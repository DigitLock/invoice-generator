<template>
  <div class="mx-auto max-w-7xl px-4 py-8">
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-2xl font-bold text-gray-900">Companies</h1>
      <RouterLink to="/companies/new"
        class="rounded-md bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700">
        Add Company
      </RouterLink>
    </div>

    <div v-if="loading" class="text-sm text-gray-500">Loading...</div>
    <div v-else-if="companies.length === 0" class="rounded-lg border border-gray-200 bg-white p-8 text-center text-gray-500">
      No companies yet. Add your first company to start creating invoices.
    </div>
    <div v-else class="space-y-3">
      <div v-for="company in companies" :key="company.id"
        class="flex items-center justify-between rounded-lg border border-gray-200 bg-white p-4">
        <div>
          <div class="font-medium text-gray-900">{{ company.name }}</div>
          <div class="text-sm text-gray-500">{{ company.contact_person }} | {{ company.address }}</div>
          <div v-if="company.vat_number" class="text-sm text-gray-400">VAT: {{ company.vat_number }}</div>
        </div>
        <div class="flex gap-3">
          <RouterLink :to="`/companies/${company.id}/edit`" class="text-sm text-blue-600 hover:text-blue-800">
            Edit
          </RouterLink>
          <button @click="handleDelete(company)" class="text-sm text-red-600 hover:text-red-800">
            Delete
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import type { CompanyResponse } from '@/types/api'
import { fetchCompanies, deleteCompany } from '@/services/companyApi'

const companies = ref<CompanyResponse[]>([])
const loading = ref(true)

async function load() {
  loading.value = true
  companies.value = await fetchCompanies()
  loading.value = false
}

async function handleDelete(company: CompanyResponse) {
  if (!confirm(`Delete "${company.name}"? This cannot be undone.`)) return
  try {
    await deleteCompany(company.id)
    await load()
  } catch (e) {
    alert(e instanceof Error ? e.message : 'Failed to delete')
  }
}

onMounted(load)
</script>
