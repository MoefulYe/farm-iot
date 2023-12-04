import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useTokenStore = defineStore(
  'token',
  () => {
    const token = ref('')
    const isLogin = () => token.value !== ''
    const setToken = (newToken: string) => {
      token.value = newToken
    }
    const clearToken = () => {
      token.value = ''
    }
    return { token, isLogin, setToken, clearToken }
  },
  {
    persist: {
      storage: window.localStorage
    }
  }
)
