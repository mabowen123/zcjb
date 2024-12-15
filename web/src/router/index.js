import { createRouter, createWebHistory } from 'vue-router'
import {useAdminUserStore} from '@/pinia/modules/admin_user.js'

const router = createRouter({
  history: createWebHistory(), routes: [
    {
      path: '/',
      redirect: '/login'
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/view/login/index.vue')
    },
    {
      path: '/:catchAll(.*)',
      meta:
          {
            closeTab: true
          },
      component: () => import('@/view/error/index.vue')
    },
    {
      path: '/layout',
      name: "Layout",
      component: () => import('@/view/layout/index.vue'),
      children: [
        {
          path: '/dashboard',
          name: 'Dashboard',
          component: () => import('@/view/dashboard/index.vue')
        },
      ]
    }
  ],
})

const whiteList = ['Login']

router.beforeEach(async (to, _) => {
  const token = useAdminUserStore().token
  if (token) {
    if (whiteList.indexOf(to.name) > -1) {
      return {name: 'Dashboard'}
    }
  } else if (whiteList.indexOf(to.name) <= -1) {
    return {name: 'Login'}
  }
})

export default router
