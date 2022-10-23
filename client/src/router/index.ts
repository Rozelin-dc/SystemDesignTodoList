import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

interface RouteMeta {
  title: string
}

type IRouteRecordRaw = RouteRecordRaw & {
  meta?: RouteMeta
  children?: IRouteRecordRaw[]
}

const sideBarRoutes: IRouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('@/components/HelloWorld.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes: sideBarRoutes
})

export default router
