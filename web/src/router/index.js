import {createRouter, createWebHistory} from 'vue-router'
import {useAdminUserStore} from '@/pinia/modules/admin_user.js'
import tipRoutes from "@/router/tip/index.js";
import movieRouter from "@/router/movie/index.js";

const router = createRouter({
    history: createWebHistory(), routes: [
        {

            path: '/redirect/:path*',
            component: () => import('@/view/dashboard/index.vue')
        },
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
            meta: {
                closeTab: true
            },
            component: () => import('@/view/error/index.vue')
        },
        {
            name: "layout",
            component: () => import('@/view/layout/index.vue'),
        },
    ],
})

export let childrenRouter = [
    {
        path: '/dashboard',
        name: 'Dashboard',
        component: () => import('@/view/dashboard/index.vue'),
        meta: {
            title: '数据大盘',
            icon: 'dashboard',
            noCache: true
        }
    },
]
childrenRouter = [...childrenRouter, ...tipRoutes, ...movieRouter]
childrenRouter.forEach(item => {
    router.addRoute('layout', item)
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
