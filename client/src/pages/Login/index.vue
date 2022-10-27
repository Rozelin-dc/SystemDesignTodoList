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
const login = async () => {
  try {
    loading.value = true
    await meStore.login(inputData)
    ElMessage({
      message: 'ログインに成功しました',
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
  <div class="change-user-name-container">
    <div class="title">TODO List</div>
    <el-form :model="inputData" label-position="top">
      <el-form-item prop="userName" label="ユーザー名">
        <el-input v-model="inputData.userName" />
      </el-form-item>
      <el-form-item prop="password" label="パスワード">
        <el-input v-model="inputData.password" type="password" show-password />
      </el-form-item>
    </el-form>

    <el-button class="button" type="primary" :loading="loading" @click="login">
      ログイン
    </el-button>
  </div>
</template>

<style lang="scss" scoped>
.change-user-name-container {
  width: 500px;
  margin: 0 auto;
  padding-top: 20vh;

  .title {
    font-size: 30px;
    font-weight: bold;
    margin-bottom: 20px;
    text-align: center;
  }

  .button {
    width: 100%;
    margin-top: 10px;
  }
}
</style>
