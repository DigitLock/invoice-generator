<template>
  <div class="space-y-4">
    <div class="flex items-center justify-between">
      <h3 class="text-lg font-medium text-gray-900">Bank Accounts</h3>
      <button v-if="!showForm" type="button" @click="startAdd"
        class="text-sm text-blue-600 hover:text-blue-800">+ Add Bank Account</button>
    </div>

    <div v-if="accounts.length === 0 && !showForm" class="text-sm text-gray-500">
      No bank accounts yet.
    </div>

    <div v-for="account in accounts" :key="account.id"
      class="flex items-start justify-between rounded-lg border border-gray-200 bg-white p-4">
      <div class="space-y-1 text-sm">
        <div class="font-medium text-gray-900">
          {{ account.bank_name }}
          <span v-if="account.is_default"
            class="ml-2 inline-block rounded-full bg-blue-100 text-blue-700 px-2 py-0.5 text-xs">Default</span>
        </div>
        <div class="text-gray-600">{{ account.account_holder }}</div>
        <div class="text-gray-500">IBAN: {{ account.iban }}</div>
        <div class="text-gray-500">SWIFT: {{ account.swift }} | {{ account.currency }}</div>
        <div v-if="account.bank_address" class="text-gray-400">{{ account.bank_address }}</div>
      </div>
      <div class="flex gap-2">
        <button @click="startEdit(account)" class="text-sm text-blue-600 hover:text-blue-800">Edit</button>
        <button @click="handleDelete(account)" class="text-sm text-red-600 hover:text-red-800">Delete</button>
      </div>
    </div>

    <BankAccountForm
      v-if="showForm"
      :account="editingAccount ?? undefined"
      @save="handleSave"
      @cancel="cancelForm"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { BankAccountResponse } from '@/types/api'
import type { BankAccountInput } from '@/services/bankAccountApi'
import { fetchBankAccounts, createBankAccount, updateBankAccount, deleteBankAccount } from '@/services/bankAccountApi'
import BankAccountForm from './BankAccountForm.vue'

const props = defineProps<{
  companyId: number
}>()

const accounts = ref<BankAccountResponse[]>([])
const showForm = ref(false)
const editingAccount = ref<BankAccountResponse | null>(null)

async function loadAccounts() {
  accounts.value = await fetchBankAccounts(props.companyId)
}

function startAdd() {
  editingAccount.value = null
  showForm.value = true
}

function startEdit(account: BankAccountResponse) {
  editingAccount.value = account
  showForm.value = true
}

function cancelForm() {
  showForm.value = false
  editingAccount.value = null
}

async function handleSave(data: BankAccountInput) {
  if (editingAccount.value) {
    await updateBankAccount(editingAccount.value.id, data)
  } else {
    await createBankAccount(props.companyId, data)
  }
  cancelForm()
  await loadAccounts()
}

async function handleDelete(account: BankAccountResponse) {
  if (!confirm(`Delete bank account "${account.bank_name}"?`)) return
  await deleteBankAccount(account.id)
  await loadAccounts()
}

onMounted(loadAccounts)
</script>
