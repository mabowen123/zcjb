import service from '@/utils/request'

export const login = (data) => {
    return service({
        url: '/admin/user/login',
        method: 'post',
        data,
        loadingOption: {
            fullscreen: true,
            text: '登录中，请稍候...'
        }
    })
}

export const logout = () => {
    return service({
        url: '/admin/user/logout',
        method: 'post'
    })
}