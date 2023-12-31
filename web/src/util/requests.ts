import axios, { AxiosError, type AxiosRequestConfig } from 'axios'
import { useTokenStore } from '@/stores/token'

const service = axios.create({
  baseURL: '/api',
  timeout: 30000
})

service.interceptors.request.use(
  (config) => {
    const token = useTokenStore().token
    config.headers.Authorization = token
    window.$loading.start()
    return config
  },
  (err) => {
    window.$loading.error()
    return Promise.reject(err)
  }
)

service.interceptors.response.use(
  (response) => {
    window.$loading.finish()
    return response.data.data
  },
  (err) => {
    window.$loading.error()
    if (err.response === undefined) {
      window.$message.warning(err.message)
      return Promise.reject(err.message)
    } else {
      if ((err.response.headers['content-type'] as string).startsWith('application/json')) {
        window.$message.warning(err.response.data.msg)
        return Promise.reject(err.response.data.msg)
      } else {
        window.$message.warning('unknown error')
        return Promise.reject('unknown error')
      }
    }
  }
)

export default service

export function request<T, R>(config: AxiosRequestConfig<T>): Promise<R> {
  return service(config) as unknown as Promise<R>
}
