import { v4 } from "uuid";
import { connect, MqttClient } from "mqtt";
import { EQMX_PASSWORD, EQMX_URL, EQMX_USERNAME } from "./constants";
import logger from "./logger";

export interface State {
  born_at: string;
  uuid: string;
  longitude: number;
  latitude: number;
  weight: number;
  health: number;
  hp: number;
  token: string;
  passwd: string;
}

export default class Cow {
  private state: State;
  public getState(): State {
    return this.state;
  }
  private client: MqttClient;
  private timeOut?: NodeJS.Timeout;
  private constructor(state: State) {
    this.state = state;
    this.client = connect(EQMX_URL, {
      clientId: state.uuid,
      clean: true,
      connectTimeout: 4000,
      username: EQMX_USERNAME,
      password: EQMX_PASSWORD,
    });
    this.timeOut = undefined;
    this.client.subscribeAsync(`cow/${this.state.uuid}/command/#`);
    this.client.on("message", (topic, payload) => {
      switch (topic) {
        default:
      }
    });
  }

  public static async fromState(state: State): Promise<Cow> {
    const cow = new Cow(state);
    if (state.token === "") {
      logger.log(
        "info",
        `cow ${state.uuid} has no token, requesting from iot-server`
      );
      cow.state.token = await cow.login();
      logger.log(
        "info",
        `cow ${cow.state.uuid} login with token ${cow.state.token}`
      );
    }
    return cow;
  }

  private async login(): Promise<string> {
    const topic = `cow/${this.state.uuid}/login-reply`;
    return "";
  }
}
