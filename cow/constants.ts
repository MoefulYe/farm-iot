export const HEARTBEAT_INTERVAL = 1000 * 60 * 5;
export const LOGGER_LEVEL = process.env.LOGGER_LEVEL || "debug";
export const LOCATION = [
  [120.35, 30.31],
  [120.37, 30.31],
  [120.37, 30.29],
  [120.35, 30.29],
];

export const LONGITUDE_MIN = Math.min(...LOCATION.map((point) => point[0]));
export const LONGITUDE_MAX = Math.max(...LOCATION.map((point) => point[0]));
export const LATITUDE_MIN = Math.min(...LOCATION.map((point) => point[1]));
export const LATITUDE_MAX = Math.max(...LOCATION.map((point) => point[1]));
export const LONGITUDE_CENTER = (LONGITUDE_MAX + LONGITUDE_MIN) / 2;
export const LATITUDE_CENTER = (LATITUDE_MAX + LATITUDE_MIN) / 2;

export const LONGITUDE_BOUND = [LONGITUDE_MIN, LONGITUDE_MAX];
export const LATITUDE_BOUND = [LATITUDE_MIN, LATITUDE_MAX];
export const CENTER = [LONGITUDE_CENTER, LATITUDE_CENTER];

export const fixLongitude = (longitude: number) =>
  Math.max(LONGITUDE_MIN, Math.min(LONGITUDE_MAX, longitude));
export const fixLatitude = (latitude: number) =>
  Math.max(LATITUDE_MIN, Math.min(LATITUDE_MAX, latitude));

export const EQMX_HOST = process.env.EQMX_HOST!;
export const EQMX_PORT = process.env.EQMX_PORT!;
export const EQMX_URL = `mqtt://${EQMX_HOST}:${EQMX_PORT}`;
export const EQMX_USERNAME = process.env.EQMX_USERNAME!;
export const EQMX_PASSWORD = process.env.EQMX_PASSWORD!;
