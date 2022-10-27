<script lang="ts" setup>
import { AxiosError } from 'axios'
import { PropType, ref } from 'vue'
import { Task } from '@/lib/apis'
import { parseDay } from '@/util/day'
import { TaskStatus } from '@/types/taskStatus'
import { useTask } from '@/store/task'
import { showErrorMessage } from '@/util/showErrorMessage'
import EditTaskDialog from './EditTaskDialog.vue'

const props = defineProps({
  task: {
    type: Object as PropType<Task>,
    required: true
  },
  index: {
    type: Number,
    required: true
  }
})

const showEditDialog = ref(false)

const loading = ref(false)
const taskStore = useTask()
const taskComplete = async () => {
  try {
    loading.value = true
    await taskStore.editTask(props.index, props.task.taskId, {
      taskName: props.task.taskName,
      status: TaskStatus.COMPLETE,
      timeLimit: props.task.timeLimit
    })
  } catch (e: any) {
    const err: AxiosError = e
    showErrorMessage(err)
  } finally {
    loading.value = false
  }
}

const taskDelete = async () => {
  try {
    loading.value = true
    await taskStore.deleteTask(props.index, props.task.taskId)
  } catch (e: any) {
    const err: AxiosError = e
    showErrorMessage(err)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="task-table">
    <span>{{ task.taskName }}</span>
    <span>{{ task.timeLimit ? parseDay(task.timeLimit) : 'なし' }}</span>
    <span>{{ task.status === TaskStatus.COMPLETE ? '完了済み' : '未完' }}</span>
    <span class="buttons">
      <el-button
        v-if="task.status === TaskStatus.INCOMPLETE"
        type="success"
        :loading="loading"
        @click="taskComplete"
        round
      >
        完了
      </el-button>
      <el-button
        v-else
        type="danger"
        :loading="loading"
        @click="taskDelete"
        round
      >
        削除
      </el-button>
      <el-button type="info" @click="showEditDialog = true" round>
        編集
      </el-button>
    </span>
  </div>
  <edit-task-dialog v-model="showEditDialog" :task="task" :index="index" />
</template>

<style lang="scss">
.task-table {
  .el-button {
    margin: 0;
  }
}
</style>

<style lang="scss" scoped>
@import '../style.scss';

.buttons {
  display: grid;
  grid-template-columns: 1fr 1fr;
  column-gap: 10px;
}
</style>
