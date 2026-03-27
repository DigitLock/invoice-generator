<template>
  <div class="mx-auto max-w-7xl px-4 py-8">
    <div class="flex items-center justify-between mb-8">
      <h1 class="text-2xl font-bold text-gray-900">Dashboard</h1>
      <RouterLink to="/invoices/new"
        class="rounded-md bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700">
        New Invoice
      </RouterLink>
    </div>

    <div class="rounded-lg border border-gray-200 bg-white p-6">
      <h2 class="text-lg font-medium text-gray-900 mb-4">Recent Invoices</h2>

      <div v-if="loading" class="text-sm text-gray-500">Loading...</div>
      <div v-else-if="invoices.length === 0" class="text-sm text-gray-500">
        No invoices yet. Create your first invoice to get started.
      </div>
      <table v-else class="w-full text-sm">
        <thead>
          <tr class="border-b text-left text-gray-600">
            <th class="pb-2 font-medium">Invoice</th>
            <th class="pb-2 font-medium">Client</th>
            <th class="pb-2 font-medium">Date</th>
            <th class="pb-2 font-medium text-right">Total</th>
            <th class="pb-2 font-medium">Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="inv in invoices" :key="inv.id" class="border-b last:border-0">
            <td class="py-3">
              <RouterLink :to="`/invoices/${inv.id}`" class="text-blue-600 hover:text-blue-800">
                {{ inv.invoice_number }}
              </RouterLink>
            </td>
            <td class="py-3 text-gray-700">{{ inv.client_name }}</td>
            <td class="py-3 text-gray-500">{{ inv.issue_date }}</td>
            <td class="py-3 text-right text-gray-900">{{ inv.total }} {{ inv.currency }}</td>
            <td class="py-3">
              <span :class="statusClass(inv.status)"
                class="inline-block rounded-full px-2 py-0.5 text-xs font-medium">
                {{ inv.status }}
              </span>
              <span v-if="inv.is_overdue"
                class="ml-1 inline-block rounded-full bg-red-100 text-red-700 px-2 py-0.5 text-xs font-medium">
                Overdue
              </span>
            </td>
          </tr>
        </tbody>
      </table>

      <div v-if="invoices.length > 0" class="mt-4">
        <RouterLink to="/invoices" class="text-sm text-blue-600 hover:text-blue-800">
          View all invoices
        </RouterLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import { fetchInvoices } from '@/services/invoiceApi'
import type { InvoiceListItem } from '@/types/api'

const invoices = ref<InvoiceListItem[]>([])
const loading = ref(true)

function statusClass(status: string): string {
  const classes: Record<string, string> = {
    draft: 'bg-gray-100 text-gray-700',
    sent: 'bg-blue-100 text-blue-700',
    partially_paid: 'bg-orange-100 text-orange-700',
    paid: 'bg-green-100 text-green-700',
    cancelled: 'bg-red-100 text-red-700',
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

onMounted(async () => {
  try {
    const response = await fetchInvoices({ page: 1, page_size: 10 })
    invoices.value = response.invoices
  } catch {
    // API error handled by api.ts (401 → redirect)
  } finally {
    loading.value = false
  }
})
</script>
