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
const rules = reactive(getRules(['userName', 'password']))
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

const inputData = reactive<Omit<UpdateUser, 'newPassword'>>({
  userName: meStore.getMe?.userName ?? '',
  password: ''
})

const loading = ref(false)
const confirmUpdate = async () => {
  if (!isFormValid.value) {
    return
  }

  try {
    loading.value = true
    await meStore.changeMeData({
      ...inputData,
      newPassword: inputData.password
    })
    ElMessage({
      message: 'ユーザー名を更新しました',
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
  <div class="change-user-name-container">
    <el-form
      ref="formRef"
      :model="inputData"
      :rules="rules"
      label-position="top"
    >
      <el-form-item prop="userName" label="ユーザー名">
        <el-input v-model="inputData.userName" />
      </el-form-item>
      <el-form-item prop="password" label="パスワードを入力して更新">
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
.change-user-name-container {
  padding: 10px 5px;
}
</style>
