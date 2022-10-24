import { defineStore } from 'pinia'
import api, { Task, UpdateTask } from '@/lib/apis'

export const useTask = defineStore('task', {
  state: (): { tasks: Task[] } => ({ tasks: [] }),
  getters: {
    getTasks: state => state.tasks
  },
  actions: {
    async setTasks(limit: number, offset: number) {
      const { data } = await api.getTasks(limit, offset)
      this.tasks = data.tasks
      return data.hasNext
    },
    async editTask(idx: number, id: string, newData: UpdateTask) {
      const { data } = await api.patchTask(id, newData)
      this.tasks[idx] = data
    },
    async deleteTask(idx: number, id: string) {
      await api.deleteTask(id)
      this.tasks = this.tasks.filter((_, index) => index != idx)
    }
  }
})
