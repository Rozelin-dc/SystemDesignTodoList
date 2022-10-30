<script lang="ts" setup>
import { AxiosError } from 'axios'
import { reactive, ref, watchEffect } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, FormInstance } from 'element-plus'
import { useMe } from '@/store/me'
import { UserSimple } from '@/lib/apis'
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

const inputData = reactive<UserSimple>({
  userName: '',
  password: ''
})

const loading = ref(false)
const router = useRouter()
const confirmCreate = async () => {
  if (!isFormValid.value) {
    return
  }

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
    <el-form
      ref="formRef"
      :model="inputData"
      :rules="rules"
      label-position="top"
    >
      <el-form-item prop="userName" label="ユーザー名">
        <el-input v-model="inputData.userName" />
      </el-form-item>
      <el-form-item prop="password" label="パスワード">
        <el-input
          v-model="inputData.password"
          type="password"
          show-password
          @keyup.enter="confirmCreate"
        />
      </el-form-item>
    </el-form>

    <el-button
      type="primary"
      :loading="loading"
      :disabled="!isFormValid"
      @click="confirmCreate"
    >
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
