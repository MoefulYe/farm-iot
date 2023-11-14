import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    name: 'home',
    path: '/',
    component: () => import('../views/task-view.vue')
  },
  {
    name: 'task',
    path: '/task',
    component: () => import('../views/task-view.vue')
  },
  {
    name: 'group',
    path: '/group',
    component: () => import('../views/task-group-view.vue')
  },
  {
    name: 'group/task',
    path: '/group/:id',
    component: () => import('../views/task-group-view.vue')
  },
  {
    name: 'balance',
    path: '/balance',
    component: () => import('../views/balance-view.vue')
  },
  {
    name: 'balance-stat',
    path: '/balance/stat',
    component: () => import('../views/balance-stat-view.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// router.beforeEach((to, from, next) => {})

// router.afterEach((to, from, failure) => {})

export default router
