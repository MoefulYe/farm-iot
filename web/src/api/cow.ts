import { request } from '@/util/requests'
import dayjs, { Dayjs } from 'dayjs'
import type { Paginated } from './types'

export enum CowQueryFilter {
  Alive = 'alive',
  Dead = 'dead',
  All = 'all'
}

export interface CowQueryParams {
  page: number
  size: number
  filter: CowQueryFilter
}

export interface CowInfo {
  id: string
  born_at: Dayjs
  dead_at?: Dayjs
  reason?: string
  parent?: string
}

export const fetchCowInfoByUuid = async (uuid: string): Promise<CowInfo> => {
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

export const fetchCowInfo = async (query: CowQueryParams): Promise<Paginated<CowInfo>> => {
  let { cnt, data } = await request<any, any>({
    method: 'get',
    url: `/cow/`,
    params: query
  })
  data = data === null ? [] : data
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

export const spawnCow = async (): Promise<void> =>
  request({
    method: 'POST',
    url: '/cow/spawn/'
  })

export const KillCow = async (cows: string[]): Promise<void> =>
  request({
    method: 'POST',
    url: '/cow/kill/',
    data: {
      cows
    }
  })
