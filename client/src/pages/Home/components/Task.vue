<script lang="ts" setup>
import { PropType, ref } from 'vue'
import { Task } from '@/lib/apis'
import { parseDay } from '@/util/day'
import { TaskStatus } from '@/types/taskStatus'

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

const formattedTimeLimit = ref(
  props.task.timeLimit ? parseDay(props.task.timeLimit) : 'なし'
)

const taskStatus = ref(
  props.task.status === TaskStatus.COMPLETE ? '完了済み' : '未完'
)
</script>

<template>
  <div class="task-table">
    <span>{{ task.taskName }}</span>
    <span>{{ formattedTimeLimit }}</span>
    <span>{{ taskStatus }}</span>
    <span class="buttons">
      <el-button
        v-if="task.status === TaskStatus.INCOMPLETE"
        type="success"
        round
      >
        完了
      </el-button>
      <el-button v-else type="danger" round> 削除 </el-button>
      <el-button type="info" round>編集</el-button>
    </span>
  </div>
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
