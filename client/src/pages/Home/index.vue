<script lang="ts" setup>
import { AxiosError } from 'axios'
import { onMounted } from 'vue'
import { useTask } from '@/store/task'
import { showErrorMessage } from '@/util/showErrorMessage'
import TaskComponent from './components/Task.vue'

const taskStore = useTask()

onMounted(async () => {
  try {
    await taskStore.setTasks(20, 0)
  } catch (e: any) {
    const err: AxiosError = e
    showErrorMessage(err)
  }
})
</script>

<template>
  <div class="task-list-container">
    <div class="task-table task-table-header">
      <span>タスク名</span>
      <span>期限</span>
      <span>状態</span>
    </div>
    <div
      v-for="(task, idx) in taskStore.getTasks"
      :key="task.taskId"
      class="task-item"
    >
      <task-component :task="task" :index="idx" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
@import './style.scss';

.task-list-container {
  padding: 10px 5px 0;

  .task-table-header {
    font-weight: bold;
  }

  .task-item {
    margin: 10px 0;
  }
}
</style>
