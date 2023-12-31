import "join"
import "sql"

option task = {
  name: "sample-to-pg",
  cron: "0 * * * *"
}

left = from(bucket: "farm")
  |> range(start: -1h)
  |> filter(fn: (r) => r["_measurement"] == "income")
  |> filter(fn: (r) => r["type"] == "kill")
  |> aggregateWindow(every: 30m, fn: sum, createEmpty: false)
  |> map(fn: (r) => ({
    when: r._time,
    in:  r._value
  }))

right = from(bucket: "farm")
  |> range(start: -1h)
  |> filter(fn: (r) => r["_measurement"] == "outcome")
  |> aggregateWindow(every: 30m, fn: sum, createEmpty: false)
  |> map(fn: (r) => ({
    when: r._time,
    out: r._value
  }))

join.inner(
  left: left,
  right: right,
  on: (l, r) => l.when == r.when,
  as: (l, r) => ({l with out: r.out})
)
  |> sql.to(
    driverName: "postgres",
    dataSourceName: "postgresql://farmer:mysecretpassword@124.221.89.92:5432/farm-iot?sslmode=disable",
    table: "balances"
  )

