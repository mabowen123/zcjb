<template>
  <el-menu
      unique-opened
      class="el-menu-vertical-demo"
      router
      default-active="/dashboard"
      style="height: calc(100vh - 100px);"
  >
    <SideItem :dataList="childrenRouter"/>
  </el-menu>
</template>

<script>
import SideItem from "./sideItem.vue"
import {childrenRouter} from '@/router/index.js'

export default {
  name: "SideBar",
  data() {
    return {
      childrenRouter
    }
  },
  components: {
    SideItem
  }
}

</script>