const mqtt = require("mqtt");
const protobuf = require("protobufjs");
const protopath = "../proto/cow";
const { v4: uuidv4 } = require("uuid");
const farm = require("../farm.json");
const fs = require("fs");
const COWS = "./cow.json";
const INTERVAL = 1000 * 60 * 5;

const cows = [];

class Cow {
  constructor(state) {
    this.state = state;
    this.client = mqtt.connect(Cow.connectUrl, {
      clientId: this.uuid,
      clean: true,
      connectTimeout: 4000,
      username: "admin",
      password: "admin",
      reconnectPeriod: 1000,
    });
    this.handler = null;
    this.client.subscribeAsync(`cow/${this.state.uuid}/command/#`, { qos: 1 });
    this.client.on("message", (t, payload, packet) => {
      if (t == `cow/${this.state.uuid}/command/kill`) this.die();
    });
  }

  static async newCow() {
    const cow = new Cow({
      born_at: new Date().toISOString(),
      uuid: uuidv4(),
      longitude:
        (farm.location[0][0] +
          farm.location[1][0] +
          farm.location[2][0] +
          farm.location[3][0]) /
        4,
      latitude:
        (farm.location[0][1] +
          farm.location[1][1] +
          farm.location[2][1] +
          farm.location[3][1]) /
        4,
      weight: Cow.weight,
      health: Cow.health,
      hp: Cow.hp,
      token: "",
      passwd: Math.random().toString(16).slice(3),
    });
    const token = await cow.register();
    console.log(token);
    cow.state.token = token;
    return cow;
  }

  static async fromState(state) {
    const cow = new Cow(state);
    if (cow.state.token == "") {
      cow.state.token = await cow.login();
    }
    return cow;
  }

  static health = 0;
  static hp = 100;
  static weight = 5;
  static host = "124.221.89.92";
  static port = "1883";
  static connectUrl = `mqtt://${Cow.host}:${Cow.port}`;
  async run() {
    this.stateTrans();
    if (this.state.hp <= 0) {
      return "ill";
    } else {
      await this.keepalive();
      return "ok";
    }
  }

  async die() {
    this.client.end();
    clearInterval(this.handler);
    const idx = cows.findIndex((cow) => cow.state.uuid == this.state.uuid);
    if (idx != -1) {
      cows.splice(idx, 1);
    }
  }
  async ill() {
    const root2 = await protobuf.load(`${protopath}/die.proto`);
    const DieMsg = root2.lookupType("farm.cow.DieMsg"); //todo
    const payload = {
      timestamp: this.now(),
      uuid: this.state.uuid,
      reason: "ill",
      token: this.state.token,
      weight: this.state.weight,
      health: this.state.health,
      latitude: this.state.latitude,
      longitude: this.state.longitude,
    };
    const message = DieMsg.create(payload);
    const buf = DieMsg.encode(message).finish();
    this.client.publishAsync(`cow/die`, buf, { qos: 2, retain: false });
    this.client.end();
    clearInterval(this.handler);
    const idx = cows.findIndex((cow) => cow.state.uuid == this.state.uuid);
    if (idx != -1) {
      cows.splice(idx, 1);
    }
  }

  async register() {
    const topic = `cow/${this.state.uuid}/register-reply`;
    const root = await protobuf.load(`${protopath}/register.proto`);
    const RegisterReq = root.lookupType("farm.cow.RegisterReq");
    const RegisterResp = root.lookupType("farm.cow.Registerconst tokenResp");
    const payload = {
      bornAt: this.state.born_at,
      uuid: this.state.uuid,
      passwd: this.state.passwd,
    };
    const message = RegisterReq.create(payload);
    const buf = RegisterReq.encode(message).finish();
    await this.client.subscribeAsync(`cow/${this.state.uuid}/register-reply`, {
      qos: 0,
    });
    await this.client.publishAsync("cow/register", buf, {
      qos: 0,
      retain: false,
    });
    const data = await new Promise((resolve, reject) => {
      const event = this.client.on("message", (t, payload, packet) => {
        if (t == topic) {
          const data = RegisterResp.decode(payload);
          event.removeAllListeners("message");
          resolve(data);
        } else {
          event.removeAllListeners("message");
          reject("error");
        }
      });
    });
    await this.client.unsubscribeAsync(
      `cow/${this.state.uuid}/register-reply`,
      { qos: 0 }
    );
    if (data.status != 0) {
      throw new Error("invalid passwd or uuid");
    }
    return data.token;
  }
  async login() {
    const topic = `cow/${this.state.uuid}/login-reply`;
    const root2 = await protobuf.load(`${protopath}/login.proto`);
    const LoginReq = root2.lookupType("farm.cow.LoginReq"); //todo
    const LoginResp = root2.lookupType("farm.cow.LoginResp"); //todo
    const payload = { uuid: this.state.uuid, passwd: this.state.passwd };
    const message = LoginReq.create(payload);
    const buf = LoginReq.encode(message).finish();
    await this.client.subscribeAsync(`cow/${this.state.uuid}/login-reply`, {
      qos: 0,
    });
    await this.client.publishAsync("cow/login", buf, { qos: 0, retain: false });
    const data = await new Promise((resolve, reject) => {
      const event = this.client.on("message", (t, payload, packet) => {
        if (t == topic) {
          const data = LoginResp.decode(payload);
          event.removeAllListeners("message");
          resolve(data);
        } else {
          event.removeAllListener("message");
          reject("error");
        }
      });
    });
    await this.client.unsubscribeAsync(`cow/${this.state.uuid}/login-reply`, {
      qos: 0,
    });
    if (data.status != 0) {
      throw new Error("invalid passwd or uuid");
    }
    return data.token;
  }
  async keepalive() {
    console.log(
      `[${this.now()}] ${this.state.uuid} ${this.state.longitude} ${
        this.state.latitude
      } ${this.state.weight} ${this.state.health} ${this.state.hp}`
    );
    const root = await protobuf.load(`${protopath}/keep_live.proto`);
    const GeoCoordinate = root.lookupType("farm.cow.GeoCoordinate");
    const KeepAliveMsg = root.lookupType("farm.cow.KeepAliveMsg");
    const msg1 = GeoCoordinate.create({
      latitude: this.state.latitude,
      longitude: this.state.longitude,
    });
    const payload = {
      timestamp: this.now(),
      token: this.state.token,
      geo: msg1,
      weight: this.state.weight,
      health: this.state.health,
    };
    const message = KeepAliveMsg.create(payload);
    const buf = KeepAliveMsg.encode(message).finish();
    this.client.publish("cow/keep-alive", buf, { qos: 0, retain: false });
  }
  now() {
    return new Date().toISOString();
  }
  stateTrans() {
    this.state.health = Math.min(
      this.state.health + (Math.random() - 0.6) * 0.01,
      1
    );
    this.state.weight += 0.1 * (this.state.health - 0.3);
    this.state.hp += Math.min(this.state.health, 0);
    const [longitude, latitude] = this.newGeo();
    this.state.longitude = longitude;
    this.state.latitude = latitude;
  }
  newGeo() {
    const delta_long = (Math.random() - 0.5) * 0.0005;
    const delta_lang = (Math.random() - 0.5) * 0.0005;
    let long = this.state.longitude + delta_long;
    let lang = this.state.latitude + delta_lang;
    const max_long = Math.max(...farm.location.map((p) => p[0]));
    const min_long = Math.min(...farm.location.map((P) => P[0]));
    const max_lang = Math.max(...farm.location.map((p) => p[1]));
    const min_lang = Math.min(...farm.location.map((P) => P[1]));
    long = Math.min(max_long, long);
    long = Math.max(min_long, long);
    lang = Math.min(max_lang, lang);
    lang = Math.max(min_lang, lang);
    return [long, lang];
  }
}

const topEntry = async () => {
  await recover();

  for (const cow of cows) {
    cow.handler = setIntferval(async () => {
      switch (await cow.run()) {
        case "ok":
          break;
        case "ill":
          await cow.ill();
          break;
        case "reproduce":
          break;
        default:
          throw Error("unreachable");
      }
    }, INTERVAL);
  }
  process.on("SIGINT", save);
  process.on("SIGTERM", save);
};

const save = async () => {
  const str = JSON.stringify(cows.map((cow) => cow.state));
  try {
    fs.writeFileSync(COWS, str, { flag: "w", encoding: "utf-8" });
  } catch (err) {
    console.error(err);
  }
  process.exit(0);
};

const recover = async () => {
  const content = fs.readFileSync(COWS, {
    encoding: "utf-8",
    flag: "r",
  });
  const data = JSON.parse(content);
  cows = await Promise.all(data.map((state) => Cow.fromState(state)));
};

topEntry();
