<script lang="ts" setup>
import { sidebarRoutes } from '@/router'
</script>

<template>
  <div class="sidebar-container">
    <div class="title">TODO List</div>
    <div v-for="route in sidebarRoutes" :key="route.name" class="sidebar-item">
      <router-link v-if="route.meta" :to="route" class="link">{{
        route.meta.title
      }}</router-link>
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

    .link {
      color: white;
      text-decoration: none;
    }
  }
}
</style>
