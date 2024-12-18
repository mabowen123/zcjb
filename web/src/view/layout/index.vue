<template>
  <div class="common-layout">
    <el-container>
      <el-header style="height: auto ">
    <Header/>
      </el-header>
      <el-container>
        <el-aside>
          <el-col :span="12">
            <SideBar/>
          </el-col>
        </el-aside>
        <el-main>
          <AppMain/>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>
<script lang="ts" setup>
import SideBar from "./components/sideBar/index.vue"
import AppMain from "./components/appMain/index.vue"
import Header from "./components/header/index.vue"

defineOptions({
  components: {
    SideBar,
    AppMain,
    Header
  }
})

</script>
<style>
.el-menu--horizontal > .el-menu-item:nth-child(1) {
  margin-right: auto;
}
</style>


