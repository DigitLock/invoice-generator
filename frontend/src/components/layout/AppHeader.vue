<template>
  <header class="border-b border-gray-200 bg-white">
    <div class="mx-auto max-w-7xl px-4 py-4 flex items-center justify-between">
      <RouterLink :to="isAuthenticated ? '/dashboard' : '/'" class="text-xl font-semibold text-gray-900">
        Invoice Generator
      </RouterLink>
      <nav class="flex items-center gap-4">
        <template v-if="isAuthenticated">
          <RouterLink to="/dashboard" class="text-sm text-gray-600 hover:text-gray-900">Dashboard</RouterLink>
          <RouterLink to="/companies" class="text-sm text-gray-600 hover:text-gray-900">Companies</RouterLink>
          <RouterLink to="/clients" class="text-sm text-gray-600 hover:text-gray-900">Clients</RouterLink>
          <RouterLink to="/invoices" class="text-sm text-gray-600 hover:text-gray-900">Invoices</RouterLink>
          <span class="text-sm text-gray-500">{{ user?.name }}</span>
          <button @click="handleLogout" class="text-sm text-red-600 hover:text-red-800">
            Log out
          </button>
        </template>
        <RouterLink v-else to="/login" class="text-sm text-blue-600 hover:text-blue-800">
          Sign In
        </RouterLink>
      </nav>
    </div>
  </header>
</template>

<script setup lang="ts">
import { RouterLink, useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()
const { isAuthenticated, user } = storeToRefs(authStore)

function handleLogout() {
  authStore.logout()
  router.push('/')
}
</script>
