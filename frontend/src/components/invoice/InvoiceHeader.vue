<template>
  <fieldset class="space-y-4">
    <legend class="text-lg font-medium text-gray-900">Invoice Details</legend>
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
      <div>
        <label for="invoiceNumber" class="block text-sm font-medium text-gray-700">Invoice Number</label>
        <input id="invoiceNumber" v-model="invoice.invoiceNumber" type="text" maxlength="50"
          class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label for="issueDate" class="block text-sm font-medium text-gray-700">Issue Date</label>
        <input id="issueDate" v-model="invoice.issueDate" type="date"
          class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label for="dueDate" class="block text-sm font-medium text-gray-700">Due Date</label>
        <input id="dueDate" v-model="invoice.dueDate" type="date"
          class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <CurrencySelector />
      <div>
        <label for="vatRate" class="block text-sm font-medium text-gray-700">VAT Rate (%)</label>
        <input id="vatRate" v-model="invoice.vatRate" type="text"
          class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label for="contractReference" class="block text-sm font-medium text-gray-700">Contract Reference</label>
        <input id="contractReference" v-model="invoice.contractReference" type="text" maxlength="255"
          class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label for="externalReference" class="block text-sm font-medium text-gray-700">External Reference</label>
        <input id="externalReference" v-model="invoice.externalReference" type="text" maxlength="255"
          class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
    </div>
  </fieldset>
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useInvoiceStore } from '@/stores/invoice'
import CurrencySelector from './CurrencySelector.vue'

const { invoice } = storeToRefs(useInvoiceStore())
</script>
