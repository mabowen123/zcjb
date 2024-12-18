const tipRoutes = [
    {
        path: '/tip',
        name: 'Tip',
        hidden: true,
        meta: {
            title: "线报管理",
            noCache: true
        },
        children: [
            {
                path: '/tip/list',
                component: () => import('@/view/tip/data/index.vue'),
                name: 'TipList',
                meta: {
                    title: '数据管理',
                    noCache: true
                }
            },
            {
                path: '/tip/source',
                component: () => import('@/view/tip/source/index.vue'),
                name: 'TipSource',
                meta: {
                    title: '来源管理',
                    noCache: true
                }
            },
        ]
    },
]
export default tipRoutes