declare module "bun" {
  interface Env {
    EQMX_HOST: string;
    EQMX_PORT: string;
    EQMX_USERNAME: string;
    EQMX_PASSWORD: string;
    FARM: string;
    LOGGER_LEVEL: string;
  }
}
