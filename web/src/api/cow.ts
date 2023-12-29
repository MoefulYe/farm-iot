import { request } from '@/util/requests'
import dayjs, { Dayjs } from 'dayjs'
import type { Pagination, Paginated } from './types'

export interface CowInfo {
  id: string
  born_at: Dayjs
  dead_at?: Dayjs
  reason?: string
  parent: string
}

export const GetCowInfoByUuid = async (uuid: string): Promise<CowInfo> => {
  const { dead_at, born_at, ...other } = await request<any, any>({
    method: 'get',
    url: `/cow/${uuid}/`
  })
  const dead_at_dayjs = dead_at ? dayjs(dead_at) : undefined
  const born_at_dayjs = dayjs(born_at)
  return {
    ...other,
    born_at: born_at_dayjs,
    dead_at: dead_at_dayjs
  }
}

export const GetCowInfo = async (query: Pagination): Promise<Paginated<CowInfo>> => {
  const { cnt, data } = await request<any, any>({
    method: 'get',
    url: `/cow/`,
    params: query
  })
  const d = (data as any[]).map((item) => {
    const { dead_at, born_at, ...other } = item
    const dead_at_dayjs = dead_at ? dayjs(dead_at) : undefined
    const born_at_dayjs = dayjs(born_at)
    return {
      born_at: born_at_dayjs,
      dead_at: dead_at_dayjs,
      ...other
    }
  })
  return {
    cnt,
    data: d
  }
}
