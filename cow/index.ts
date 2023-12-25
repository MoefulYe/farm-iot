import { MqttClient, connect, connectAsync } from "mqtt";
import Cow, { State } from "./Cow";
import { EQMX_PASSWORD, EQMX_URL, EQMX_USERNAME, FARM } from "./constants";
import logger from "./logger";

let farm: Cow[] = [];

const recover = async () => {
  const file = Bun.file(FARM);
  if (await file.exists()) {
    const states: State[] = await file.json();
    farm = await Promise.all(states.map((state) => Cow.fromState(state)));
    logger.info(`recover states from \`${FARM}\``);
  } else {
    farm.push(await Cow.newCow());
  }
};

const save = async () => {
  const json = JSON.stringify(farm.map((cow) => cow.getState()));
  const dest = Bun.file(FARM);
  Bun.write(dest, json);
  logger.info(`save states to \`${FARM}\``);
  process.exit(0);
};

const onCowDie = (uuid: string) => {
  const idx = farm.findIndex((cow) => cow.getState().uuid === uuid);
  if (idx !== -1) {
    farm.splice(idx, 1);
  }
};

const onCowBreed = (child: Cow) => {
  farm.push(child);
  child.activate(onCowDie, onCowBreed);
};

connect(EQMX_URL, {
  clientId: "spawner",
  clean: true,
  connectTimeout: 4000,
  username: EQMX_USERNAME,
  password: EQMX_PASSWORD,
})
  .subscribe("spawner/spawn")
  .on("message", async () => {
    const cow = await Cow.newCow();
    logger.info(`spawn cow ${cow.getState().uuid}`);
    farm.push(cow);
    cow.activate(onCowDie, onCowBreed);
  });

await recover();
farm.forEach((cow) => cow.activate(onCowDie, onCowBreed));
process.on("SIGTERM", save);
process.on("SIGINT", save);
