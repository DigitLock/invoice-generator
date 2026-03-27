<template>
  <fieldset class="space-y-4">
    <legend class="text-lg font-medium text-gray-900">Payment Details</legend>
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
      <div>
        <label for="bankName" class="block text-sm font-medium text-gray-700">Bank Name</label>
        <input id="bankName" v-model="invoice.bankAccount.bankName" type="text"
          :class="['mt-1 block w-full rounded-md shadow-sm focus:ring-1 focus:ring-blue-500',
            errors.bankName ? 'border border-red-500 focus:border-red-500' : 'border border-gray-300 focus:border-blue-500']" />
        <p v-if="errors.bankName" class="mt-1 text-sm text-red-600">{{ errors.bankName }}</p>
      </div>
      <div>
        <label for="bankAddress" class="block text-sm font-medium text-gray-700">Bank Address</label>
        <input id="bankAddress" v-model="invoice.bankAccount.bankAddress" type="text"
          class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label for="accountHolder" class="block text-sm font-medium text-gray-700">Account Holder</label>
        <input id="accountHolder" v-model="invoice.bankAccount.accountHolder" type="text"
          class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
      </div>
      <div>
        <label for="iban" class="block text-sm font-medium text-gray-700">IBAN</label>
        <input id="iban" v-model="invoice.bankAccount.iban" type="text"
          :class="['mt-1 block w-full rounded-md shadow-sm focus:ring-1 focus:ring-blue-500',
            errors.iban ? 'border border-red-500 focus:border-red-500' : 'border border-gray-300 focus:border-blue-500']" />
        <p v-if="errors.iban" class="mt-1 text-sm text-red-600">{{ errors.iban }}</p>
      </div>
      <div>
        <label for="swift" class="block text-sm font-medium text-gray-700">SWIFT / BIC</label>
        <input id="swift" v-model="invoice.bankAccount.swift" type="text"
          :class="['mt-1 block w-full rounded-md shadow-sm focus:ring-1 focus:ring-blue-500',
            errors.swift ? 'border border-red-500 focus:border-red-500' : 'border border-gray-300 focus:border-blue-500']" />
        <p v-if="errors.swift" class="mt-1 text-sm text-red-600">{{ errors.swift }}</p>
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
