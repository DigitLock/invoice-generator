<template>
  <form @submit.prevent="handleSubmit" class="space-y-4 rounded-lg border border-gray-200 bg-gray-50 p-4">
    <h4 class="text-sm font-medium text-gray-900">{{ isEdit ? 'Edit Bank Account' : 'Add Bank Account' }}</h4>
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
      <div>
        <label class="block text-sm font-medium text-gray-700">Bank Name</label>
        <input v-model="form.bank_name" type="text" required
          class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700">Bank Address</label>
        <input v-model="form.bank_address" type="text" required
          class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700">Account Holder</label>
        <input v-model="form.account_holder" type="text" required
          class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700">IBAN</label>
        <input v-model="form.iban" type="text" required
          class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700">SWIFT / BIC</label>
        <input v-model="form.swift" type="text" required
          class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700">Currency</label>
        <input v-model="form.currency" type="text" required maxlength="3" placeholder="EUR"
          class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500 uppercase" />
      </div>
      <div class="flex items-center gap-2 sm:col-span-2">
        <input v-model="form.is_default" type="checkbox" id="is_default"
          class="rounded border-gray-300 text-blue-600 focus:ring-blue-500" />
        <label for="is_default" class="text-sm text-gray-700">Set as default account</label>
      </div>
    </div>
    <p v-if="error" class="text-sm text-red-600">{{ error }}</p>
    <div class="flex gap-2">
      <button type="submit" :disabled="saving"
        class="rounded-md bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700 disabled:opacity-50">
        {{ saving ? 'Saving...' : (isEdit ? 'Update' : 'Add') }}
      </button>
      <button type="button" @click="$emit('cancel')"
        class="rounded-md border border-gray-300 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">
        Cancel
      </button>
    </div>
  </form>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import type { BankAccountResponse } from '@/types/api'
import type { BankAccountInput } from '@/services/bankAccountApi'

const props = defineProps<{
  account?: BankAccountResponse
}>()

const emit = defineEmits<{
  save: [data: BankAccountInput]
  cancel: []
}>()

const isEdit = !!props.account
const saving = ref(false)
const error = ref('')

const form = reactive<BankAccountInput>({
  bank_name: props.account?.bank_name ?? '',
  bank_address: props.account?.bank_address ?? '',
  account_holder: props.account?.account_holder ?? '',
  iban: props.account?.iban ?? '',
  swift: props.account?.swift ?? '',
  currency: props.account?.currency ?? 'EUR',
  is_default: props.account?.is_default ?? false,
})

function handleSubmit() {
  emit('save', { ...form, currency: form.currency.toUpperCase() })
}
</script>
