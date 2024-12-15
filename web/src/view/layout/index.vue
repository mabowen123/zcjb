<template>
  <div class="common-layout">
    <el-container>
      <el-header style="height: auto ">
        <el-menu
            :default-active="activeIndex2"
            class="el-menu-demo"
            mode="horizontal"
            :ellipsis="false"
            @select="handleSelect"
        >
          <el-menu-item index="0">
            <img
                style="width: 50px"
                src="../../../public/logo.svg"
                alt="Element logo"
            />
          </el-menu-item>
          <el-menu-item index="1">Processing Center</el-menu-item>
          <el-sub-menu index="2">
            <template #title>Workspace</template>
            <el-menu-item index="2-1">item one</el-menu-item>
            <el-menu-item index="2-2">item two</el-menu-item>
            <el-menu-item index="2-3">item three</el-menu-item>
            <el-sub-menu index="2-4">
              <template #title>item four</template>
              <el-menu-item index="2-4-1">item one</el-menu-item>
              <el-menu-item index="2-4-2">item two</el-menu-item>
              <el-menu-item index="2-4-3">item three</el-menu-item>
            </el-sub-menu>
          </el-sub-menu>
        </el-menu>
      </el-header>
      <el-container>
        <el-aside style="width: 400px">
          <el-row class="tac">
            <el-col :span="12">
              <el-menu
                  :default-active=activeIndex
                  class="el-menu-vertical-demo"
                  router
                  @open="handleOpen"
                  @close="handleClose"
              >
                <el-sub-menu index="1">
                  <template #title>
                    <el-icon>
                      <location/>
                    </el-icon>
                    <span>Navigator One</span>
                  </template>
                  <el-menu-item-group title="Group One">
                    <el-menu-item index="1-1">item one</el-menu-item>
                    <el-menu-item index="/dashboard" key="/dashboard">数据大盘</el-menu-item>
                  </el-menu-item-group>
                </el-sub-menu>
                <el-menu-item index="/dashboard" key="/dashboard">
                  <el-icon>
                    <icon-menu/>
                  </el-icon>
                  <span>数据大盘</span>
                </el-menu-item>
                <el-menu-item index="3" disabled>
                  <el-icon>
                    <document/>
                  </el-icon>
                  <span>Navigator Three</span>
                </el-menu-item>
                <el-menu-item index="4">
                  <el-icon>
                    <setting/>
                  </el-icon>
                  <span>Navigator Four</span>
                </el-menu-item>
              </el-menu>
            </el-col>
          </el-row>
        </el-aside>
        <el-main>
          这是首页
          <router-view v-show="activeIndex !='/dashboard'" :key="activeIndex"/>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>
<script lang="ts" setup>
import {ref} from 'vue'
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from '@element-plus/icons-vue'

const activeIndex = ref('1')
const activeIndex2 = ref('1')
const handleSelect = (key, keyPath) => {
  console.log(key, keyPath)
}

const handleOpen = (key, keyPath) => {
  console.log(key, keyPath)
}
const handleClose = (key, keyPath) => {
  console.log(key, keyPath)
}
</script>
<style>
.el-menu--horizontal > .el-menu-item:nth-child(1) {
  margin-right: auto;
}
</style>


