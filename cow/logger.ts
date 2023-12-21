import winston from "winston";
import { LOGGER_LEVEL } from "./constants";

export default winston.createLogger({
  level: LOGGER_LEVEL,
  transports: [
    new winston.transports.Console({
      format: winston.format.combine(
        winston.format.colorize(),
        winston.format.timestamp(),
        winston.format.printf((info) => {
          return `[${info.timestamp} ${info.level}] ${info.message}`;
        })
      ),
    }),
  ],
});
