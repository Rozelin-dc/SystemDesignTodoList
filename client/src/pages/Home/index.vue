<script lang="ts" setup>
import { AxiosError } from 'axios'
import { onMounted, ref, watchEffect } from 'vue'
import { ElLoading } from 'element-plus'
import { useTask } from '@/store/task'
import { showErrorMessage } from '@/util/showErrorMessage'
import TaskComponent from './components/Task.vue'
import NewTaskDialog from './components/NewTaskDialog.vue'

const showCreateDialog = ref(false)

const taskStore = useTask()

const loadingEle = ref<HTMLDivElement | undefined>(undefined)
watchEffect(onCleanup => {
  const { value } = loadingEle
  if (!value || !taskStore.getHasNext) {
    return
  }

  const observer = new IntersectionObserver(
    entries => {
      for (const entry of entries) {
        if (!entry.isIntersecting) {
          return
        }
      }

      void taskStore.setTasks(10, taskStore.getTasks.length)
    },
    {
      threshold: 0.5
    }
  )
  observer.observe(value)

  onCleanup(() => observer.disconnect())
})

onMounted(async () => {
  try {
    await taskStore.setTasks(20, 0)

    if (loadingEle.value) {
      ElLoading.service({ target: loadingEle.value })
    }
  } catch (e: any) {
    const err: AxiosError = e
    showErrorMessage(err)
  }
})
</script>

<template>
  <div class="task-list-container">
    <div class="create-task-button">
      <el-button type="primary" @click="showCreateDialog = true" round>
        タスク作成
      </el-button>
    </div>

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
    <div v-if="taskStore.getHasNext" ref="loadingEle" class="loading" />
  </div>
  <new-task-dialog v-model="showCreateDialog" />
</template>

<style lang="scss" scoped>
@import './style.scss';

.task-list-container {
  padding: 10px 5px 0;

  .create-task-button {
    display: flex;
    justify-content: flex-end;
    margin-bottom: 10px;
  }

  .task-table-header {
    font-weight: bold;
  }

  .task-item {
    margin: 10px 0;
  }

  .loading {
    height: 50px;
  }
}
</style>
