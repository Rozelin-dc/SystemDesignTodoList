<script lang="ts" setup>
import { AxiosError } from 'axios'
import { reactive, ref, watchEffect } from 'vue'
import { ElMessage, FormInstance } from 'element-plus'
import { useMe } from '@/store/me'
import { UpdateUser } from '@/lib/apis'
import { showErrorMessage } from '@/util/showErrorMessage'
import { getRules } from '@/util/validate'

const meStore = useMe()

const formRef = ref<FormInstance>()
const rules = reactive(getRules(['password', 'newPassword']))
const isFormValid = ref(false)
watchEffect(() => {
  const { value } = formRef
  if (!value) {
    isFormValid.value = false
    return
  }

  if (inputData.password.length === 0) {
    isFormValid.value = false
    return
  }

  value.validate(isValid =>
    isValid ? (isFormValid.value = true) : (isFormValid.value = false)
  )
})

const inputData = reactive<Omit<UpdateUser, 'userName'>>({
  newPassword: '',
  password: ''
})

const loading = ref(false)
const confirmUpdate = async () => {
  if (!meStore.getMe || !isFormValid.value) {
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
  <div class="change-password-container">
    <el-form
      ref="formRef"
      :model="inputData"
      :rules="rules"
      label-position="top"
    >
      <el-form-item prop="newPassword" label="新しいパスワード">
        <el-input
          v-model="inputData.newPassword"
          type="password"
          show-password
        />
      </el-form-item>
      <el-form-item prop="password" label="現在のパスワードを入力して更新">
        <el-input
          v-model="inputData.password"
          type="password"
          show-password
          @keyup.enter="confirmUpdate"
        />
      </el-form-item>
    </el-form>

    <el-button
      type="primary"
      :loading="loading"
      :disabled="!isFormValid"
      @click="confirmUpdate"
    >
      更新
    </el-button>
  </div>
</template>

<style lang="scss" scoped>
.change-password-container {
  padding: 10px 5px;
}
</style>
