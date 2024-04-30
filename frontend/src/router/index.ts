import { createRouter, createWebHashHistory } from 'vue-router'
import Body from '@/components/Body.vue'
import Book from '@/components/Book.vue'
import Books from '@/components/Books.vue'
import Login from '@/components/Login.vue'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Body
    },
    {
      path: '/books',
      name: 'BooksComposition',
      meta: {
        requiresAuth: true
      },
      component: Books
    },
    {
      path: '/books/:id',
      name: 'Book',
      component: Book
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    }
  ]
})

router.beforeEach(async (to, from) => {
  const authStore = useAuthStore()
  const authRequired = to.matched.some((record) => record.meta.requiresAuth)

  if (authRequired && !authStore.isAuthenticated()) {
    await authStore.validateToken()

    if (!authStore.isAuthenticated()) {
      authStore.returnUrl = to.fullPath
      return '/login'
    }
  }
})

export default router
