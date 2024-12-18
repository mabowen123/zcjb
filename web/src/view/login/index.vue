<template>
  <div class="login-container">
    <el-card class="login-card">
      <h2>用户登录</h2>
      <el-form :model="loginFormData" ref="loginForm" :rules="rules" label-width="100px">
        <el-form-item label="账号" prop="username">
          <el-input v-model="loginFormData.username" placeholder="请输入账号"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input type="password" v-model="loginFormData.password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin">登录</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import {reactive, ref} from 'vue'
import {ElMessage} from 'element-plus'
import {useAdminUserStore} from '@/pinia/modules/admin_user'
import {useRouterStore} from '@/pinia/modules/router'
import router, { childrenRouter } from '@/router/index.js'

const loginFormData = reactive({
  username: '',
  password: ''
})

const loginForm = ref(null)
const rules = {
  username: [{required: true, message: '请输入账号', trigger: 'blur'}],
  password: [{required: true, message: '请输入密码', trigger: 'blur'}]
}

const handleLogin = () => {
  loginForm.value.validate(async (v) => {
    if (!v) {
      // 未通过前端静态验证
      ElMessage.error({message: '请正确填写登录信息'})
      return
    }
    useAdminUserStore().login(loginFormData).then(res => {
      if (res) {
        ElMessage.success({message: '登录成功'})
        useRouterStore().setRouterList(childrenRouter)
        router.push({name: 'Dashboard'})
      }
    })
  })
}

</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.login-card {
  width: 400px;
}
</style>
