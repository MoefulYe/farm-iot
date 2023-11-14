import '@/assets/tailwind.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import router from '@/router'
import App from '@/App.vue'

const entry = async () => {
  const pinia = createPinia().use(piniaPluginPersistedstate)
  createApp(App).use(pinia).use(router).mount('#app')
}

entry()
