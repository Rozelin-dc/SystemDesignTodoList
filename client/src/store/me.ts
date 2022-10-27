import { defineStore } from 'pinia'
import api, { UserWithoutPass, UpdateUser } from '@/lib/apis'

export const useMe = defineStore('me', {
  state: (): { me: UserWithoutPass | undefined } => ({ me: undefined }),
  getters: {
    getMe: state => state.me
  },
  actions: {
    async setMe() {
      const { data } = await api.getUserMe()
      this.me = data
    },
    async logout() {
      await api.postLogout()
      this.me = undefined
    },
    async changeMeData(newData: UpdateUser) {
      if (!this.me) {
        throw new Error('not logged in')
      }
      const { data } = await api.patchUser(this.me.userId, newData)
      this.me = data
    },
    async deleteMe(pass: string) {
      if (!this.me) {
        throw new Error('not logged in')
      }
      await api.deleteUser(this.me.userId, pass)
      this.me = undefined
    }
  }
})
