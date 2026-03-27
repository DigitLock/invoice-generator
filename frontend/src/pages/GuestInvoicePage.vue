<template>
  <div class="mx-auto max-w-4xl px-4 py-8">
    <h1 class="mb-8 text-2xl font-bold text-gray-900">Create Invoice</h1>
    <InvoiceForm ref="formRef" @submit="handleSubmit">
      <template #actions>
        <button type="submit"
          class="w-full sm:w-auto rounded-md bg-blue-600 px-6 py-2 text-white font-medium hover:bg-blue-700">
          Generate PDF
        </button>
        <RouterLink to="/"
          class="w-full sm:w-auto rounded-md border border-gray-300 px-6 py-2 text-center text-gray-700 font-medium hover:bg-gray-50">
          Cancel
        </RouterLink>
      </template>
    </InvoiceForm>

    <ConfirmDialog
      :visible="showConfirm"
      @confirm="handleConfirm"
      @cancel="showConfirm = false"
    />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { RouterLink } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useInvoiceStore } from '@/stores/invoice'
import InvoiceForm from '@/components/invoice/InvoiceForm.vue'
import ConfirmDialog from '@/components/invoice/ConfirmDialog.vue'

const { invoice } = storeToRefs(useInvoiceStore())
const formRef = ref<InstanceType<typeof InvoiceForm> | null>(null)
const showConfirm = ref(false)

function handleSubmit() {
  const validation = formRef.value?.validation
  if (!validation) return

  if (!validation.validate(invoice.value)) return

  showConfirm.value = true
}

async function handleConfirm() {
  showConfirm.value = false
  const { generatePdf } = await import('@/services/pdfGenerator')
  await generatePdf(invoice.value)
}
</script>
