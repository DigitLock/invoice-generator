<template>
  <fieldset class="space-y-4">
    <legend class="text-lg font-medium text-gray-900">Buyer (Client)</legend>
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
      <div>
        <label for="buyerName" class="block text-sm font-medium text-gray-700">Company Name</label>
        <input id="buyerName" v-model="invoice.buyer.name" type="text"
          :class="['mt-1 block w-full rounded-md shadow-sm focus:ring-1 focus:ring-blue-500',
            errors.buyerName ? 'border border-red-500 focus:border-red-500' : 'border border-gray-300 focus:border-blue-500']" />
        <p v-if="errors.buyerName" class="mt-1 text-sm text-red-600">{{ errors.buyerName }}</p>
      </div>
      <div>
        <label for="buyerContact" class="block text-sm font-medium text-gray-700">Contact Person</label>
        <input id="buyerContact" v-model="invoice.buyer.contactPerson" type="text"
          class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label for="buyerEmail" class="block text-sm font-medium text-gray-700">Email</label>
        <input id="buyerEmail" v-model="invoice.buyer.email" type="email"
          :class="['mt-1 block w-full rounded-md shadow-sm focus:ring-1 focus:ring-blue-500',
            errors.buyerEmail ? 'border border-red-500 focus:border-red-500' : 'border border-gray-300 focus:border-blue-500']" />
        <p v-if="errors.buyerEmail" class="mt-1 text-sm text-red-600">{{ errors.buyerEmail }}</p>
      </div>
      <div class="sm:col-span-2">
        <label for="buyerAddress" class="block text-sm font-medium text-gray-700">Address</label>
        <input id="buyerAddress" v-model="invoice.buyer.address" type="text"
          :class="['mt-1 block w-full rounded-md shadow-sm focus:ring-1 focus:ring-blue-500',
            errors.buyerAddress ? 'border border-red-500 focus:border-red-500' : 'border border-gray-300 focus:border-blue-500']" />
        <p v-if="errors.buyerAddress" class="mt-1 text-sm text-red-600">{{ errors.buyerAddress }}</p>
      </div>
      <div>
        <label for="buyerVat" class="block text-sm font-medium text-gray-700">VAT Number</label>
        <input id="buyerVat" v-model="invoice.buyer.vatNumber" type="text"
          class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label for="buyerReg" class="block text-sm font-medium text-gray-700">Registration Number</label>
        <input id="buyerReg" v-model="invoice.buyer.regNumber" type="text"
          class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
    </div>
  </fieldset>
</template>

<script setup lang="ts">
import { inject } from 'vue'
import { storeToRefs } from 'pinia'
import { useInvoiceStore } from '@/stores/invoice'
import type { useValidation } from '@/composables/useValidation'

const { invoice } = storeToRefs(useInvoiceStore())
const { errors } = inject('validation') as ReturnType<typeof useValidation>
</script>
