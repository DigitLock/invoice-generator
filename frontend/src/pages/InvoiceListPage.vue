<template>
  <div class="mx-auto max-w-7xl px-4 py-8">
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-2xl font-bold text-gray-900">Invoices</h1>
      <RouterLink to="/invoices/new"
        class="rounded-md bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700">
        New Invoice
      </RouterLink>
    </div>

    <!-- Filters -->
    <div class="mb-6 flex flex-wrap items-end gap-4">
      <div>
        <label class="block text-xs font-medium text-gray-500">Status</label>
        <select v-model="filters.status" @change="loadPage(1)"
          class="mt-1 rounded-md border border-gray-300 px-3 py-1.5 text-sm shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500">
          <option value="">All</option>
          <option value="draft">Draft</option>
          <option value="sent">Sent</option>
          <option value="partially_paid">Partially Paid</option>
          <option value="paid">Paid</option>
          <option value="cancelled">Cancelled</option>
        </select>
      </div>
      <div>
        <label class="block text-xs font-medium text-gray-500">From</label>
        <input v-model="filters.date_from" type="date" @change="loadPage(1)"
          class="mt-1 rounded-md border border-gray-300 px-3 py-1.5 text-sm shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label class="block text-xs font-medium text-gray-500">To</label>
        <input v-model="filters.date_to" type="date" @change="loadPage(1)"
          class="mt-1 rounded-md border border-gray-300 px-3 py-1.5 text-sm shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label class="block text-xs font-medium text-gray-500">Search</label>
        <input v-model="filters.search" type="text" placeholder="Invoice # or client..." @keyup.enter="loadPage(1)"
          class="mt-1 rounded-md border border-gray-300 px-3 py-1.5 text-sm shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
    </div>

    <div v-if="loading" class="text-sm text-gray-500">Loading...</div>
    <div v-else-if="invoices.length === 0" class="rounded-lg border border-gray-200 bg-white p-8 text-center text-gray-500">
      No invoices found.
    </div>
    <div v-else>
      <div class="-mx-4 overflow-x-auto px-4 sm:mx-0 sm:px-0">
      <table class="w-full min-w-[700px] text-sm">
        <thead>
          <tr class="border-b text-left text-gray-600">
            <th class="pb-2 font-medium">Invoice #</th>
            <th class="pb-2 font-medium">Company</th>
            <th class="pb-2 font-medium">Client</th>
            <th class="pb-2 font-medium">Date</th>
            <th class="pb-2 font-medium text-right">Total</th>
            <th class="pb-2 font-medium">Status</th>
            <th class="pb-2 font-medium text-right">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="inv in invoices" :key="inv.id" class="border-b hover:bg-gray-50">
            <td class="py-3">
              <RouterLink :to="`/invoices/${inv.id}`" class="text-blue-600 hover:text-blue-800">
                {{ inv.invoice_number }}
              </RouterLink>
            </td>
            <td class="py-3 text-gray-700">{{ inv.company_name }}</td>
            <td class="py-3 text-gray-700">{{ inv.client_name }}</td>
            <td class="py-3 text-gray-500">{{ inv.issue_date }}</td>
            <td class="py-3 text-right text-gray-900">{{ inv.total }} {{ inv.currency }}</td>
            <td class="py-3">
              <StatusBadge :status="inv.status" />
              <OverdueBadge :is-overdue="inv.is_overdue" class="ml-1" />
            </td>
            <td class="py-3 text-right">
              <div class="flex justify-end gap-2">
                <RouterLink :to="`/invoices/${inv.id}`" class="text-gray-500 hover:text-gray-700" title="View">
                  View
                </RouterLink>
                <RouterLink :to="`/invoices/${inv.id}/edit`" class="text-blue-600 hover:text-blue-800" title="Edit">
                  Edit
                </RouterLink>
                <button @click="handleDownloadPdf(inv)" class="text-gray-500 hover:text-gray-700" title="PDF">
                  PDF
                </button>
                <button @click="handleDelete(inv)" class="text-red-600 hover:text-red-800" title="Delete">
                  Delete
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      </div>

      <!-- Pagination -->
      <div v-if="pagination" class="mt-4 flex items-center justify-between text-sm text-gray-600">
        <span>Page {{ pagination.page }} of {{ pagination.total_pages }} ({{ pagination.total_items }} total)</span>
        <div class="flex gap-2">
          <button @click="loadPage(pagination!.page - 1)" :disabled="!pagination.has_previous"
            class="rounded-md border border-gray-300 px-3 py-1 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed">
            Previous
          </button>
          <button @click="loadPage(pagination!.page + 1)" :disabled="!pagination.has_next"
            class="rounded-md border border-gray-300 px-3 py-1 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed">
            Next
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import type { InvoiceListItem, PaginationMeta } from '@/types/api'
import { fetchInvoices, deleteInvoice, downloadInvoicePdf } from '@/services/invoiceApi'
import StatusBadge from '@/components/invoice/StatusBadge.vue'
import OverdueBadge from '@/components/invoice/OverdueBadge.vue'

const invoices = ref<InvoiceListItem[]>([])
const pagination = ref<PaginationMeta | null>(null)
const loading = ref(true)

const filters = reactive({
  status: '',
  date_from: '',
  date_to: '',
  search: '',
})

async function loadPage(page: number) {
  loading.value = true
  try {
    const response = await fetchInvoices({
      page,
      page_size: 20,
      status: filters.status || undefined,
      date_from: filters.date_from || undefined,
      date_to: filters.date_to || undefined,
      search: filters.search || undefined,
    })
    invoices.value = response.invoices
    pagination.value = response.pagination
  } catch {
    // handled by api.ts
  } finally {
    loading.value = false
  }
}

async function handleDelete(inv: InvoiceListItem) {
  if (!confirm(`Delete invoice ${inv.invoice_number}?`)) return
  try {
    await deleteInvoice(inv.id)
    await loadPage(pagination.value?.page ?? 1)
  } catch (e) {
    alert(e instanceof Error ? e.message : 'Failed to delete')
  }
}

async function handleDownloadPdf(inv: InvoiceListItem) {
  await downloadInvoicePdf(inv.id, inv.invoice_number)
}

onMounted(() => loadPage(1))
</script>
