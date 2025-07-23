import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import DashboardView from '../pages/DashboardPage.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Dashboard',
    component: DashboardView,
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router