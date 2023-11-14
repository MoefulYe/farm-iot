import { request } from '@/util/requests'
import { Dayjs } from 'dayjs'
import type { Pagination, Paginated } from './types'

export interface CowInfo {
  id: string
  born_at: Dayjs
  dead_at?: Dayjs
  reason?: string
}

export const GetCowInfoByUuid = async (uuid: string): Promise<CowInfo> =>
  request({
    method: 'get',
    url: `/cow/${uuid}`
  })

export const GetCowInfo = async (query: Pagination): Promise<Paginated<CowInfo>> =>
  request({
    method: 'get',
    url: `/cow`,
    params: query
  })
