import {defineStore} from "pinia";

export const useRouterStore = defineStore('user_router', {
    state: () => {
        return {
            routerList: null
        }
    },
    actions: {
        setRouterList(routerList) {
            this.routerList = routerList
        }
    }
})