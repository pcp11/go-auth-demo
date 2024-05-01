import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { Book } from '@/services/api'
import { axiosInstance } from '@/services/api'
import axios from 'axios'

const API_URL: string = import.meta.env.VITE_API_URL

export const useBookStore = defineStore('book', () => {
  const books = ref<Book[]>([])
  const isLoading = ref<Boolean>(false)

  const fetchBooks = async () => {
    isLoading.value = true

    try {
      const response = await axiosInstance.get(`${API_URL}/books`)

      books.value = response.data.data
      isLoading.value = false
    } catch (error) {
      if (axios.isAxiosError(error)) {
        console.log(error.message)
      } else {
        console.log(error)
      }
    }
  }

  return {
    books,
    isLoading,
    fetchBooks
  }
})
