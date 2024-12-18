<template>
  <el-menu
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
  >
    <el-menu-item>
      <img
          style="width: 50px"
          src="@/assets/logo.svg"
          alt="Element logo"
      />
    </el-menu-item>
    <el-sub-menu index="0">
      <template #title>个人中心</template>
      <el-menu-item index="0-1" @click="logout">登出</el-menu-item>
      <el-menu-item index="0-2" @click="clear()">清理缓存</el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>
<script setup>
import {useAdminUserStore} from '@/pinia/modules/admin_user'
import {useRouterStore} from '@/pinia/modules/router.js'
import router, {childrenRouter} from '@/router/index.js'

defineOptions({
  name: 'Header'
})
const logout = () => {
  useAdminUserStore().logout().then(_ => {
    router.push('/login')
  })
}
const clear = () => {
  useRouterStore().setRouterList(childrenRouter)
}

</script>
