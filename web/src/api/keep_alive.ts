import { request } from '@/util/requests'
import type { Dayjs } from 'dayjs'
import dayjs from 'dayjs'

export interface KeepAlive {
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

export interface QueryParams {
  start?: string
  stop?: string
  fields: Field[]
}

export const GetKeepAlive = async ({ fields, ...other }: QueryParams): Promise<KeepAlive[]> => {
  const arr: any[] = await request<any, any>({
    method: 'get',
    url: '/cow/keep-alive/',
    params: {
      fields: fields.join(','),
      ...other
    }
  })
  return arr.map((item) => {
    const { time, ...other } = item
    const time_dayjs = dayjs(time)
    return {
      time: time_dayjs,
      ...other
    }
  })
}

export const GetKeepAliveByUuid = async (
  uuid: string,
  { fields, ...other }: QueryParams
): Promise<KeepAlive[]> => {
  const arr: any[] = await request({
    method: 'get',
    url: `/cow/keep-alive/${uuid}/`,
    params: {
      fields: fields.join(','),
      ...other
    }
  })
  return arr.map((item) => {
    const { time, ...other } = item
    const time_dayjs = dayjs(time)
    return {
      time: time_dayjs,
      ...other
    }
  })
}
