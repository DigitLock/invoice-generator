<template>
  <div class="mx-auto max-w-4xl px-4 py-8">
    <div v-if="loading" class="text-sm text-gray-500">Loading...</div>
    <template v-else-if="invoice">
      <!-- Header -->
      <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between mb-8">
        <div>
          <h1 class="text-2xl font-bold text-gray-900">{{ invoice.invoice_number }}</h1>
          <div class="mt-1 flex items-center gap-2">
            <StatusBadge :status="invoice.status" />
            <OverdueBadge :is-overdue="invoice.is_overdue" />
          </div>
        </div>
        <div class="flex flex-wrap gap-2">
          <button @click="handleDownloadPdf"
            class="rounded-md bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700">
            Download PDF
          </button>
          <RouterLink :to="`/invoices/${invoice.id}/edit`"
            class="rounded-md border border-gray-300 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">
            Edit
          </RouterLink>
          <button @click="handleDelete"
            class="rounded-md border border-red-300 px-4 py-2 text-sm font-medium text-red-600 hover:bg-red-50">
            Delete
          </button>
        </div>
      </div>

      <!-- Status management -->
      <div v-if="invoice.status !== 'paid' && invoice.status !== 'cancelled'"
        class="mb-8 flex flex-col gap-4 rounded-lg border border-gray-200 bg-white p-4 sm:flex-row sm:items-end">
        <div class="flex-1">
          <StatusSelect :current-status="invoice.status" @change="handleStatusChange" />
        </div>
        <div v-if="invoice.status !== 'draft'" class="flex items-center gap-2">
          <input type="checkbox" id="overdue" :checked="invoice.is_overdue" @change="handleOverdueToggle"
            class="rounded border-gray-300 text-red-600 focus:ring-red-500" />
          <label for="overdue" class="text-sm text-gray-700">Mark as overdue</label>
        </div>
      </div>

      <!-- Invoice details -->
      <div class="space-y-6">
        <!-- Dates & references -->
        <div class="rounded-lg border border-gray-200 bg-white p-4">
          <div class="grid grid-cols-2 gap-4 text-sm sm:grid-cols-4">
            <div>
              <span class="text-gray-500">Issue Date</span>
              <div class="font-medium">{{ invoice.issue_date }}</div>
            </div>
            <div>
              <span class="text-gray-500">Due Date</span>
              <div class="font-medium">{{ invoice.due_date }}</div>
            </div>
            <div v-if="invoice.contract_reference">
              <span class="text-gray-500">Contract Ref</span>
              <div class="font-medium">{{ invoice.contract_reference }}</div>
            </div>
            <div v-if="invoice.external_reference">
              <span class="text-gray-500">External Ref</span>
              <div class="font-medium">{{ invoice.external_reference }}</div>
            </div>
          </div>
        </div>

        <!-- Seller / Buyer -->
        <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div class="rounded-lg border border-gray-200 bg-white p-4">
            <h3 class="mb-2 text-sm font-medium text-gray-500">From</h3>
            <div class="text-sm space-y-1" v-if="invoice.company">
              <div class="font-medium text-gray-900">{{ invoice.company.name }}</div>
              <div>{{ invoice.company.contact_person }}</div>
              <div class="text-gray-600">{{ invoice.company.address }}</div>
              <div v-if="invoice.company.phone" class="text-gray-500">{{ invoice.company.phone }}</div>
              <div v-if="invoice.company.vat_number" class="text-gray-500">VAT: {{ invoice.company.vat_number }}</div>
              <div v-if="invoice.company.reg_number" class="text-gray-500">Reg: {{ invoice.company.reg_number }}</div>
            </div>
          </div>
          <div class="rounded-lg border border-gray-200 bg-white p-4">
            <h3 class="mb-2 text-sm font-medium text-gray-500">Bill To</h3>
            <div class="text-sm space-y-1" v-if="invoice.client">
              <div class="font-medium text-gray-900">{{ invoice.client.name }}</div>
              <div v-if="invoice.client.contact_person">{{ invoice.client.contact_person }}</div>
              <div v-if="invoice.client.email">{{ invoice.client.email }}</div>
              <div class="text-gray-600">{{ invoice.client.address }}</div>
              <div v-if="invoice.client.vat_number" class="text-gray-500">VAT: {{ invoice.client.vat_number }}</div>
              <div v-if="invoice.client.reg_number" class="text-gray-500">Reg: {{ invoice.client.reg_number }}</div>
            </div>
          </div>
        </div>

        <!-- Line items -->
        <div class="rounded-lg border border-gray-200 bg-white p-4">
          <h3 class="mb-3 text-sm font-medium text-gray-500">Line Items</h3>
          <table class="w-full text-sm">
            <thead>
              <tr class="border-b text-left text-gray-600">
                <th class="pb-2 font-medium">Description</th>
                <th class="pb-2 font-medium text-center w-20">Qty</th>
                <th class="pb-2 font-medium text-right w-28">Unit Price</th>
                <th class="pb-2 font-medium text-right w-28">Total</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in invoice.items" :key="item.id" class="border-b last:border-0">
                <td class="py-2">{{ item.description }}</td>
                <td class="py-2 text-center">{{ item.quantity }}</td>
                <td class="py-2 text-right">{{ item.unit_price }} {{ invoice.currency }}</td>
                <td class="py-2 text-right">{{ item.total }} {{ invoice.currency }}</td>
              </tr>
            </tbody>
          </table>

          <div class="mt-4 flex justify-end">
            <dl class="w-64 space-y-2 text-sm">
              <div class="flex justify-between">
                <dt class="text-gray-600">Subtotal</dt>
                <dd>{{ invoice.subtotal }} {{ invoice.currency }}</dd>
              </div>
              <div class="flex justify-between">
                <dt class="text-gray-600">VAT ({{ invoice.vat_rate }}%)</dt>
                <dd>{{ invoice.vat_amount }} {{ invoice.currency }}</dd>
              </div>
              <div class="flex justify-between border-t pt-2 font-semibold">
                <dt>Total</dt>
                <dd>{{ invoice.total }} {{ invoice.currency }}</dd>
              </div>
            </dl>
          </div>
        </div>

        <!-- Bank account -->
        <div v-if="invoice.bank_account" class="rounded-lg border border-gray-200 bg-white p-4">
          <h3 class="mb-2 text-sm font-medium text-gray-500">Payment Details</h3>
          <div class="text-sm space-y-1">
            <div><span class="text-gray-500">Account Holder:</span> {{ invoice.bank_account.account_holder }}</div>
            <div><span class="text-gray-500">Bank:</span> {{ invoice.bank_account.bank_name }}</div>
            <div v-if="invoice.bank_account.bank_address"><span class="text-gray-500">Bank Address:</span> {{ invoice.bank_account.bank_address }}</div>
            <div><span class="text-gray-500">IBAN:</span> {{ invoice.bank_account.iban }}</div>
            <div><span class="text-gray-500">SWIFT:</span> {{ invoice.bank_account.swift }}</div>
          </div>
        </div>

        <!-- Notes -->
        <div v-if="invoice.notes" class="rounded-lg border border-gray-200 bg-white p-4">
          <h3 class="mb-2 text-sm font-medium text-gray-500">Notes</h3>
          <p class="text-sm text-gray-700 whitespace-pre-wrap">{{ invoice.notes }}</p>
        </div>
      </div>

      <div class="mt-6">
        <RouterLink to="/invoices" class="text-sm text-blue-600 hover:text-blue-800">Back to invoices</RouterLink>
      </div>
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { RouterLink, useRoute, useRouter } from 'vue-router'
import type { InvoiceResponse } from '@/types/api'
import { fetchInvoice, deleteInvoice, updateInvoiceStatus, updateInvoiceOverdue, downloadInvoicePdf } from '@/services/invoiceApi'
import StatusBadge from '@/components/invoice/StatusBadge.vue'
import OverdueBadge from '@/components/invoice/OverdueBadge.vue'
import StatusSelect from '@/components/invoice/StatusSelect.vue'

const route = useRoute()
const router = useRouter()
const invoice = ref<InvoiceResponse | null>(null)
const loading = ref(true)

async function load() {
  loading.value = true
  const id = Number(route.params.id)
  invoice.value = await fetchInvoice(id)
  loading.value = false
}

async function handleStatusChange(newStatus: string) {
  if (!invoice.value) return
  if (!confirm(`Change status to "${newStatus}"?`)) return
  invoice.value = await updateInvoiceStatus(invoice.value.id, newStatus)
}

async function handleOverdueToggle() {
  if (!invoice.value) return
  invoice.value = await updateInvoiceOverdue(invoice.value.id, !invoice.value.is_overdue)
}

async function handleDownloadPdf() {
  if (!invoice.value) return
  await downloadInvoicePdf(invoice.value.id, invoice.value.invoice_number)
}

async function handleDelete() {
  if (!invoice.value) return
  if (!confirm(`Delete invoice ${invoice.value.invoice_number}?`)) return
  await deleteInvoice(invoice.value.id)
  router.push('/invoices')
}

onMounted(load)
</script>
