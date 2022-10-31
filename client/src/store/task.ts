import { defineStore } from 'pinia'
import api, { NewTask, Task, UpdateTask } from '@/lib/apis'

export interface TaskFilter {
  status: -1 | 0 | 1
  name: string
}

export const useTask = defineStore('task', {
  state: (): { tasks: Task[]; hasNext: boolean; filter: TaskFilter } => ({
    tasks: [],
    hasNext: false,
    filter: {
      status: -1,
      name: ''
    }
  }),
  getters: {
    getTasks: state => state.tasks,
    getHasNext: state => state.hasNext,
    getFilter: state => state.filter
  },
  actions: {
    async setFilter(filter: TaskFilter) {
      this.filter = filter
      await this.setTasks(20, 0)
    },
    async setTasks(limit: number, offset: number) {
      const { data } = await api.getTasks(
        limit,
        offset,
        this.filter.status === -1 ? undefined : this.filter.status,
        this.filter.name === '' ? undefined : this.filter.name
      )
      if (offset === 0) {
        this.tasks = data.tasks
      } else {
        this.tasks = this.tasks.concat(data.tasks)
      }
      this.hasNext = data.hasNext
    },
    async editTask(idx: number, id: string, newData: UpdateTask) {
      const { data } = await api.patchTask(id, {
        ...newData,
        timeLimit: newData.timeLimit === '' ? undefined : newData.timeLimit
      })
      this.tasks[idx] = data
    },
    async deleteTask(idx: number, id: string) {
      await api.deleteTask(id)
      this.tasks = this.tasks.filter((_, index) => index != idx)
    },
    async createTask(task: NewTask) {
      const { data } = await api.createTask({
        ...task,
        timeLimit: task.timeLimit === '' ? undefined : task.timeLimit
      })
      this.tasks.unshift(data)
    }
  }
})
