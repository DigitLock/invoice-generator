<template>
  <fieldset class="space-y-4">
    <legend class="text-lg font-medium text-gray-900">Line Items</legend>
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
    <button v-if="invoice.items.length < 10" type="button" @click="addItem"
      class="text-sm text-blue-600 hover:text-blue-800">
      + Add Item
    </button>
  </fieldset>
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useInvoiceStore, createEmptyLineItem } from '@/stores/invoice'

const { invoice } = storeToRefs(useInvoiceStore())

function addItem() {
  invoice.value.items.push(createEmptyLineItem())
}

function removeItem(index: number) {
  invoice.value.items.splice(index, 1)
}
</script>
