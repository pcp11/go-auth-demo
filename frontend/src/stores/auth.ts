import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { User } from '@/services/api'
import { axiosInstance } from '@/services/api'
import router from '@/router'
import axios from 'axios'

const API_URL: string = import.meta.env.VITE_API_URL

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>()
  const returnUrl = ref<string>()

  const login = async (email: string, password: string) => {
    await axiosInstance.post(API_URL + '/login', { email: email, password: password })
    await validateToken()

    router.push({ path: returnUrl.value || '/' })
  }

  const logout = async () => {
    await axiosInstance.post(API_URL + '/logout')
    user.value = null
  }

  const validateToken = async () => {
    try {
      const response = await axiosInstance.post(API_URL + '/auth')
      user.value = response.data.user
    } catch (error) {
      if (axios.isAxiosError(error)) {
        console.log(error.message)
      } else {
        console.log(error)
      }
    }
  }

  const isAuthenticated = () => !!user.value

  return { user, returnUrl, login, logout, validateToken, isAuthenticated }
})
