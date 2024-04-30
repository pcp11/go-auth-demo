import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { Book } from '@/services/api'
import { axiosInstance } from '@/services/api'

const API_URL: string = import.meta.env.VITE_API_URL

export const useBookStore = defineStore('book', () => {
  const books = ref<Book[]>([])
  const isLoading = ref<Boolean>(false)

  const fetchBooks = async () => {
    isLoading.value = true

    const response = await axiosInstance.get(`${API_URL}/books`)

    books.value = response.data.data
    isLoading.value = false
  }

  return {
    books,
    isLoading,
    fetchBooks
  }
})
