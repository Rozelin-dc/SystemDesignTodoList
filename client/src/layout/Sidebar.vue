<script lang="ts" setup>
import { sidebarRoutes } from '@/router'
</script>

<template>
  <div class="sidebar-container">
    <div class="title">TODO List</div>
    <div v-for="route in sidebarRoutes" :key="route.name" class="sidebar-item">
      <span v-if="route.meta">{{ route.meta.title }}</span>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.sidebar-container {
  color: white;
  background-color: #2b2f3a;
  padding: 5px;

  .title {
    font-size: 30px;
    font-weight: bold;
    margin: 5px 0 20px;
  }

  .sidebar-item {
    margin: 10px 0;
  }
}
</style>
