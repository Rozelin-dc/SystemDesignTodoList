import { defineStore } from 'pinia'
import api, { NewTask, Task, UpdateTask } from '@/lib/apis'

export const useTask = defineStore('task', {
  state: (): { tasks: Task[]; hasNext: boolean } => ({
    tasks: [],
    hasNext: true
  }),
  getters: {
    getTasks: state => state.tasks,
    getHasNext: state => state.hasNext
  },
  actions: {
    async setTasks(limit: number, offset: number) {
      const { data } = await api.getTasks(limit, offset)
      if (offset === 0) {
        this.tasks = data.tasks
      } else {
        this.tasks = this.tasks.concat(data.tasks)
      }
      this.hasNext = data.hasNext
    },
    async editTask(idx: number, id: string, newData: UpdateTask) {
      const { data } = await api.patchTask(id, newData)
      this.tasks[idx] = data
    },
    async deleteTask(idx: number, id: string) {
      await api.deleteTask(id)
      this.tasks = this.tasks.filter((_, index) => index != idx)
    },
    async createTask(task: NewTask) {
      const { data } = await api.createTask(task)
      this.tasks.unshift(data)
    }
  }
})
