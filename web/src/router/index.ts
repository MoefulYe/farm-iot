import loginModal from '@/components/login-modal'
import { useTokenStore } from '@/stores/token'
import { createRouter, createWebHashHistory, type RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    name: 'home',
    path: '/',
    redirect: {
      name: 'dashboard'
    }
  },
  {
    name: 'dashboard',
    path: '/dashboard',
    component: () => import('../views/dashboard-view.vue')
  },
  {
    name: 'cow-manage',
    path: '/cow',
    component: () => import('../views/cow-view.vue')
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
  history: createWebHashHistory(),
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
