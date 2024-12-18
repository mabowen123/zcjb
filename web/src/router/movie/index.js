const movieRouter = [
    {
        path: '/movie',
        name: 'Movie',
        meta: {
            title: "影视管理",
            noCache: true
        },
        children: [
            {
                path: '/movie/list',
                component: () => import('@/view/movie/data/index.vue'),
                name: 'MovieList',
                meta: {
                    title: '数据管理',
                    noCache: true
                }
            },
            {
                path: '/movie/source',
                component: () => import('@/view/movie/source/index.vue'),
                name: 'MovieSource',
                meta: {
                    title: '来源管理',
                    noCache: true
                }
            },
        ]
    }
]

export default movieRouter
