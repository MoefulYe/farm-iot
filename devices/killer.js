const mqtt = require('mqtt')
const protobuf = require('protobufjs')
const protopath = "../proto/cow"
const { v4: uuidv4 } = require('uuid')
const INTERVAL = 1000 * 5
const { InfluxDB, FluxTableMetaData } = require('@influxdata/influxdb-client')
const { url, token, org, threshold } = require('./env/env.js')

const queryApi = new InfluxDB({ url, token }).getQueryApi(org)

const fluxQuery1 =
    'from(bucket:"farm-iot") |> range(start: -2d) |> filter(fn: (r) => r._measurement == "cow")'
//|> filter(fn: (r) => "weight")`
//|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")`
const fluxQuery =
    `from(bucket: "farm-iot")
|> range(start: -5m)
|> filter(fn: (r) => r["_measurement"] == "cow")
|> filter(fn: (r) => r["_field"] == "weight")
|> yield(name: "mean")`

async function iterateRows() {
    console.log('*** IterateRows ***')
    let res = new Array()
    for await (const { values, tableMeta } of queryApi.iterateRows(fluxQuery)) {
        // the following line creates an object for each row
        const o = tableMeta.toObject(values)
        //console.log(JSON.stringify(o, null, 2))
        //console.log(`${o._time} ${o._measurement} in  ${o._field}=${o._value},${o.uuid}`)
        //console.log(tableMeta.get(row, '_value'))
        // alternatively, you can get only a specific column value without
        // the need to create an object for every row
        // console.log(tableMeta.get(row, '_time'))
        res.push([o._value, o.uuid])
    }
    console.log('\nIterateRows SUCCESS')
    //console.log(res)
    return res
}



const killer = mqtt.connect("mqtt://124.221.89.92:1883", {
    clientId: "killer",
    clean: true,
    connectTimeout: 4000,
    username: 'admin',
    password: 'admin',
    reconnectPeriod: 1000,
})

killer.subscribe("killer/kill")

killer.on("message", async (t, payload, packet) => {
    if (t == "killer/kill") {
        const root2 = await protobuf.load(`${protopath}/die.proto`);
        const DieMsg = root2.lookupType('farm.cow.DieMsg');
        const KillReq = root2.lookupType('farm.cow.Diereq')

        const msg = KillReq.decode(payload)
        const uuid = msg.uuid
        killer.publishAsync(`cow/${uuid}/command/kill`, '', { qos: 2, retain: false })
        var payload1 = { timestamp: now(), uuid: uuid, reason: "killed" }
        var message = DieMsg.create(payload1);
        const buf = DieMsg.encode(message).finish();
        killer.publishAsync(`cow/${uuid}/die`, buf, { qos: 2, retain: false })
    }

})
setInterval(async () => {
    let res = await iterateRows()
    console.log(res)
    if (res.length > threshold) {
        res.sort(function (a, b) {
            return a[0] - b[0];
        })
        for (i = threshold; i < res.length; i++) {
            token = res[i][1]

            const root2 = await protobuf.load(`${protopath}/die.proto`);
            const DieMsg = root2.lookupType('farm.cow.DieMsg');
            killer.publishAsync(`cow/${uuid}/command/kill`, '', { qos: 2, retain: false })
            var payload1 = { timestamp: now(), uuid: uuid, reason: "killed" }
            var message = DieMsg.create(payload1);
            const buf = DieMsg.encode(message).finish();
            killer.publishAsync(`cow/${uuid}/die`, buf, { qos: 2, retain: false })
        }
    }
}, INTERVAL);





function now() {
    return new Date().toISOString()
}