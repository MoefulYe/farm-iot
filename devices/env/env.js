/** InfluxDB v2 URL */
const url = 'http://124.221.89.92:8086'
/** InfluxDB authorization token */
const token = 'EFr0bgEQip96t-RL99r6rURvBzj0MFi4LtC-vCpIKaQYu4CjKm5M59xXakfL2NtLMArwPlXUykhrinwJVD53Zg=='
/** Organization within InfluxDB  */
const org = "farm-iot"
/**InfluxDB bucket used in examples  */
const bucket = 'farm-iot'
// ONLY onboarding example
/**InfluxDB user  */
const username = 'my-user'
/**InfluxDB password  */
const password = 'my-password'

const threshold = 200
module.exports = { url, token, org, bucket, username, password, threshold }