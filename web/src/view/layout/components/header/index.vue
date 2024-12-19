<template>
  <el-menu
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
  >
    <el-menu-item>
      <img
          src="@/assets/logo.svg"
          style="width: 100px;height: 100px"
          alt="Element logo"
      />
    </el-menu-item>
    <el-sub-menu index="0">
      <template #title>个人中心</template>
      <el-menu-item index="0-1" @click="logout">登出</el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>
<script>
import {useAdminUserStore} from '@/pinia/modules/admin_user'
import router from '@/router/index.js'

export default {
  name: 'Header',
  methods: {
    logout() {
      useAdminUserStore().logout().then(_ => {
        router.push('/login')
      })
    }
  }

}
</script>
