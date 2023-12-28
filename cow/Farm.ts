import { MqttClient, connect } from "mqtt";
import Cow, { State } from "./Cow";
import { EQMX_PASSWORD, EQMX_URL, EQMX_USERNAME, FARM } from "./constants";
import logger from "./logger";

export default class Farm {
  private static readonly BUTCH_INTERVAL = 1234 * 46 * 23;
  private cows: Cow[];
  private client: MqttClient;
  private constructor(cows: Cow[]) {
    this.cows = cows;
    this.client = connect(EQMX_URL, {
      clientId: "farm",
      clean: true,
      connectTimeout: 4000,
      username: EQMX_USERNAME,
      password: EQMX_PASSWORD,
    });
  }
  public static async load(): Promise<Farm> {
    const file = Bun.file(FARM);
    if (await file.exists()) {
      const states: State[] = await file.json();
      const cows = await Promise.all(
        states.map((state) => Cow.fromState(state))
      );
      logger.info(`load farm from \`${FARM}\``);
      return new Farm(cows);
    } else {
      const cow = await Cow.newCow();
      logger.info(`create new cow ${cow.getState().uuid}`);
      return new Farm([cow]);
    }
  }

  public async dump(): Promise<void> {
    const json = JSON.stringify(this.cows.map((cow) => cow.getState()));
    const dest = Bun.file(FARM);
    Bun.write(dest, json);
    logger.info(`dump farm to \`${FARM}\``);
    process.exit(0);
  }

  private listen() {
    this.client.subscribe("farm/spawn");
    this.client.on("message", async (topic) => {
      switch (topic) {
        case "farm/spawn":
          this.spawn();
          break;
      }
    });
  }

  private async spawn() {
    const cow = await Cow.newCow();
    logger.info(`spawn cow ${cow.getState().uuid}`);
    this.cows.push(cow);
    cow.activate();
  }

  public register(child: Cow) {
    this.cows.push(child);
    child.activate();
  }

  public logout(uuid: string) {
    const idx = this.cows.findIndex((cow) => cow.getState().uuid === uuid);
    if (idx !== -1) {
      this.cows.splice(idx, 1);
    }
  }

  public activate() {
    this.listen();
    this.butch();
    setInterval(() => this.butch(), Farm.BUTCH_INTERVAL);
    this.cows.forEach((cow) => cow.activate());
  }

  private butch() {
    const num = this.cows.sort(
      (a, b) => a.getState().weight - b.getState().weight
    ).length;
    if (num < 100) {
      logger.info(`cows num is ${num} and less than 100, no cow is killed`);
    } else if (num < 200) {
      logger.info(`cows num is ${num} and less than 200, 10 cows are killed`);
      this.cows.splice(num - 10, 10).forEach((cow) => cow.butch());
    } else if (num < 300) {
      logger.info(`cows num is ${num} and less than 300, 20 cows are killed`);
      this.cows.splice(num - 20, 20).forEach((cow) => cow.butch());
    } else if (num < 400) {
      logger.info(`cows num is ${num} and less than 400, 30 cows are killed`);
      this.cows.splice(num - 30, 30).forEach((cow) => cow.butch());
    } else {
      logger.info(
        `cows num is ${num} and more than 400, ${num - 300} cows are killed`
      );
      this.cows.splice(300, num - 300).forEach((cow) => cow.butch());
    }
  }
}

export const farm = await Farm.load();
