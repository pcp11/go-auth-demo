import { ref } from 'vue'
import { defineStore } from 'pinia'
import axios from 'axios'

// Set withCredentials to true for all requests
//axios.defaults.withCredentials = true;

const API_URL: string = import.meta.env.VITE_API_URL

export const useAuthStore = defineStore('auth', () => {
  const user = ref('')

  const login = async (email: string, password: string) => {
    await axios.post(API_URL + '/login', {
      email: email,
      password: password
    },  {
      withCredentials: true // Necessary to receive cookies
    });
    const response = await validateToken();
    console.log(response);
  }

  const validateToken = () => {
    return axios.post(API_URL + '/auth')
  }

  return { user, login, validateToken }
})
