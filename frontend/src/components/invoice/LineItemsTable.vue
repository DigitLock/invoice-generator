<template>
  <fieldset class="space-y-4">
    <legend class="text-lg font-medium text-gray-900">Line Items</legend>
    <p v-if="errors.lineItems" class="text-sm text-red-600">{{ errors.lineItems }}</p>

    <!-- Mobile: stacked cards -->
    <div class="space-y-4 sm:hidden">
      <div v-for="(item, index) in invoice.items" :key="index"
        class="rounded-lg border border-gray-200 bg-white p-3 space-y-3">
        <div class="flex items-center justify-between">
          <span class="text-sm font-medium text-gray-500">Item {{ index + 1 }}</span>
          <button v-if="invoice.items.length > 1" type="button" @click="removeItem(index)"
            class="text-red-500 hover:text-red-700 text-lg leading-none">&times;</button>
        </div>
        <div>
          <label class="block text-xs text-gray-500">Description</label>
          <input v-model="item.description" type="text" maxlength="500"
            class="mt-1 w-full rounded-md border border-gray-300 px-2 py-1.5 text-sm shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
        </div>
        <div class="grid grid-cols-3 gap-2">
          <div>
            <label class="block text-xs text-gray-500">Qty</label>
            <input v-model="item.quantity" type="text"
              class="mt-1 w-full rounded-md border border-gray-300 px-2 py-1.5 text-sm shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div>
            <label class="block text-xs text-gray-500">Unit Price</label>
            <input v-model="item.unitPrice" type="text"
              class="mt-1 w-full rounded-md border border-gray-300 px-2 py-1.5 text-sm shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </div>
          <div>
            <label class="block text-xs text-gray-500">Total</label>
            <div class="mt-1 px-2 py-1.5 text-sm text-right text-gray-700">{{ item.total }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Desktop: table -->
    <div class="hidden sm:block">
    <table class="w-full text-sm">
      <thead>
        <tr class="border-b text-left text-gray-700">
          <th class="pb-2">#</th>
          <th class="pb-2">Description</th>
          <th class="pb-2 w-28">Quantity</th>
          <th class="pb-2 w-32">Unit Price</th>
          <th class="pb-2 w-32">Total</th>
          <th class="pb-2 w-10"></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in invoice.items" :key="index" class="border-b">
          <td class="py-2 pr-2 text-gray-500">{{ index + 1 }}</td>
          <td class="py-2 pr-2">
            <input v-model="item.description" type="text" maxlength="500"
              class="w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </td>
          <td class="py-2 pr-2">
            <input v-model="item.quantity" type="text"
              class="w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </td>
          <td class="py-2 pr-2">
            <input v-model="item.unitPrice" type="text"
              class="w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
          </td>
          <td class="py-2 pr-2 text-right">{{ item.total }}</td>
          <td class="py-2">
            <button v-if="invoice.items.length > 1" type="button" @click="removeItem(index)"
              class="text-red-500 hover:text-red-700">&times;</button>
          </td>
        </tr>
      </tbody>
    </table>
    </div>
    <button v-if="invoice.items.length < 10" type="button" @click="addItem"
      class="text-sm text-blue-600 hover:text-blue-800">
      + Add Item
    </button>
  </fieldset>
</template>

<script setup lang="ts">
import { inject } from 'vue'
import { storeToRefs } from 'pinia'
import { useInvoiceStore, createEmptyLineItem } from '@/stores/invoice'
import type { useValidation } from '@/composables/useValidation'

const { invoice } = storeToRefs(useInvoiceStore())
const { errors } = inject('validation') as ReturnType<typeof useValidation>

function addItem() {
  invoice.value.items.push(createEmptyLineItem())
}

function removeItem(index: number) {
  invoice.value.items.splice(index, 1)
}
</script>
