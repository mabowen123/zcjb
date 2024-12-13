<template>
  <div class="login-container">
    <el-card class="login-card">
      <h2 class="login-title">用户登录</h2>
      <el-form :model="loginForm" :rules="loginRules" ref="loginFormRef" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="loginForm.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input
              v-model="loginForm.password"
              placeholder="请输入密码"
              type="password"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin">登录</el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import {ref, reactive} from 'vue';
import {ElMessage} from 'element-plus';
import {FormInstance} from 'element-plus'

const loginForm = reactive({
  username: '',
  password: '',
});

const loginRules = reactive({
  username: [
    {required: true, message: '请输入用户名', trigger: 'blur'},
    {min: 3, message: '用户名长度至少为3位', trigger: 'blur'},
  ],
  password: [
    {required: true, message: '请输入密码', trigger: 'blur'},
    {min: 6, message: '密码长度至少为6位', trigger: 'blur'},
  ],
});

const loginFormRef = ref(FormInstance);

const handleLogin = async () => {
  console.log(!loginFormRef.validate())

};
const resetForm = () => {
  loginFormRef.value.resetFields();
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f4f4f4;
}

.login-card {
  width: 360px;
}

.login-title {
  text-align: center;
  margin-bottom: 20px;
}
</style>
