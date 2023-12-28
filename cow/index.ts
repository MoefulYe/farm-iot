import { farm } from "./Farm";

farm.activate();
process.on("SIGTERM", () => farm.dump());
process.on("SIGINT", () => farm.dump());
