<script lang="ts" setup>
import { AxiosError } from 'axios'
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { useMe } from '@/store/me'
import { UserSimple } from '@/lib/apis'
import { showErrorMessage } from '@/util/showErrorMessage'

const meStore = useMe()

const inputData = reactive<UserSimple>({
  userName: '',
  password: ''
})

const loading = ref(false)
const router = useRouter()
const confirmCreate = async () => {
  try {
    loading.value = true
    await meStore.createMe(inputData)
    ElMessage({
      message: 'アカウントを作成しました',
      type: 'success'
    })
    router.push({ name: 'Home' })
  } catch (e: any) {
    const err: AxiosError = e
    showErrorMessage(err)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="create-user-container">
    <div class="title">アカウント作成</div>
    <el-form :model="inputData" label-position="top">
      <el-form-item prop="userName" label="ユーザー名">
        <el-input v-model="inputData.userName" />
      </el-form-item>
      <el-form-item prop="password" label="パスワード">
        <el-input v-model="inputData.password" type="password" show-password />
      </el-form-item>
    </el-form>

    <el-button type="primary" :loading="loading" @click="confirmCreate">
      作成
    </el-button>
  </div>
</template>

<style lang="scss" scoped>
.create-user-container {
  padding: 40px 30px;

  .title {
    font-size: 24px;
    font-weight: bold;
    margin-bottom: 20px;
  }
}
</style>
