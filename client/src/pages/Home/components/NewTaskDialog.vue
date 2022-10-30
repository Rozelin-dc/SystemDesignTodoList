<script lang="ts" setup>
import { AxiosError } from 'axios'
import { reactive, ref, watchEffect } from 'vue'
import { ElMessage, FormInstance } from 'element-plus'
import { useTask } from '@/store/task'
import { NewTask } from '@/lib/apis'
import { showErrorMessage } from '@/util/showErrorMessage'
import { getRules } from '@/util/validate'

defineProps({
  modelValue: {
    type: Boolean,
    required: true
  }
})

const emits = defineEmits(['update:modelValue'])

const formRef = ref<FormInstance>()
const rules = reactive(getRules(['taskName']))
const isFormValid = ref(false)
watchEffect(() => {
  const { value } = formRef
  if (!value) {
    isFormValid.value = false
    return
  }

  if (newTask.taskName.length === 0) {
    isFormValid.value = false
    return
  }

  value.validate(isValid =>
    isValid ? (isFormValid.value = true) : (isFormValid.value = false)
  )
})

const newTask = reactive<NewTask>({
  taskName: '',
  timeLimit: ''
})
const loading = ref(false)
const taskStore = useTask()
const confirmCreate = async () => {
  try {
    loading.value = true
    await taskStore.createTask(newTask)
    ElMessage({
      message: 'タスクを作成しました',
      type: 'success'
    })
    emits('update:modelValue', false)
  } catch (e: any) {
    const err: AxiosError = e
    showErrorMessage(err)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <el-dialog
    :model-value="modelValue"
    title="タスク作成"
    @close="$emit('update:modelValue', false)"
  >
    <el-form ref="formRef" :model="newTask" :rules="rules" label-position="top">
      <el-form-item prop="taskName" label="タスク名">
        <el-input v-model="newTask.taskName" />
      </el-form-item>
      <el-form-item prop="timeLimit" label="期限">
        <el-date-picker
          v-model="newTask.timeLimit"
          type="date"
          format="YYYY/MM/DD"
          value-format="YYYY/MM/DD"
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        :loading="loading"
        :disabled="!isFormValid"
        @click="confirmCreate"
      >
        確定
      </el-button>
    </template>
  </el-dialog>
</template>
