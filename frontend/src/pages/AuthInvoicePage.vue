<template>
  <div class="mx-auto max-w-4xl px-4 py-8">
    <h1 class="mb-8 text-2xl font-bold text-gray-900">{{ isEdit ? 'Edit Invoice' : 'New Invoice' }}</h1>

    <div v-if="loadingSetup" class="text-sm text-gray-500">Loading...</div>
    <form v-else @submit.prevent="handleSubmit" class="space-y-8">
      <!-- Entity selectors (authorized mode) -->
      <fieldset class="space-y-4">
        <legend class="text-lg font-medium text-gray-900">Invoice Setup</legend>
        <div class="grid grid-cols-1 gap-4 sm:grid-cols-3">
          <div>
            <label class="block text-sm font-medium text-gray-700">Company (Seller)</label>
            <select v-model="selectedCompanyId" @change="onCompanyChange" required
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500">
              <option :value="0" disabled>Select company...</option>
              <option v-for="c in companies" :key="c.id" :value="c.id">{{ c.name }}</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Client (Buyer)</label>
            <select v-model="selectedClientId" required
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500">
              <option :value="0" disabled>Select client...</option>
              <option v-for="cl in activeClients" :key="cl.id" :value="cl.id">{{ cl.name }}</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Bank Account</label>
            <select v-model="selectedBankAccountId" required :disabled="bankAccounts.length === 0"
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500 disabled:bg-gray-100">
              <option :value="0" disabled>{{ bankAccounts.length ? 'Select account...' : 'Select company first' }}</option>
              <option v-for="ba in bankAccounts" :key="ba.id" :value="ba.id">
                {{ ba.bank_name }} ({{ ba.currency }}) {{ ba.is_default ? '- Default' : '' }}
              </option>
            </select>
          </div>
        </div>
      </fieldset>

      <!-- Invoice header fields -->
      <fieldset class="space-y-4">
        <legend class="text-lg font-medium text-gray-900">Invoice Details</legend>
        <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
          <div v-if="isEdit">
            <label class="block text-sm font-medium text-gray-700">Invoice Number</label>
            <input v-model="form.invoice_number" type="text" maxlength="50"
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div v-else>
            <label class="block text-sm font-medium text-gray-700">Invoice Number</label>
            <input value="Auto-generated" disabled
              class="mt-1 block w-full rounded-md border border-gray-200 bg-gray-50 px-3 py-2 text-gray-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Issue Date</label>
            <input v-model="form.issue_date" type="date" required
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Due Date</label>
            <input v-model="form.due_date" type="date" required
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Currency</label>
            <select v-model="form.currency"
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500">
              <option value="EUR">EUR</option>
              <option value="RSD">RSD</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">VAT Rate (%)</label>
            <input v-model="form.vat_rate" type="text"
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Contract Reference</label>
            <input v-model="form.contract_reference" type="text" maxlength="255"
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">External Reference</label>
            <input v-model="form.external_reference" type="text" maxlength="255"
              class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
        </div>
      </fieldset>

      <!-- Line items -->
      <fieldset class="space-y-4">
        <legend class="text-lg font-medium text-gray-900">Line Items</legend>
        <div class="-mx-4 overflow-x-auto px-4 sm:mx-0 sm:px-0">
        <table class="w-full min-w-[500px] text-sm">
          <thead>
            <tr class="border-b text-left text-gray-700">
              <th class="pb-2">#</th>
              <th class="pb-2">Description</th>
              <th class="pb-2 w-28">Quantity</th>
              <th class="pb-2 w-32">Unit Price</th>
              <th class="pb-2 w-32">Total</th>
              <th class="pb-2 w-10"></th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in items" :key="index" class="border-b">
              <td class="py-2 pr-2 text-gray-500">{{ index + 1 }}</td>
              <td class="py-2 pr-2">
                <input v-model="item.description" type="text" maxlength="500"
                  class="w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
              </td>
              <td class="py-2 pr-2">
                <input v-model="item.quantity" type="text"
                  class="w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
              </td>
              <td class="py-2 pr-2">
                <input v-model="item.unit_price" type="text"
                  class="w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
              </td>
              <td class="py-2 pr-2 text-right">{{ calcItemTotal(item) }}</td>
              <td class="py-2">
                <button v-if="items.length > 1" type="button" @click="items.splice(index, 1)"
                  class="text-red-500 hover:text-red-700">&times;</button>
              </td>
            </tr>
          </tbody>
        </table>
        </div>
        <button v-if="items.length < 10" type="button" @click="addItem"
          class="text-sm text-blue-600 hover:text-blue-800">+ Add Item</button>
      </fieldset>

      <!-- Totals -->
      <div class="flex justify-end">
        <dl class="w-64 space-y-2 text-sm">
          <div class="flex justify-between">
            <dt class="text-gray-600">Subtotal</dt>
            <dd>{{ subtotal }} {{ form.currency }}</dd>
          </div>
          <div class="flex justify-between">
            <dt class="text-gray-600">VAT ({{ form.vat_rate || '0' }}%)</dt>
            <dd>{{ vatAmount }} {{ form.currency }}</dd>
          </div>
          <div class="flex justify-between border-t pt-2 font-semibold">
            <dt>Total</dt>
            <dd>{{ total }} {{ form.currency }}</dd>
          </div>
        </dl>
      </div>

      <!-- Notes -->
      <div>
        <label class="block text-sm font-medium text-gray-700">Notes</label>
        <textarea v-model="form.notes" rows="3" maxlength="2000"
          class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>

      <p v-if="error" class="text-sm text-red-600">{{ error }}</p>

      <div class="flex flex-col gap-3 sm:flex-row">
        <button type="submit" :disabled="saving"
          class="w-full sm:w-auto rounded-md bg-blue-600 px-6 py-2 font-medium text-white hover:bg-blue-700 disabled:opacity-50">
          {{ saving ? 'Saving...' : (isEdit ? 'Update Invoice' : 'Create Invoice') }}
        </button>
        <RouterLink to="/invoices"
          class="w-full sm:w-auto rounded-md border border-gray-300 px-6 py-2 text-center font-medium text-gray-700 hover:bg-gray-50">
          Cancel
        </RouterLink>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, watch, onMounted } from 'vue'
import { RouterLink, useRouter, useRoute } from 'vue-router'
import type { CompanyResponse, ClientResponse, BankAccountResponse } from '@/types/api'
import { fetchCompanies } from '@/services/companyApi'
import { fetchClients } from '@/services/clientApi'
import { fetchBankAccounts } from '@/services/bankAccountApi'
import { createInvoice, updateInvoice, fetchInvoice } from '@/services/invoiceApi'
import type { InvoiceItemInput } from '@/services/invoiceApi'

const router = useRouter()
const route = useRoute()

const invoiceId = computed(() => {
  const id = route.params.id
  return id ? Number(id) : null
})
const isEdit = computed(() => !!invoiceId.value)

const loadingSetup = ref(true)
const saving = ref(false)
const error = ref('')

const companies = ref<CompanyResponse[]>([])
const clients = ref<ClientResponse[]>([])
const bankAccounts = ref<BankAccountResponse[]>([])
const activeClients = computed(() => clients.value.filter((c) => c.status === 'active'))

const selectedCompanyId = ref(0)
const selectedClientId = ref(0)
const selectedBankAccountId = ref(0)

const form = reactive({
  invoice_number: '',
  issue_date: new Date().toISOString().slice(0, 10),
  due_date: '',
  currency: 'EUR',
  vat_rate: '0.00',
  contract_reference: '',
  external_reference: '',
  notes: '',
})

const items = ref<InvoiceItemInput[]>([
  { description: '', quantity: '', unit_price: '' },
])

function addItem() {
  items.value.push({ description: '', quantity: '', unit_price: '' })
}

function calcItemTotal(item: InvoiceItemInput): string {
  const qty = parseFloat(item.quantity)
  const price = parseFloat(item.unit_price)
  if (isNaN(qty) || isNaN(price)) return '0.00'
  return (Math.round(qty * price * 100) / 100).toFixed(2)
}

const subtotal = computed(() => {
  return items.value
    .reduce((sum, item) => sum + parseFloat(calcItemTotal(item)), 0)
    .toFixed(2)
})

const vatAmount = computed(() => {
  const rate = parseFloat(form.vat_rate) || 0
  return (Math.round(parseFloat(subtotal.value) * (rate / 100) * 100) / 100).toFixed(2)
})

const total = computed(() => {
  return (parseFloat(subtotal.value) + parseFloat(vatAmount.value)).toFixed(2)
})

async function onCompanyChange() {
  selectedBankAccountId.value = 0
  if (selectedCompanyId.value) {
    bankAccounts.value = await fetchBankAccounts(selectedCompanyId.value)
    const defaultAccount = bankAccounts.value.find((ba) => ba.is_default)
    if (defaultAccount) {
      selectedBankAccountId.value = defaultAccount.id
    }
  } else {
    bankAccounts.value = []
  }
}

watch(() => selectedCompanyId.value, () => {
  if (!isEdit.value) onCompanyChange()
})

async function loadSetup() {
  loadingSetup.value = true
  ;[companies.value, clients.value] = await Promise.all([
    fetchCompanies(),
    fetchClients(),
  ])

  if (invoiceId.value) {
    const inv = await fetchInvoice(invoiceId.value)
    selectedCompanyId.value = inv.company_id
    selectedClientId.value = inv.client_id
    selectedBankAccountId.value = inv.bank_account_id
    form.invoice_number = inv.invoice_number
    form.issue_date = inv.issue_date
    form.due_date = inv.due_date
    form.currency = inv.currency
    form.vat_rate = inv.vat_rate
    form.contract_reference = inv.contract_reference ?? ''
    form.external_reference = inv.external_reference ?? ''
    form.notes = inv.notes ?? ''
    items.value = (inv.items ?? []).map((it) => ({
      description: it.description,
      quantity: it.quantity,
      unit_price: it.unit_price,
    }))
    if (items.value.length === 0) {
      items.value = [{ description: '', quantity: '', unit_price: '' }]
    }
    bankAccounts.value = await fetchBankAccounts(inv.company_id)
  }

  loadingSetup.value = false
}

async function handleSubmit() {
  if (!selectedCompanyId.value || !selectedClientId.value || !selectedBankAccountId.value) {
    error.value = 'Please select a company, client, and bank account'
    return
  }

  saving.value = true
  error.value = ''
  try {
    const data = {
      company_id: selectedCompanyId.value,
      client_id: selectedClientId.value,
      bank_account_id: selectedBankAccountId.value,
      issue_date: form.issue_date,
      due_date: form.due_date,
      currency: form.currency,
      vat_rate: form.vat_rate || '0.00',
      contract_reference: form.contract_reference || null,
      external_reference: form.external_reference || null,
      notes: form.notes || null,
      items: items.value.filter((it) => it.description.trim()),
    }

    if (invoiceId.value) {
      await updateInvoice(invoiceId.value, {
        ...data,
        invoice_number: form.invoice_number,
      })
    } else {
      await createInvoice(data)
    }
    router.push('/invoices')
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Failed to save invoice'
  } finally {
    saving.value = false
  }
}

onMounted(loadSetup)
</script>
