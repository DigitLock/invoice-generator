<template>
  <div class="mx-auto max-w-7xl px-4 py-8">
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-2xl font-bold text-gray-900">Clients</h1>
      <RouterLink to="/clients/new"
        class="rounded-md bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700">
        Add Client
      </RouterLink>
    </div>

    <div class="mb-4 flex gap-2">
      <button v-for="opt in statusOptions" :key="opt.value"
        @click="statusFilter = opt.value"
        :class="[
          'rounded-full px-3 py-1 text-sm font-medium',
          statusFilter === opt.value
            ? 'bg-blue-600 text-white'
            : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
        ]">
        {{ opt.label }}
      </button>
    </div>

    <div v-if="loading" class="text-sm text-gray-500">Loading...</div>
    <div v-else-if="clients.length === 0" class="rounded-lg border border-gray-200 bg-white p-8 text-center text-gray-500">
      No clients found.
    </div>
    <div v-else class="space-y-3">
      <div v-for="client in clients" :key="client.id"
        class="flex items-center justify-between rounded-lg border border-gray-200 bg-white p-4">
        <div>
          <div class="flex items-center gap-2">
            <span class="font-medium text-gray-900">{{ client.name }}</span>
            <span :class="client.status === 'active' ? 'bg-green-100 text-green-700' : 'bg-gray-100 text-gray-500'"
              class="rounded-full px-2 py-0.5 text-xs font-medium">
              {{ client.status }}
            </span>
          </div>
          <div class="text-sm text-gray-500">
            {{ client.address }}
            <span v-if="client.contact_person"> | {{ client.contact_person }}</span>
          </div>
          <div v-if="client.contract_reference" class="text-sm text-gray-400">
            Contract: {{ client.contract_reference }}
          </div>
        </div>
        <div class="flex gap-3">
          <RouterLink :to="`/clients/${client.id}/edit`" class="text-sm text-blue-600 hover:text-blue-800">
            Edit
          </RouterLink>
          <button @click="handleDelete(client)" class="text-sm text-red-600 hover:text-red-800">
            Delete
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import type { ClientResponse } from '@/types/api'
import { fetchClients, deleteClient } from '@/services/clientApi'

const statusOptions = [
  { value: '', label: 'All' },
  { value: 'active', label: 'Active' },
  { value: 'inactive', label: 'Inactive' },
]

const clients = ref<ClientResponse[]>([])
const loading = ref(true)
const statusFilter = ref('')

async function load() {
  loading.value = true
  clients.value = await fetchClients(statusFilter.value || undefined)
  loading.value = false
}

async function handleDelete(client: ClientResponse) {
  if (!confirm(`Delete "${client.name}"? This cannot be undone.`)) return
  try {
    await deleteClient(client.id)
    await load()
  } catch (e) {
    alert(e instanceof Error ? e.message : 'Failed to delete')
  }
}

watch(statusFilter, load)
onMounted(load)
</script>
