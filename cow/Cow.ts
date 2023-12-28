import { v4 } from "uuid";
import { connect, MqttClient } from "mqtt";
import {
  EQMX_PASSWORD,
  EQMX_URL,
  EQMX_USERNAME,
  HEARTBEAT_INTERVAL,
  LATITUDE_CENTER,
  LONGITUDE_CENTER,
} from "./constants";
import logger from "./logger";
import {
  encodeLoginReq,
  decodeLoginResp,
  Status as LoginStatus,
} from "./protoc/login";
import {
  encodeRegisterReq,
  decodeRegisterResp,
  Status as RegisterStatus,
} from "./protoc/register";
import { encodeHeartBeat } from "./protoc/heartbeat";
import { encodeDie } from "./protoc/command";
import { farm } from "./Farm";
import dayjs from "dayjs";

export interface State {
  born_at: string;
  uuid: string;
  longitude: number;
  latitude: number;
  weight: number;
  health: number;
  healthPoint: number;
  token: string;
  passwd: string;
  parent?: string;
  chilren: string[];
}

export const fix = (val: number, [minBound, maxBound]: [number, number]) =>
  Math.max(minBound, Math.min(maxBound, val));
export default class Cow {
  private static readonly INIT_HP = 15;
  private static readonly INIT_WEIGHT = 5;
  private static readonly INIT_HEALTH = 1;

  private state: State;
  public getState(): State {
    return this.state;
  }
  private client: MqttClient;
  public timeOut?: NodeJS.Timeout;

  private constructor(state: State) {
    this.state = state;
    this.client = connect(EQMX_URL, {
      clientId: `cow-${this.state.uuid}`,
      clean: true,
      connectTimeout: 4000,
      username: EQMX_USERNAME,
      password: EQMX_PASSWORD,
    });
  }

  public static async newCow(): Promise<Cow> {
    const cow = new Cow({
      born_at: dayjs().format(),
      uuid: v4(),
      longitude: LONGITUDE_CENTER,
      latitude: LATITUDE_CENTER,
      weight: this.INIT_WEIGHT,
      health: this.INIT_HEALTH,
      healthPoint: this.INIT_HP,
      token: "",
      passwd: Math.random().toString(16).slice(3),
      chilren: [],
    });
    await cow.register();
    logger.info(`new cow ${cow.state.uuid}`);
    return cow;
  }

  private async breed(): Promise<Cow> {
    const cow = new Cow({
      born_at: dayjs().format(),
      uuid: v4(),
      longitude: this.state.longitude,
      latitude: this.state.latitude,
      weight: Cow.INIT_WEIGHT,
      health: Cow.INIT_HEALTH,
      healthPoint: Cow.INIT_HP,
      token: "",
      passwd: Math.random().toString(16).slice(3),
      parent: this.state.uuid,
      chilren: [],
    });
    await cow.register();
    this.state.chilren.push(cow.state.uuid);
    logger.info(`cow-${this.state.uuid} breed child cow-${cow.state.uuid}`);
    return cow;
  }

  public static async fromState(state: State): Promise<Cow> {
    const cow = new Cow(state);
    logger.info(`recover cow-${cow.state.uuid} from state`);
    if (state.token === "") {
      logger.info(`cow ${state.uuid} has no token, requesting from iot-server`);
      await cow.login();
    }
    return cow;
  }

  private listen() {
    this.client.subscribeAsync([
      `cow/${this.state.uuid}/command/+`,
      "cow/broadcast/command/+",
    ]);
    const kill = `cow/${this.state.uuid}/command/kill`;
    const banish = `cow/${this.state.uuid}/command/banish`;
    const cure = "cow/broadcast/command/cure";
    this.client.on("message", (topic, _) => {
      switch (topic) {
        case kill:
          this.kill();
          break;
        case banish:
          this.banish();
          break;
        case cure:
          this.cure();
          break;
        default:
      }
    });
  }

  private async login(): Promise<void> {
    const reply = `cow/${this.state.uuid}/login-reply`;
    const msg = encodeLoginReq({
      uuid: this.state.uuid,
      passwd: this.state.passwd,
    });
    await this.client.subscribeAsync(reply);
    await this.client.publishAsync("cow/login", Buffer.from(msg));

    this.state.token = await new Promise<string>((resolve, reject) => {
      const event = this.client.on("message", (topic, payload) => {
        event.removeAllListeners("message");
        if (topic === reply) {
          const resp = decodeLoginResp(payload);
          resp.status = resp.status ?? LoginStatus.STATUS_OK;
          if (resp.token === undefined) {
            const msg = `cow-${this.state.uuid} login failed! get invalid package`;
            logger.error(msg);
            reject(msg);
          } else if (resp.status !== LoginStatus.STATUS_OK) {
            const msg = `cow-${this.state.uuid} login failed! status code ${resp.status}`;
            logger.error(msg);
            reject(msg);
          } else {
            const token = resp.token;
            logger.info(
              `cow-${this.state.uuid} login success with token ${token}`
            );
            resolve(token);
          }
        } else {
          const msg = `cow-${this.state.uuid} login failed! expect topic ${reply}, found ${topic}`;
          logger.error(msg);
          reject(msg);
        }
      });
    });

    await this.client.unsubscribeAsync(reply);
  }

  private async register(): Promise<void> {
    const reply = `cow/${this.state.uuid}/register-reply`;
    const req = encodeRegisterReq({
      born_at: this.state.born_at,
      uuid: this.state.uuid,
      passwd: this.state.passwd,
      parent: this.state.parent,
    });
    await this.client.subscribeAsync(reply);
    await this.client.publishAsync("cow/register", Buffer.from(req));
    this.state.token = await new Promise<string>((resolve, reject) => {
      const event = this.client.on("message", (topic, payload) => {
        event.removeAllListeners("message");
        if (topic === reply) {
          const resp = decodeRegisterResp(payload);
          resp.status = resp.status ?? RegisterStatus.STATUS_OK;
          if (resp.uuid === undefined || resp.token === undefined) {
            const msg = `cow-${this.state.uuid} register failed! get invalid package`;
            logger.error(msg);
            reject(msg);
          } else if (resp.status !== RegisterStatus.STATUS_OK) {
            const msg = `cow-${this.state.uuid} register failed! status code ${resp.status}`;
            logger.error(msg);
            reject(msg);
          } else {
            const token = resp.token;
            logger.info(
              `cow-${this.state.uuid} register success with token ${token}`
            );
            resolve(token);
          }
        } else {
          const msg = `cow-${this.state.uuid} register failed! expect topic ${reply}, found ${topic}`;
          logger.error(msg);
          reject(msg);
        }
      });
    });

    await this.client.unsubscribeAsync(reply);
  }

  private cure(): void {
    this.state.health = Math.min(this.state.health + 0.2, 1);
  }

  private async heartBeat(): Promise<void> {
    const heartBeat = {
      timestamp: dayjs().format(),
      token: this.state.token,
      longitude: this.state.longitude,
      latitude: this.state.latitude,
      weight: this.state.weight,
      health: this.state.health,
    };
    logger.debug(
      `cow-${this.state.uuid} longitude: ${this.state.longitude}, latitude: ${this.state.latitude}, weight: ${this.state.weight}, health: ${this.state.health}, hp: ${this.state.healthPoint}`
    );
    const pkt = encodeHeartBeat(heartBeat);
    await this.client.publishAsync("cow/heartbeat", Buffer.from(pkt));
  }

  private mutate() {
    this.mutateHealth();
    this.mutateWeight();
    this.mutateLocation();
  }

  private mutateHealth() {
    this.state.health = fix(
      this.state.health + (Math.random() - 0.8) * 0.05,
      [0, 1]
    );
    this.state.healthPoint -= 1 - this.state.health;
  }

  private mutateLocation() {
    const delta_lo = (Math.random() - 0.5) * 0.0005;
    const delta_la = (Math.random() - 0.5) * 0.0005;
    this.state.longitude += delta_lo;
    this.state.latitude += delta_la;
  }

  private mutateWeight() {
    this.state.weight += 10 * (this.state.health - 0.2);
  }

  private async banish() {
    const longitude = this.state.longitude;
    const latitude = this.state.latitude;
    let delta_lo = LONGITUDE_CENTER - longitude;
    let delta_la = LATITUDE_CENTER - latitude;
    const distance = Math.sqrt(delta_lo ** 2 + delta_la ** 2);
    const unit = 0.002;
    delta_la *= unit / distance;
    delta_lo *= unit / distance;
    this.state.longitude += delta_lo;
    this.state.latitude += delta_la;
    logger.info(
      `cow-${this.state.uuid} is banished from (${longitude}, ${latitude}) to (${this.state.longitude}, ${this.state.latitude})`
    );
  }

  private async ill() {
    const msg = encodeDie({
      timestamp: dayjs().format(),
      reason: "ill",
      uuid: this.state.uuid,
      token: this.state.token,
      weight: this.state.weight,
      health: this.state.health,
      latitude: this.state.latitude,
      longitude: this.state.longitude,
    });
    await this.client.publishAsync("cow/die", Buffer.from(msg));
    this.client.end();
    clearInterval(this.timeOut);
    farm.logout(this.state.uuid);
    logger.info(`cow-${this.state.uuid} is died`);
  }

  private isDead(): boolean {
    return this.state.healthPoint <= 0;
  }

  private breedable(): boolean {
    return (
      this.state.weight > 100 &&
      this.state.chilren.length < 4 &&
      Math.random() > 0.8
    );
  }

  public async kill() {
    const msg = encodeDie({
      timestamp: dayjs().format(),
      reason: "kill",
      uuid: this.state.uuid,
      token: this.state.token,
      weight: this.state.weight,
      health: this.state.health,
      latitude: this.state.latitude,
      longitude: this.state.longitude,
    });
    await this.client.publishAsync("cow/die", Buffer.from(msg));
    this.client.end();
    clearInterval(this.timeOut);
    farm.logout(this.state.uuid);
    logger.info(`cow-${this.state.uuid} is killed`);
  }

  public async butch() {
    const msg = encodeDie({
      timestamp: dayjs().format(),
      reason: "kill",
      uuid: this.state.uuid,
      token: this.state.token,
      weight: this.state.weight,
      health: this.state.health,
      latitude: this.state.latitude,
      longitude: this.state.longitude,
    });
    await this.client.publishAsync("cow/die", Buffer.from(msg));
    this.client.end();
    clearInterval(this.timeOut);
    logger.info(`cow-${this.state.uuid} is killed`);
  }

  private async run(): Promise<void> {
    this.mutate();
    if (this.isDead()) {
      await this.ill();
      return;
    }
    await this.heartBeat();
    if (this.breedable()) {
      const child = await this.breed();
      farm.register(child);
    }
  }

  public async activate(): Promise<void> {
    this.listen();
    this.timeOut = setInterval(() => this.run(), HEARTBEAT_INTERVAL);
    logger.info(`activate cow-${this.state.uuid}`);
  }
}
