const mqtt = require('mqtt')
import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb';
const protobuf = require('protobufjs')
const protopath = "../proto/cow"



run().catch(err => console.log(err));
async function run() {
    //console.log(buf.toString('utf8')); // Gnarly string that contains "Bill"
    //const obj = RegisterReq.decode(buf);
    //console.log(obj);
    const host = 'broker.emqx.io'
    const port = '1883'
    const uuid = `mqtt_${Math.random().toString(16).slice(3)}`

    const connectUrl = `mqtt://${host}:${port}`
    const client = mqtt.connect(connectUrl, {
        clientId: uuid,
        clean: true,
        connectTimeout: 4000,
        username: 'emqx',
        password: 'public',
        reconnectPeriod: 1000,
    })

    const topic = `/cow/${uuid}`
    const topicrep = `${topic}/register-reply`
    client.on('connect', async () => {
        var token = await register(client, uuid, passwd)
        client.subscribe([topic], () => {
            console.log(`Subscribe to topic '${topic}'`)
        })
        setInterval(() => {
            keepalive(client, token, x, y, weight, health)
        }, 1000 * 60);
    })
    client.on('message', (topic, payload) => {
        console.log('Received Message:', topic, payload.toString())
    })
}

function getstamp() {
    var timestamp = Timestamp.fromDate(new Date());
    var date = timestamp.toDate();
    return data
}

async function register(client, uuid, passwd) {
    const topic = `/cow/${uuid}`
    const topicrep = `${topic}/register-reply`
    const root = await protobuf.load(`${protopath}/register.proto`);
    const RegisterReq = root.lookupType('farm.cow.RegisterReq');
    const RegisterResp = root.lookupType('farm.cow.RegisterResp');
    var timestamp = getstamp()
    var payload = { born_at: timestamp, uuid: uuid, passwd: passwd }
    var message = RegisterReq.create(payload);
    const buf = RegisterReq.encode(message).finish();
    var token
    await client.publish('/cow/register', buf, { qos: 0, retain: false }, (error) => {
        client.on('message', (topicrep, payload) => {
            const respmessage = RegisterResp.decode(payload)
            if (respmessage.status == 0 && respmessage.uuid == uuid)
                token = respmessage.token
            else
                token = login(client, uuid, passwd)
        })
    })
    return token
}

async function login(client, uuid, passwd) {
    const topic = `/cow/${uuid}`
    //const topicrep = `${topic}/login-reply`
    const root2 = await protobuf.load(`${protopath}/login.proto`);
    const LoginReq = root2.lookupType('farm.cow.LoginReq');//todo
    const LoginResp = root2.lookupType('farm.cow.RegisterResp');//todo
    var payload = { uuid: uuid, passwd: passwd }
    var message = LoginReq.create(payload);
    const buf = LoginReq.encode(message).finish();
    await client.publish('/cow/login', buf, { qos: 0, retain: false }, (error) => {
        client.on('message', (topicrep, payload) => {
            const respmessage = LoginResp.decode(payload)
            if (respmessage.status == 0)
                token = respmessage.token
            else {
                console.log("uuid or passwd error")
                return null
            }
        })
    })
    return token
}

async function keepalive(client, token, x, y, weight, health) {
    const root = await protobuf.load(`${protopath}/keep_live.proto`);
    const GeoCoordinate = root.lookupType('farm.cow.GeoCoordinate');
    const KeepAliveMsg = root.lookupType('farm.cow.KeepAliveMsg');
    var geo = GeoCoordinate.create({ latitude: x, longitude: y })
    var payload = { timestamp: getstamp(), token: token, geo: geo, weight: weight, health: health }
    var message = KeepAliveMsg.create(payload);
    const buf = KeepAliveMsg.encode(message).finish();
    client.publish('/cow/keep-alive', buf, { qos: 0, retain: false }, (error) => {
        if (err)
            console.log(err)
    })
}