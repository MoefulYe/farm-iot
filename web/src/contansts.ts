export const LATITUDE_LOWWER_BOUNDARY = parseFloat(import.meta.env.VITE_LATITUDE_LOWWER_BOUNDARY)
export const LATITUDE_UPPER_BOUNDARY = parseFloat(import.meta.env.VITE_LATITUDE_UPPER_BOUNDARY)
export const LONGITUDE_LOWWER_BOUNDARY = parseFloat(import.meta.env.VITE_LONGITUDE_LOWWER_BOUNDARY)
export const LONGITUDE_UPPER_BOUNDARY = parseFloat(import.meta.env.VITE_LONGITUDE_UPPER_BOUNDARY)
export const LATITUDE_BOUNDARY = [LATITUDE_LOWWER_BOUNDARY, LATITUDE_UPPER_BOUNDARY]
export const LONGITUDE_BOUNDARY = [LONGITUDE_LOWWER_BOUNDARY, LONGITUDE_UPPER_BOUNDARY]
export const CENTER = [
  (LONGITUDE_LOWWER_BOUNDARY + LONGITUDE_UPPER_BOUNDARY) / 2,
  (LATITUDE_LOWWER_BOUNDARY + LATITUDE_UPPER_BOUNDARY) / 2
]
export const POLYGON = [
  [LONGITUDE_LOWWER_BOUNDARY, LATITUDE_LOWWER_BOUNDARY],
  [LONGITUDE_LOWWER_BOUNDARY, LATITUDE_UPPER_BOUNDARY],
  [LONGITUDE_UPPER_BOUNDARY, LATITUDE_UPPER_BOUNDARY],
  [LONGITUDE_UPPER_BOUNDARY, LATITUDE_LOWWER_BOUNDARY]
]

export const BMAP_KEY: string = import.meta.env.VITE_BMAP_SDK_KEY

export const NULL = '00000000-0000-0000-0000-000000000000'
