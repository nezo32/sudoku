/**
 * router/index.ts
 *
 * Automatic routes for `./src/pages/*.vue`
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router/auto'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes:
  [
    {
      path:'/',
      name:'Главная страница',
      component: () => import('@/views/MainView.vue')
    }
    ,
    {
      path: '/auth',
      name: 'Main',
      component: () => import("@/views/AuthView.vue"),
      redirect:"/auth/login",
      children:[{
        path: "register",
        name: "Регестрация",
        component: () => import("@/components/Auth/Register.vue")
      },
      {
        path: "login",
        name: "Login",
        component: () => import("@/components/Auth/LogIn.vue"),
        
      }]
    },
    {
      path: '/profile',
      name: 'Профиль',
      component: () => import("@/views/GameView.vue")
    },
  ],
})

// Workaround for https://github.com/vitejs/vite/issues/11804
router.onError((err, to) => {
  if (err?.message?.includes?.('Failed to fetch dynamically imported module')) {
    if (!localStorage.getItem('vuetify:dynamic-reload')) {
      console.log('Reloading page to fix dynamic import error')
      localStorage.setItem('vuetify:dynamic-reload', 'true')
      location.assign(to.fullPath)
    } else {
      console.error('Dynamic import error, reloading page did not fix it', err)
    }
  } else {
    console.error(err)
  }
})

router.isReady().then(() => {
  localStorage.removeItem('vuetify:dynamic-reload')
})

export default router
