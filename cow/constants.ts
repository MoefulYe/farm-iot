export const HEARTBEAT_INTERVAL = 1000 * 60 * 5;
export const LOGGER_LEVEL = Bun.env.LOGGER_LEVEL;
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

export const LONGITUDE_BOUND: [number, number] = [LONGITUDE_MIN, LONGITUDE_MAX];
export const LATITUDE_BOUND: [number, number] = [LATITUDE_MIN, LATITUDE_MAX];
export const CENTER: [number, number] = [LONGITUDE_CENTER, LATITUDE_CENTER];

export const fix = (val: number, [minBound, maxBound]: [number, number]) =>
  Math.max(minBound, Math.min(maxBound, val));

export const fixLongitude = (longitude: number) =>
  fix(longitude, LONGITUDE_BOUND);
export const fixLatitude = (latitude: number) => fix(latitude, LATITUDE_BOUND);

export const EQMX_HOST = Bun.env.EQMX_HOST;
export const EQMX_PORT = Bun.env.EQMX_PORT;
export const EQMX_URL = `mqtt://${EQMX_HOST}:${EQMX_PORT}`;
export const EQMX_USERNAME = Bun.env.EQMX_USERNAME;
export const EQMX_PASSWORD = Bun.env.EQMX_PASSWORD;
export const FARM = Bun.env.FARM;
