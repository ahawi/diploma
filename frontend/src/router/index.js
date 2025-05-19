import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/pages/HomePage.vue'
import AnimalsPage from '@/pages/AnimalsPage.vue'

const routes = [
  { path: '/', component: HomePage },
  { path: '/animals', component: AnimalsPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
