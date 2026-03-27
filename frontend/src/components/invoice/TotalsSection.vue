<template>
  <div class="flex justify-end">
    <dl class="w-64 space-y-2 text-sm">
      <div class="flex justify-between">
        <dt class="text-gray-600">Subtotal</dt>
        <dd>{{ invoice.subtotal }}</dd>
      </div>
      <div class="flex justify-between">
        <dt class="text-gray-600">VAT ({{ invoice.vatRate }}%)</dt>
        <dd>{{ invoice.vatAmount }}</dd>
      </div>
      <div class="flex justify-between border-t pt-2 font-semibold">
        <dt>Total</dt>
        <dd>{{ invoice.total }} {{ invoice.currency }}</dd>
      </div>
    </dl>
  </div>
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useInvoiceStore } from '@/stores/invoice'

const { invoice } = storeToRefs(useInvoiceStore())
</script>
