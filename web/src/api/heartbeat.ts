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
  health = 'health',
  weight = 'weight',
  longitude = 'longitude',
  latitude = 'latitude'
}

export const FieldName = {
  [Field.health]: '健康',
  [Field.weight]: '体重',
  [Field.latitude]: '纬度',
  [Field.longitude]: '经度'
}

export interface QueryParams {
  start?: string
  stop?: string
  fields: Field[]
}

export const GetKeepAlive = async ({ fields, ...other }: QueryParams): Promise<Heartbeat[]> => {
  const arr: any[] | null = await request<any, any>({
    method: 'get',
    url: '/cow/keep-alive',
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

export const GetKeepAliveByUuid = async (
  uuid: string,
  { fields, ...other }: QueryParams
): Promise<Heartbeat[]> => {
  const arr: any[] | null = await request({
    method: 'get',
    url: `/cow/keep-alive/${uuid}`,
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
