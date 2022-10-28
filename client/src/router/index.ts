import { AxiosError } from 'axios'
import {
  createRouter,
  createWebHistory,
  RouteLocation,
  RouteRecordRaw
} from 'vue-router'
import Layout from '@/layout/index.vue'
import { useMe } from '@/store/me'
import { showErrorMessage } from '@/util/showErrorMessage'

interface RouteMeta {
  title?: string
  isPublic?: boolean
}

type IRouteRecordRaw = RouteRecordRaw & {
  meta?: RouteMeta
  children?: IRouteRecordRaw[]
}

type IRoute = Omit<RouteLocation, 'meta'> & {
  meta?: RouteMeta
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
  },
  {
    path: 'change-password',
    name: 'ChangePassword',
    component: () => import('@/pages/ChangePassword/index.vue'),
    meta: { title: 'パスワード変更' }
  },
  {
    path: 'delete-account',
    name: 'DeleteAccount',
    component: () => import('@/pages/DeleteAccount/index.vue'),
    meta: { title: 'アカウント削除' }
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
    component: () => import('@/pages/Login/index.vue')
  },
  {
    path: '/create-account',
    name: 'CreateAccount',
    component: () => import('@/pages/CreateAccount/index.vue')
  }
].map((route: IRouteRecordRaw) => {
  return {
    ...route,
    meta: { isPublic: true }
  }
})

const router = createRouter({
  history: createWebHistory(),
  routes: constantRouts.concat(publicRoutes)
})

router.beforeEach(async (to: IRoute, _, next) => {
  if (to.meta && to.meta.isPublic) {
    next()
  }

  const meStore = useMe()
  if (meStore.getMe) {
    next()
  }

  try {
    await meStore.setMe()
    next()
  } catch (e: any) {
    const err: AxiosError = e
    showErrorMessage(err)
    next({ name: 'Login' })
  }
})

export default router
