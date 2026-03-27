<template>
  <div class="mx-auto max-w-4xl px-4 py-8">
    <h1 class="mb-8 text-2xl font-bold text-gray-900">{{ isEdit ? 'Edit Client' : 'New Client' }}</h1>

    <form @submit.prevent="handleSubmit" class="space-y-6">
      <fieldset class="space-y-4">
        <legend class="text-lg font-medium text-gray-900">Client Details</legend>
        <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div>
            <label class="block text-sm font-medium text-gray-700">Company Name</label>
            <input v-model="form.name" type="text" required
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Contact Person</label>
            <input v-model="form.contact_person" type="text"
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Email</label>
            <input v-model="form.email" type="email"
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div class="sm:col-span-2">
            <label class="block text-sm font-medium text-gray-700">Address</label>
            <input v-model="form.address" type="text" required
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">VAT Number</label>
            <input v-model="form.vat_number" type="text"
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Registration Number</label>
            <input v-model="form.reg_number" type="text"
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
        </div>
      </fieldset>

      <fieldset class="space-y-4">
        <legend class="text-lg font-medium text-gray-900">Contract Information</legend>
        <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div>
            <label class="block text-sm font-medium text-gray-700">Contract Reference</label>
            <input v-model="form.contract_reference" type="text" maxlength="255"
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Status</label>
            <select v-model="form.status"
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500">
              <option value="active">Active</option>
              <option value="inactive">Inactive</option>
            </select>
          </div>
          <div class="sm:col-span-2">
            <label class="block text-sm font-medium text-gray-700">Contract Notes</label>
            <textarea v-model="form.contract_notes" rows="3" maxlength="2000"
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
        </div>
      </fieldset>

      <p v-if="error" class="text-sm text-red-600">{{ error }}</p>

      <div class="flex flex-col gap-3 sm:flex-row">
        <button type="submit" :disabled="saving"
          class="w-full sm:w-auto rounded-md bg-blue-600 px-6 py-2 font-medium text-white hover:bg-blue-700 disabled:opacity-50">
          {{ saving ? 'Saving...' : (isEdit ? 'Update Client' : 'Create Client') }}
        </button>
        <RouterLink to="/clients"
          class="w-full sm:w-auto rounded-md border border-gray-300 px-6 py-2 text-center font-medium text-gray-700 hover:bg-gray-50">
          Cancel
        </RouterLink>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { RouterLink, useRouter, useRoute } from 'vue-router'
import { fetchClient, createClient, updateClient } from '@/services/clientApi'
import type { ClientInput } from '@/services/clientApi'

const router = useRouter()
const route = useRoute()

const clientId = computed(() => {
  const id = route.params.id
  return id ? Number(id) : null
})
const isEdit = computed(() => !!clientId.value)
const saving = ref(false)
const error = ref('')

const form = reactive<ClientInput>({
  name: '',
  contact_person: '',
  email: '',
  address: '',
  vat_number: '',
  reg_number: '',
  contract_reference: '',
  contract_notes: '',
  status: 'active',
})

async function loadClient() {
  if (!clientId.value) return
  const client = await fetchClient(clientId.value)
  form.name = client.name
  form.contact_person = client.contact_person ?? ''
  form.email = client.email ?? ''
  form.address = client.address
  form.vat_number = client.vat_number ?? ''
  form.reg_number = client.reg_number ?? ''
  form.contract_reference = client.contract_reference ?? ''
  form.contract_notes = client.contract_notes ?? ''
  form.status = client.status
}

async function handleSubmit() {
  saving.value = true
  error.value = ''
  try {
    const data: ClientInput = {
      ...form,
      contact_person: form.contact_person || null,
      email: form.email || null,
      vat_number: form.vat_number || null,
      reg_number: form.reg_number || null,
      contract_reference: form.contract_reference || null,
      contract_notes: form.contract_notes || null,
    }
    if (clientId.value) {
      await updateClient(clientId.value, data)
    } else {
      await createClient(data)
    }
    router.push('/clients')
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Failed to save'
  } finally {
    saving.value = false
  }
}

onMounted(loadClient)
</script>
