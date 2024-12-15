import {defineStore} from 'pinia'
import {login, logout} from '@/api/admin_user/admin_user.js'

export const useAdminUserStore = defineStore('admin_user',
    {
        state: () => {
            return {
                token: null
            }
        },
        getters: {
            getToken() {
                return this.token
            }
        },
        actions: {
            setToken(token) {
                if (token) {
                    this.token = token
                }
            },
            removeToken() {
                this.token = ""
            },
            login(loginInfo) {
                return login(loginInfo).then(res => {
                    this.setToken(res.token)
                    return true
                }).catch(_ => {
                    return false
                })
            },
            logout() {
                return logout().then(_ => {
                    this.removeToken()
                    return true
                }).catch(_ => {
                    return false
                })
            }
        }

    })