<script lang="ts" setup>
import Sidebar from './Sidebar.vue'
import MyHeader from './Header.vue'
</script>

<template>
  <div class="layout-container">
    <sidebar class="sidebar" />
    <my-header class="header" />
    <router-view class="main" />
  </div>
</template>

<style lang="scss" scoped>
.layout-container {
  width: 100%;
  height: 100%;
  display: grid;
  grid-template-columns: 180px 1fr;
  grid-template-rows: 40px 1fr;

  .sidebar {
    grid-column: 1 / 2;
    grid-row: 1 / 3;
  }

  .header {
    grid-column: 2 / 3;
    grid-row: 1 / 2;
  }

  .main {
    grid-column: 2 / 3;
    grid-row: 2 / 3;
  }
}
</style>
