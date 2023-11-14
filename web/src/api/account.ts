import { request } from '@/util/requests'
import type { Paginated } from './types'
import dayjs from 'dayjs'

enum Category {
  Other = 'other',
  Meal = 'meal',
  Fruit = 'fruit',
  Snack = 'snack',
  Consumable = 'consumable',
  Medical = 'medical',
  Transportation = 'transportation',
  Entertainment = 'entertainment',
  Social = 'social',
  Clothing = 'clothing',
  Housing = 'housing',
  Dormitory = 'dormitory',
  Subscriptions = 'subscriptions',
  Hardware = 'hardware',
  Software = 'software',
  Books = 'books',
  Education = 'education',
  Salary = 'salary',
  Borrow = 'borrow',
  Class = 'class',
  Gift = 'gift'
}

const CategoryMap = {
  [Category.Other]: '其他',
  [Category.Meal]: '餐饮',
  [Category.Fruit]: '水果',
  [Category.Snack]: '零食',
  [Category.Consumable]: '耗材',
  [Category.Medical]: '医疗',
  [Category.Transportation]: '交通',
  [Category.Entertainment]: '娱乐',
  [Category.Social]: '社交',
  [Category.Clothing]: '服饰',
  [Category.Housing]: '住房',
  [Category.Dormitory]: '宿舍',
  [Category.Subscriptions]: '订阅',
  [Category.Hardware]: '硬件',
  [Category.Software]: '软件',
  [Category.Books]: '书籍',
  [Category.Education]: '教育',
  [Category.Salary]: '工资',
  [Category.Borrow]: '出入',
  [Category.Class]: '班级',
  [Category.Gift]: '礼物'
}

export const categoryName = (category: Category): string => CategoryMap[category]

export interface Account {
  id: number
  name: string
  category: Category
  balance: number
  commet: string
  yyyy: number
  mm: number
  dd: number
}

export interface GetAccountsQueryParams {
  page?: number
  size?: number
  year?: number
  month?: number
  day?: number
  filter?: Category[]
}

export const getAccounts = (params: GetAccountsQueryParams): Promise<Paginated<Account>> =>
  request({
    method: 'get',
    url: '/account',
    params
  })

export interface CreateAccountReq {
  name?: string
  category: Category
  balance: number
  commet?: string
}

export const createAccount = (req: CreateAccountReq): Promise<Account> =>
  request({
    method: 'post',
    url: '/account',
    data: {
      ...req,
      yyyy: dayjs().year(),
      mm: dayjs().month() + 1,
      dd: dayjs().date()
    }
  })

export const getAccountbyId = (id: number): Promise<Account> =>
  request({
    method: 'get',
    url: `/account/${id}`
  })

export interface UpdateAccountReq {
  name?: string
  category?: Category
  balance?: number
  commet?: string
}

export const updateAccountbyId = (id: number, req: UpdateAccountReq): Promise<Account> =>
  request({
    method: 'patch',
    url: `/account/${id}`,
    data: req
  })

export const deleteAccountbyId = (id: number): Promise<void> =>
  request({
    method: 'delete',
    url: `/account/${id}`
  })

export interface Stats {
  other: number
  meal: number
  fruit: number
  snack: number
  consumable: number
  medical: number
  transportation: number
  entertainment: number
  social: number
  clothing: number
  housing: number
  dormitory: number
  subscriptions: number
  hardware: number
  software: number
  books: number
  education: number
  salary: number
  borrow: number
  class: number
  gift: number
  income: number
  expense: number
  total: number
}

export interface CategoryFilter {
  filter?: Category[]
}

export const getStats = (): Promise<Stats> =>
  request({
    method: 'get',
    url: '/account/stat'
  })

export interface AnnualStats {
  year: number
  other: number
  meal: number
  fruit: number
  snack: number
  consumable: number
  medical: number
  transportation: number
  entertainment: number
  social: number
  clothing: number
  housing: number
  dormitory: number
  subscriptions: number
  hardware: number
  software: number
  books: number
  education: number
  salary: number
  borrow: number
  class: number
  gift: number
  income: number
  expense: number
  total: number
}

export interface MonthlyStats {
  year: number
  month: number
  other: number
  meal: number
  fruit: number
  snack: number
  consumable: number
  medical: number
  transportation: number
  entertainment: number
  social: number
  clothing: number
  housing: number
  dormitory: number
  subscriptions: number
  hardware: number
  software: number
  books: number
  education: number
  salary: number
  borrow: number
  class: number
  gift: number
  income: number
  expense: number
  total: number
}

export const getAnnualStats = (): Promise<AnnualStats[]> =>
  request({
    method: 'get',
    url: '/account/stat/annual'
  })

export const getAnnualStat = (yyyy: number): Promise<AnnualStats> =>
  request({
    method: 'get',
    url: `/account/stat/annual/${yyyy}`
  })

export const getMonthlyStats = (): Promise<MonthlyStats[]> =>
  request({
    method: 'get',
    url: '/account/stat/monthly'
  })

export const getMonthlyStatsbyYyyy = (yyyy: number): Promise<MonthlyStats[]> =>
  request({
    method: 'get',
    url: `/account/stat/monthly/${yyyy}`
  })

export const getMonthlyStat = (yyyy: number, mm: number): Promise<MonthlyStats> =>
  request({
    method: 'get',
    url: `/account/stat/monthly/${yyyy}/${mm}`
  })
