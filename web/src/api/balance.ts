import { request } from '@/util/requests'

export interface Balance {
  when: string
  in: number
  out: number
}

export interface BalanceQueryParams {
  from?: string
  to?: string
}

export const fetchBalance = async (params: BalanceQueryParams): Promise<Balance[]> =>
  request({
    method: 'GET',
    url: '/balance/',
    params
  })
