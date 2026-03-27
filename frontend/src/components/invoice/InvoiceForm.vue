<template>
  <form @submit.prevent="$emit('submit')" class="space-y-8">
    <InvoiceHeader />
    <SellerSection />
    <BuyerSection />
    <LineItemsTable />
    <TotalsSection />
    <PaymentDetails />

    <div>
      <label for="notes" class="block text-sm font-medium text-gray-700">Notes</label>
      <textarea
        id="notes"
        v-model="invoice.notes"
        rows="3"
        maxlength="2000"
        class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500"
      />
    </div>

    <div class="flex gap-3">
      <slot name="actions" />
    </div>
  </form>
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useInvoiceStore } from '@/stores/invoice'
import InvoiceHeader from './InvoiceHeader.vue'
import SellerSection from './SellerSection.vue'
import BuyerSection from './BuyerSection.vue'
import LineItemsTable from './LineItemsTable.vue'
import TotalsSection from './TotalsSection.vue'
import PaymentDetails from './PaymentDetails.vue'

defineEmits<{
  submit: []
}>()

const { invoice } = storeToRefs(useInvoiceStore())
</script>
