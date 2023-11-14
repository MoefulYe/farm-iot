import { request } from '@/util/requests'

export interface Token {
  token: string
}

export interface LoginReq {
  username: string
  password: string
}

export interface RegisterReq {
  username: string
  password: string
}

export const Login = async (params: LoginReq): Promise<Token> =>
  request({
    method: 'post',
    url: '/login',
    data: params
  })

export const Register = async (params: RegisterReq): Promise<Token> =>
  request({
    method: 'post',
    url: '/register',
    data: params
  })
