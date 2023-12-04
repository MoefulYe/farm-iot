import { useTokenStore } from '@/stores/token'
import { request } from '@/util/requests'

export interface Token {
  token: string
}

export interface LoginReq {
  username: string
  passwd: string
}

export interface RegisterReq {
  username: string
  passwd: string
}

export const Login = async (params: LoginReq): Promise<Token> => {
  const token: Token = await request({
    method: 'post',
    url: '/login',
    data: params
  })
  useTokenStore().setToken(token.token)
  return token
}

export const Register = async (params: RegisterReq): Promise<Token> => {
  const token: Token = await request({
    method: 'post',
    url: '/register',
    data: params
  })
  useTokenStore().setToken(token.token)
  return token
}
