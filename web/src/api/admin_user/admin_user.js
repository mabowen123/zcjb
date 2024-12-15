import service from '@/utils/request'

export const login = (data) => {
  return service({
    url: '/admin/user/login', method: 'post', data
  })
}

export const logout = () => {
  return service({
    url: '/admin/user/logout', method: 'post'
  })
}