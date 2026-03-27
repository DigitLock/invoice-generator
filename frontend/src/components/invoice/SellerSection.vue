<template>
  <fieldset class="space-y-4">
    <legend class="text-lg font-medium text-gray-900">Seller (Your Company)</legend>
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
      <div>
        <label for="sellerName" class="block text-sm font-medium text-gray-700">Company Name</label>
        <input id="sellerName" v-model="invoice.seller.name" type="text"
          :class="['mt-1 block w-full rounded-md shadow-sm focus:ring-1 focus:ring-blue-500',
            errors.sellerName ? 'border border-red-500 focus:border-red-500' : 'border border-gray-300 focus:border-blue-500']" />
        <p v-if="errors.sellerName" class="mt-1 text-sm text-red-600">{{ errors.sellerName }}</p>
      </div>
      <div>
        <label for="sellerContact" class="block text-sm font-medium text-gray-700">Contact Person</label>
        <input id="sellerContact" v-model="invoice.seller.contactPerson" type="text"
          class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div class="sm:col-span-2">
        <label for="sellerAddress" class="block text-sm font-medium text-gray-700">Address</label>
        <input id="sellerAddress" v-model="invoice.seller.address" type="text"
          :class="['mt-1 block w-full rounded-md shadow-sm focus:ring-1 focus:ring-blue-500',
            errors.sellerAddress ? 'border border-red-500 focus:border-red-500' : 'border border-gray-300 focus:border-blue-500']" />
        <p v-if="errors.sellerAddress" class="mt-1 text-sm text-red-600">{{ errors.sellerAddress }}</p>
      </div>
      <div>
        <label for="sellerPhone" class="block text-sm font-medium text-gray-700">Phone</label>
        <input id="sellerPhone" v-model="invoice.seller.phone" type="text"
          class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label for="sellerVat" class="block text-sm font-medium text-gray-700">VAT Number</label>
        <input id="sellerVat" v-model="invoice.seller.vatNumber" type="text"
          class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label for="sellerReg" class="block text-sm font-medium text-gray-700">Registration Number</label>
        <input id="sellerReg" v-model="invoice.seller.regNumber" type="text"
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
