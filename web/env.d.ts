/// <reference types="vite/client" />
interface ImportMetaEnv {
    readonly VITE_LONGITUDE_LOWWER_BOUNDARY: string
    readonly VITE_LONGITUDE_UPPER_BOUNDARY: string
    readonly VITE_LATITUDE_LOWWER_BOUNDARY: string
    readonly VITE_LATITUDE_UPPER_BOUNDARY: string
    readonly VITE_BMAP_SDK_KEY: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}
