<template>
  <el-menu
      unique-opened
      class="el-menu-vertical-demo"
      router
      default-active="/dashboard">
    <SideItem :dataList="useRouterStore().routerList" />
  </el-menu>
</template>

<script setup>
import SideItem from "./sideItem.vue"
import {useRouterStore} from "@/pinia/modules/router.js";

defineOptions({
  name: "SideBar",
  components: {
    SideItem
  }
})
</script>