import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { login as loginApi } from '@/services/authApi'

interface User {
  id: number
  familyId: number
  email: string
  name: string
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('jwt_token'))
  const user = ref<User | null>(loadUser())

  const isAuthenticated = computed(() => !!token.value)

  function loadUser(): User | null {
    const stored = localStorage.getItem('user')
    if (!stored) return null
    try {
      return JSON.parse(stored)
    } catch {
      return null
    }
  }

  async function login(email: string, password: string) {
    const response = await loginApi({ email, password })

    token.value = response.token
    user.value = {
      id: response.user.id,
      familyId: response.user.family_id,
      email: response.user.email,
      name: response.user.name,
    }

    localStorage.setItem('jwt_token', response.token)
    localStorage.setItem('user', JSON.stringify(user.value))
  }

  function logout() {
    token.value = null
    user.value = null
    localStorage.removeItem('jwt_token')
    localStorage.removeItem('user')
  }

  return {
    token,
    user,
    isAuthenticated,
    login,
    logout,
  }
})
