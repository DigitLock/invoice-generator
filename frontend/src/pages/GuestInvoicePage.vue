<template>
  <div class="mx-auto max-w-4xl px-4 py-8">
    <h1 class="mb-8 text-2xl font-bold text-gray-900">Create Invoice</h1>
    <InvoiceForm @submit="handleGeneratePdf">
      <template #actions>
        <button type="submit"
          class="rounded-md bg-blue-600 px-6 py-2 text-white font-medium hover:bg-blue-700">
          Generate PDF
        </button>
        <RouterLink to="/"
          class="rounded-md border border-gray-300 px-6 py-2 text-gray-700 font-medium hover:bg-gray-50">
          Cancel
        </RouterLink>
      </template>
    </InvoiceForm>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useInvoiceStore } from '@/stores/invoice'
import InvoiceForm from '@/components/invoice/InvoiceForm.vue'
import { generatePdf } from '@/services/pdfGenerator'

const { invoice } = storeToRefs(useInvoiceStore())

async function handleGeneratePdf() {
  await generatePdf(invoice.value)
}
</script>
