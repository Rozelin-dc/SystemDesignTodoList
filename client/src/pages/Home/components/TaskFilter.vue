<script lang="ts" setup>
import { AxiosError } from 'axios'
import { reactive, ref, watchEffect } from 'vue'
import { ElMessage, FormInstance } from 'element-plus'
import { useTask } from '@/store/task'
import { TaskStatus } from '@/types/taskStatus'
import { showErrorMessage } from '@/util/showErrorMessage'

const formRef = ref<FormInstance>()
const rules = reactive({
  name: [
    { min: 0, max: 60, message: '検索文字列が長すぎます', trigger: 'blur' }
  ]
})
const isFormValid = ref(false)
watchEffect(() => {
  const { value } = formRef
  if (!value) {
    isFormValid.value = false
    return
  }

  if (inputData.name.length === 0) {
    isFormValid.value = true
    return
  }

  value.validate(isValid =>
    isValid ? (isFormValid.value = true) : (isFormValid.value = false)
  )
})

const taskStore = useTask()
const inputData = reactive(taskStore.getFilter)
const loading = ref(false)
const setFilter = async () => {
  if (!isFormValid.value) {
    return
  }

  try {
    loading.value = true
    await taskStore.setFilter(inputData)
    ElMessage({
      message: '検索条件が適用されました',
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
  <el-form
    ref="formRef"
    :inline="true"
    :model="inputData"
    :rules="rules"
    class="filter-container"
  >
    <el-form-item label="状態">
      <el-radio-group v-model="inputData.status">
        <el-radio :label="-1">全て</el-radio>
        <el-radio :label="TaskStatus.INCOMPLETE">未完</el-radio>
        <el-radio :label="TaskStatus.COMPLETE">完了済み</el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item prop="name" label="タスク名">
      <el-input v-model="inputData.name" @keyup.enter="setFilter" />
    </el-form-item>
    <el-form-item>
      <el-button
        type="primary"
        :loading="loading"
        :disabled="!isFormValid"
        @click="setFilter"
      >
        検索
      </el-button>
    </el-form-item>
  </el-form>
</template>

<style lang="scss">
.filter-container {
  .el-form-item__label {
    color: black;
    font-weight: bold;
  }
}
</style>
