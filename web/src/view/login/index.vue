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

<script lang="ts">
import {ElMessage} from 'element-plus'
import {useAdminUserStore} from '@/pinia/modules/admin_user'
import router from '@/router/index.js'

export default {
  data() {
    return {
      loginFormData: {
        username: '',
        password: ''
      },
      rules: {
        username: [{required: true, message: '请输入账号', trigger: 'blur'}],
        password: [{required: true, message: '请输入密码', trigger: 'blur'}]
      }
    }
  },
  methods: {
    handleLogin() {
      this.$refs.loginForm.validate(async (v) => {
        if (v) {
          useAdminUserStore().login(this.loginFormData).then(res => {
            if (res) {
              ElMessage.success({message: '登录成功'})
              router.push({name: 'Dashboard'})
            }
          })
        }
      })
    }
  }
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
