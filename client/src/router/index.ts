import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import { nextTick } from 'vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: "Sudoku Solvers' Playground",
      component: HomeView
    },
    {
      path: '/authentication',
      name: 'Authentication',
      component: () => import('@/components/bases/AuthBase.vue'),
      children: [
        {
          path: 'login',
          name: 'Sign-in',
          component: () => import('@/views/LoginView.vue')
        },
        {
          path: 'register',
          name: 'Sign-up',
          component: () => import('@/views/RegisterView.vue')
        }
      ]
    }
  ]
})

router.beforeEach((to) => {
  if (to.fullPath == '/') return
  nextTick(() => {
    document.title = to.name?.toString() + ' — SuSo Playground'
  })
})

export default router
