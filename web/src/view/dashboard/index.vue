<template>
  <div>
    <el-button @click="logout">登出</el-button>
  </div>
</template>
<script lang="ts" setup>
import {useAdminUserStore} from '@/pinia/modules/admin_user'
import router from '@/router/index.js'

const logout = () => {
  useAdminUserStore().logout().then(_ => {
    router.push('/login')
  })
}
</script>
