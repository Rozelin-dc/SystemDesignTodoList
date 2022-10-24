import { defineStore } from 'pinia'
import api, { UserWithoutPass } from '@/lib/apis'

export const useMe = defineStore('me', {
  state: (): { me: UserWithoutPass | undefined } => ({ me: undefined }),
  getters: {
    getMe: state => state.me
  },
  actions: {
    async setMe() {
      const { data } = await api.getUserMe()
      this.me = data
    }
  }
})
