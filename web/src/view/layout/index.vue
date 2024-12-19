<template>
  <el-container>
    <el-header style="height: 100px;padding: 0">
      <Header/>
    </el-header>
    <el-container style="max-height: 100%">
      <el-aside>
        <SideBar/>
      </el-aside>
      <el-main>
        <AppMain/>
      </el-main>
    </el-container>
  </el-container>
</template>
<script lang="ts">
import SideBar from "./components/sideBar/index.vue"
import AppMain from "./components/appMain/index.vue"
import Header from "./components/header/index.vue"

export default {
  components: {
    SideBar,
    AppMain,
    Header
  }
}

</script>
<style>
.el-menu--horizontal > .el-menu-item:nth-child(1) {
  margin-right: auto;
}
</style>


