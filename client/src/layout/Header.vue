<script lang="ts" setup>
import { AxiosError } from 'axios'
import { useRoute, useRouter } from 'vue-router'
import { useMe } from '@/store/me'
import { showErrorMessage } from '@/util/showErrorMessage'

const route = useRoute()
const router = useRouter()

const meStore = useMe()
const doLogout = async () => {
  try {
    await meStore.logout()
    router.push({ name: 'Login' })
  } catch (e: any) {
    const err: AxiosError = e
    showErrorMessage(err)
  }
}
</script>

<template>
  <div class="header-container">
    <span class="title">{{ route.meta.title }}</span>
    <span>ユーザー名: {{ meStore.getMe?.userName }}</span>
    <el-button @click="doLogout" class="button" link>ログアウト</el-button>
  </div>
</template>

<style lang="scss" scoped>
.header-container {
  display: flex;
  justify-content: end;
  align-items: center;
  background-color: rgb(167, 224, 255);
  padding: 0 5px;

  .title {
    flex: 1;
    font-size: 24px;
    font-weight: bold;
  }

  .button {
    margin-left: 10px;
  }
}
</style>
