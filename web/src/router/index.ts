import loginModal from '@/components/login-modal'
import { useTokenStore } from '@/stores/token'
import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    name: 'home',
    path: '/',
    redirect: {
      name: 'real-time'
    }
  },
  {
    name: 'real-time',
    path: '/real-time',
    component: () => import('../views/real-time.vue')
  },
  {
    name: 'cow-info',
    path: '/cow-info',
    component: () => import('../views/cow-info.vue')
  },
  {
    name: 'stat',
    path: '/stat/:uuid',
    component: () => import('../views/stat-view.vue')
  },
  {
    name: 'blank',
    path: '/blank',
    component: () => import('../views/blank-view.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach(async () => {
  if (!useTokenStore().isLogin()) {
    router.push({ name: 'blank' })
    loginModal()
  }
})

// router.afterEach((to, from, failure) => {})

export default router
