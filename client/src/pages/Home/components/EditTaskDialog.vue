<script lang="ts" setup>
import { AxiosError } from 'axios'
import { PropType, reactive, ref, watchEffect } from 'vue'
import { ElMessage, ElMessageBox, FormInstance } from 'element-plus'
import { TaskStatus } from '@/types/taskStatus'
import { useTask } from '@/store/task'
import { Task, UpdateTask } from '@/lib/apis'
import { showErrorMessage } from '@/util/showErrorMessage'
import { getRules } from '@/util/validate'

const props = defineProps({
  modelValue: {
    type: Boolean,
    required: true
  },
  task: {
    type: Object as PropType<Task>,
    required: true
  },
  index: {
    type: Number,
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

  if (editTask.taskName.length === 0) {
    isFormValid.value = false
    return
  }

  value.validate(isValid =>
    isValid ? (isFormValid.value = true) : (isFormValid.value = false)
  )
})

const editTask = reactive<UpdateTask>({
  taskName: props.task.taskName,
  status: props.task.status,
  timeLimit: props.task.timeLimit
})
const loading = ref(false)
const taskStore = useTask()
const confirmEdit = async () => {
  try {
    loading.value = true
    await taskStore.editTask(props.index, props.task.taskId, editTask)
    ElMessage({
      message: 'タスク内容を更新しました',
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

const confirmDelete = async () => {
  await ElMessageBox.confirm(
    `タスク: ${props.task.taskName}を削除しますか`,
    '確認',
    {
      confirmButtonText: 'はい',
      cancelButtonText: 'いいえ',
      type: 'warning'
    }
  )
    .then(() => deleteTask())
    // eslint-disable-next-line @typescript-eslint/no-empty-function
    .catch(() => {})
}
const deleteTask = async () => {
  try {
    loading.value = true
    await taskStore.deleteTask(props.index, props.task.taskId)
    ElMessage({
      message: 'タスクを削除しました',
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
    title="タスク編集"
    @close="$emit('update:modelValue', false)"
  >
    <el-form
      ref="formRef"
      :model="editTask"
      :rules="rules"
      label-position="top"
    >
      <el-form-item prop="taskName" label="タスク名">
        <el-input v-model="editTask.taskName" />
      </el-form-item>
      <el-form-item prop="status" label="状態">
        <el-radio-group v-model="editTask.status">
          <el-radio :label="TaskStatus.INCOMPLETE">未完</el-radio>
          <el-radio :label="TaskStatus.COMPLETE">完了済み</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item prop="timeLimit" label="期限">
        <el-date-picker
          v-model="editTask.timeLimit"
          type="date"
          format="YYYY/MM/DD"
          value-format="YYYY/MM/DD"
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button type="danger" :loading="loading" @click="confirmDelete">
        削除
      </el-button>
      <el-button
        type="primary"
        :loading="loading"
        :disabled="!isFormValid"
        @click="confirmEdit"
      >
        編集
      </el-button>
    </template>
  </el-dialog>
</template>
