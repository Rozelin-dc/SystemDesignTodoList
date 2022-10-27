<script lang="ts" setup>
import { AxiosError } from 'axios'
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useMe } from '@/store/me'
import { UpdateUser } from '@/lib/apis'
import { showErrorMessage } from '@/util/showErrorMessage'

const meStore = useMe()

const inputData = reactive<Omit<UpdateUser, 'userName'>>({
  newPassword: '',
  password: ''
})

const loading = ref(false)

const confirmUpdate = async () => {
  if (!meStore.getMe) {
    return
  }

  try {
    loading.value = true
    await meStore.changeMeData({
      ...inputData,
      userName: meStore.getMe.userName
    })
    ElMessage({
      message: 'パスワードを更新しました',
      type: 'success'
    })
  } catch (e: any) {
    const err: AxiosError = e
    showErrorMessage(err)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="update-user-name-container">
    <el-form :model="inputData" label-position="top">
      <el-form-item prop="newPassword" label="新しいパスワード">
        <el-input
          v-model="inputData.newPassword"
          type="password"
          show-password
        />
      </el-form-item>
      <el-form-item prop="password" label="現在のパスワードを入力して更新">
        <el-input v-model="inputData.password" type="password" show-password />
      </el-form-item>
    </el-form>

    <el-button type="primary" :loading="loading" @click="confirmUpdate" round>
      更新
    </el-button>
  </div>
</template>

<style lang="scss" scoped>
.update-user-name-container {
  padding: 10px 5px;
}
</style>
