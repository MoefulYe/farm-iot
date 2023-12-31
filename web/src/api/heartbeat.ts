import { request } from '@/util/requests'
import type { Dayjs } from 'dayjs'
import dayjs from 'dayjs'

export interface Heartbeat {
  id: string
  time: Dayjs
  health?: number
  weight?: number
  longitude?: number
  latitude?: number
}

export enum Field {
  Lealth = 'health',
  Weight = 'weight',
  Longitude = 'longitude',
  Latitude = 'latitude'
}

export const FieldName = {
  [Field.Lealth]: '健康',
  [Field.Weight]: '体重',
  [Field.Latitude]: '纬度',
  [Field.Longitude]: '经度'
}

export interface QueryParams {
  start?: string
  stop?: string
  fields: Field[]
}

export const fetchHeartbeat = async ({ fields, ...other }: QueryParams): Promise<Heartbeat[]> => {
  const arr: any[] | null = await request<any, any>({
    method: 'get',
    url: '/cow/heartbeat',
    params: {
      fields: fields.join(','),
      ...other
    }
  })
  if (arr === null) {
    return []
  } else {
    return arr.map((item) => {
      const { time, ...other } = item
      const time_dayjs = dayjs(time)
      return {
        time: time_dayjs,
        ...other
      }
    })
  }
}

export const fetchHeartbeatByUuid = async (
  uuid: string,
  { fields, ...other }: QueryParams
): Promise<Heartbeat[]> => {
  const arr: any[] | null = await request({
    method: 'get',
    url: `/cow/heartbeat/${uuid}`,
    params: {
      fields: fields.join(','),
      ...other
    }
  })
  if (arr === null) {
    return []
  } else {
    return arr.map((item) => {
      const { time, ...other } = item
      const time_dayjs = dayjs(time)
      return {
        time: time_dayjs,
        ...other
      }
    })
  }
}
