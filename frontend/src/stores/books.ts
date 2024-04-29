import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { Book } from 'src/services/api'
import axios from 'axios'

const API_URL: string = import.meta.env.VITE_API_URL

export const useBookStore = defineStore('book', () => {
  const books = ref<Book[]>([])
  const isLoading = ref<Boolean>(false)

  const fetchBooks = () => {
    isLoading.value = true

    return axios.get(`${API_URL}/books`)
      .then((response) => {
          books.value = response.data.data
          isLoading.value = false
      })
  }

  return {
    books,
    isLoading,
    fetchBooks
  }
})
