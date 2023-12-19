const mqtt = require('mqtt')
const protobuf = require('protobufjs')
const protopath = "../proto/cow"
const { v4: uuidv4 } = require('uuid')
const fs = require('fs')
const { kill } = require('process')
const INTERVAL = 1000 * 5


const killer = mqtt.connect(Cow.connectUrl, {
    clientId: uuidv4(),
    clean: true,
    connectTimeout: 4000,
    username: 'admin',
    password: 'admin',
    reconnectPeriod: 1000,
})

killer.subscribe("killer/command/kill")
killer.on("message", async (t, payload, packet) => {
    const root2 = await protobuf.load(`${protopath}/die.proto`);
    const DieMsg = root2.lookupType('farm.cow.DieMsg');//todo
    var payload1 = { timestamp: now(), uuid: uuid, reason: "killed" }
    var message = DieMsg.create(payload1);
    const buf = DieMsg.encode(message).finish();
    killer.publishAsync(`cow/${uuid}/die`, buf, { qos: 2, retain: false })
})
setInterval(() => {

}, interval);

function now() {
    return new Date().toISOString()
}