import axios from 'axios'
import router from '@/router'

export interface Book {
  id: number
  title: string
  authors: Array<Person>
  subjects: Array<string>
  formats: { [key: string]: string }
  created_at: string
  updated_at: string
  deleted_at: string
}

export interface Person {
  id: number
  name: string
  birth_year: number
  death_year: number
  created_at: string
  updated_at: string
  deleted_at: string
}

export interface User {
  id: number
  name: string
  email: string
  created_at: string
  updated_at: string
  deleted_at: string
}

export const axiosInstance = axios.create({
  withCredentials: true
})

// axiosInstance.interceptors.response.use(
//   response => response,
//   (error) => {
//     if (error.response.status === 401) {
//       router.push({ path: 'login' })
//       return
//     }
//     return Promise.reject(error);
//   }
// )
