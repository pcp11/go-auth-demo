import { createRouter, createWebHashHistory } from 'vue-router'
import Body from '@/components/Body.vue'
import Book from '@/components/Book.vue'
import Books from '@/components/Books.vue'
import Login from '@/components/Login.vue'

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

export default router