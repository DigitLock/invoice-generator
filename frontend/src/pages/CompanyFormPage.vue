<template>
  <div class="mx-auto max-w-4xl px-4 py-8">
    <h1 class="mb-8 text-2xl font-bold text-gray-900">{{ isEdit ? 'Edit Company' : 'New Company' }}</h1>

    <form @submit.prevent="handleSubmit" class="space-y-6">
      <fieldset class="space-y-4">
        <legend class="text-lg font-medium text-gray-900">Company Details</legend>
        <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div>
            <label class="block text-sm font-medium text-gray-700">Company Name</label>
            <input v-model="form.name" type="text" required
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Contact Person</label>
            <input v-model="form.contact_person" type="text" required
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div class="sm:col-span-2">
            <label class="block text-sm font-medium text-gray-700">Address</label>
            <input v-model="form.address" type="text" required
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Phone</label>
            <input v-model="form.phone" type="text"
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

      <p v-if="error" class="text-sm text-red-600">{{ error }}</p>

      <div class="flex flex-col gap-3 sm:flex-row">
        <button type="submit" :disabled="saving"
          class="w-full sm:w-auto rounded-md bg-blue-600 px-6 py-2 font-medium text-white hover:bg-blue-700 disabled:opacity-50">
          {{ saving ? 'Saving...' : (isEdit ? 'Update Company' : 'Create Company') }}
        </button>
        <RouterLink to="/companies"
          class="w-full sm:w-auto rounded-md border border-gray-300 px-6 py-2 text-center font-medium text-gray-700 hover:bg-gray-50">
          Cancel
        </RouterLink>
      </div>
    </form>

    <div v-if="isEdit && companyId" class="mt-10 border-t border-gray-200 pt-8">
      <BankAccountList :company-id="companyId" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { RouterLink, useRouter, useRoute } from 'vue-router'
import { fetchCompany, createCompany, updateCompany } from '@/services/companyApi'
import type { CompanyInput } from '@/services/companyApi'
import BankAccountList from '@/components/company/BankAccountList.vue'

const router = useRouter()
const route = useRoute()

const companyId = computed(() => {
  const id = route.params.id
  return id ? Number(id) : null
})
const isEdit = computed(() => !!companyId.value)
const saving = ref(false)
const error = ref('')

const form = reactive<CompanyInput>({
  name: '',
  contact_person: '',
  address: '',
  phone: '',
  vat_number: '',
  reg_number: '',
})

async function loadCompany() {
  if (!companyId.value) return
  const company = await fetchCompany(companyId.value)
  form.name = company.name
  form.contact_person = company.contact_person
  form.address = company.address
  form.phone = company.phone ?? ''
  form.vat_number = company.vat_number ?? ''
  form.reg_number = company.reg_number ?? ''
}

async function handleSubmit() {
  saving.value = true
  error.value = ''
  try {
    const data: CompanyInput = {
      ...form,
      phone: form.phone || null,
      vat_number: form.vat_number || null,
      reg_number: form.reg_number || null,
    }
    if (companyId.value) {
      await updateCompany(companyId.value, data)
    } else {
      await createCompany(data)
    }
    router.push('/companies')
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Failed to save'
  } finally {
    saving.value = false
  }
}

onMounted(loadCompany)
</script>
