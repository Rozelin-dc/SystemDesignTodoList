<script lang="ts" setup>
import { onMounted } from 'vue'
import { AxiosError } from 'axios'
import { useMe } from '@/store/me'
import { showErrorMessage } from './util/showErrorMessage';

const meStore = useMe()

onMounted(async () => {
  try {
    await meStore.setMe()
  } catch (e: any) {
    const err: AxiosError = e
    showErrorMessage(err)
  }
})
</script>

<template>
  <router-view />
</template>
