import { request } from '@/util/requests'
import type { Dayjs } from 'dayjs'

export interface KeepAlive {
  id: string
  time: Dayjs
  health?: number
  weight?: number
  longtitude?: number
  latitude?: number
}

export enum Field {
  health = 'health',
  weight = 'weight',
  longtitude = 'longtitude',
  latitude = 'latitude'
}

export interface QueryParams {
  start?: string
  stop?: string
  fields: Field[]
}

export const GetKeepAlive = async ({ fields, ...other }: QueryParams): Promise<KeepAlive[]> =>
  request({
    method: 'get',
    url: '/cow/keep-alive',
    params: {
      fields: fields.join(','),
      ...other
    }
  })

export const GetKeepAliveByUuid = async (
  uuid: string,
  { fields, ...other }: QueryParams
): Promise<KeepAlive[]> =>
  request({
    method: 'get',
    url: `/cow/keep-alive/${uuid}`,
    params: {
      fields: fields.join(','),
      ...other
    }
  })
