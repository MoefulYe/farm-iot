import { request } from '@/util/requests'
import type { Paginated } from './types'

export interface Task {
  id: number
  title: string
  urgency: number
  status: Status
  commet: string
  deadline: Date | null
  group_id: number | null
}

export enum SortBy {
  CreateAt = 'create_at',
  Deadline = 'deadline',
  Urgency = 'urgency'
}

export enum Status {
  Todo = 'todo',
  Done = 'done',
  Discard = 'discard'
}

export interface TaskQueryParams {
  filter?: Status[]
  sort_by?: SortBy
  page?: number
  size?: number
}

export const getTasks = (params: TaskQueryParams): Promise<Paginated<Task>> =>
  request({
    method: 'get',
    url: '/task',
    params
  })

export interface CreateTaskReq {
  title: string
  urgency?: number
  commet?: string
  deadline?: Date
  group_id?: number
}

export const createTask = (params: CreateTaskReq): Promise<Task> =>
  request({
    method: 'post',
    url: '/task',
    data: params
  })

export const getTaskbyId = (id: number): Promise<Task> =>
  request({
    method: 'get',
    url: `/task/${id}`
  })

export interface UpdateTaskReq {
  title?: string
  urgency?: number
  status?: Status
  commet?: string
  deadline?: Date | null
  group_id?: number | null
}

export const updateTaskbyId = (id: number, params: UpdateTaskReq): Promise<Task> =>
  request({
    method: 'patch',
    url: `/task/${id}`,
    data: params
  })

export const deleteTaskbyId = (id: number): Promise<void> =>
  request({
    method: 'delete',
    url: `/task/${id}`
  })

export interface TaskGroup {
  id: number
  name: string
  commet: string
  create_at: Date
}

export interface Pagination {
  page?: number
  size?: number
}

export const getTaskGroups = (params: Pagination): Promise<Paginated<TaskGroup>> =>
  request({
    method: 'get',
    url: '/group',
    params
  })

export interface CreateTaskGroupReq {
  name: string
  commet?: string
}

export const createTaskGroup = (params: CreateTaskGroupReq): Promise<TaskGroup> =>
  request({
    method: 'post',
    url: '/group',
    data: params
  })

export const getTaskGroupbyId = (id: number): Promise<TaskGroup> =>
  request({
    method: 'get',
    url: `/group/${id}`
  })

export interface UpdateTaskGroupReq {
  name?: string
  commet?: string
}

export const updateTaskGroupbyId = (id: number, req: UpdateTaskGroupReq): Promise<TaskGroup> =>
  request({
    method: 'patch',
    url: `/group/${id}`,
    data: req
  })

export const deleteTaskGroupbyId = (id: number): Promise<void> =>
  request({
    method: 'delete',
    url: `/group/${id}`
  })

export interface TaskGroupwithTasks {
  group: TaskGroup
  tasks: Task[]
}

export const getTaskGroupwithTask = (
  pagination: Pagination
): Promise<Paginated<TaskGroupwithTasks>> =>
  request({
    method: 'get',
    url: '/group/task',
    params: pagination
  })

export const getTaskGroupwithTaskbyId = (id: number): Promise<TaskGroupwithTasks> =>
  request({
    method: 'get',
    url: `/group/task/${id}`
  })
