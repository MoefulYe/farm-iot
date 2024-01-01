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
  born_at: string
  dead_at?: string
  reason?: string
  parent: string
}

export interface CowInfoWithChildren {
  id: string
  born_at: string
  dead_at?: string
  reason?: string
  parent: string
  edges: {
    children?: {
      id: string
    }[]
  }
}

export const fetchCowInfoByUuid = (uuid: string): Promise<CowInfoWithChildren> =>
  request<any, any>({
    method: 'get',
    url: `/cow/${uuid}/`
  })

export const fetchCowInfo = async (query: CowQueryParams): Promise<Paginated<CowInfo>> => {
  let { cnt, data } = await request<any, any>({
    method: 'get',
    url: `/cow/`,
    params: query
  })
  data = data === null ? [] : data
  return {
    cnt,
    data
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
