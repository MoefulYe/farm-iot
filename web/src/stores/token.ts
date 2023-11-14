import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useTokenStore = defineStore(
  'token',
  () => {
    const token = ref('')
    const isLogin = () => token.value !== ''
    return { token, isLogin }
  },
  {
    persist: {
      storage: window.localStorage
    }
  }
)
