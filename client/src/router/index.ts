import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Layout from '@/layout/index.vue'

interface RouteMeta {
  title: string
  isPublic?: boolean
}

type IRouteRecordRaw = RouteRecordRaw & {
  meta?: RouteMeta
  children?: IRouteRecordRaw[]
}

export const sidebarRoutes: IRouteRecordRaw[] = [
  {
    path: '',
    name: 'Home',
    component: () => import('@/pages/Home/index.vue'),
    meta: { title: 'タスク一覧' }
  },
  {
    path: 'change-user-name',
    name: 'ChangeUserName',
    component: () => import('@/pages/ChangeUserName/index.vue'),
    meta: { title: 'ユーザー名変更' }
  }
]

const constantRouts: IRouteRecordRaw[] = [
  {
    path: '/',
    component: Layout,
    children: sidebarRoutes
  }
]

const publicRoutes: IRouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/pages/Home/index.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes: constantRouts.concat(publicRoutes)
})

export default router
