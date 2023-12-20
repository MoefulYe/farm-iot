import Cow from "./Cow";
import logger from "./logger";

const farm: Cow[] = [];

const topEntry = async () => {};

const input = Bun.file("input.txt");
await Bun.write(Bun.stdout, input);
