<template>
  <router-view v-slot="{ Component, route }">
    <transition mode="out-in" name="el-fade-in-linear">
      <keep-alive>
        <component :is="Component" :key="route.fullPath"/>
      </keep-alive>
    </transition>
  </router-view>
</template>
<script>
export default {
  name: "AppMain",

}
</script>
