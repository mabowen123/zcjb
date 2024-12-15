<template>
  <div>
    <div class="w-full h-screen bg-gray-50 flex items-center justify-center">
      <div class="flex flex-col items-center text-2xl gap-4">
        <p class="text-lg">页面被神秘力量吸走了</p>
        <el-button @click="toDashboard">返回首页</el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import router from '@/router/index.js'

const toDashboard = () => {
  router.push('/dashboard')
}
</script>
