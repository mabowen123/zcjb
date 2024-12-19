<template>
  <template v-for="(item,index) in dataList" :key="index">
    <el-menu-item v-if="!item.children" :index="item.path">
      <template #title>
        <span>{{ item.meta.title }}</span>
      </template>
    </el-menu-item>
    <el-sub-menu v-else :key="item.path" :index="item.path">
      <template #title>
        <span>{{ item.meta.title }}</span>
      </template>
      <side-item :dataList="item.children"></side-item>
    </el-sub-menu>
  </template>
</template>
<script>
export default {
  name: 'SiteItem',
  props: {
    dataList: {
      type: Array,
      default: []
    }
  }
}

</script>
