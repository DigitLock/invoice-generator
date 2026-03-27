<template>
  <div>
    <label for="currency" class="block text-sm font-medium text-gray-700">Currency</label>
    <div class="mt-1 flex gap-2">
      <select id="currency" v-model="selectedOption"
        class="block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500">
        <option value="EUR">EUR</option>
        <option value="RSD">RSD</option>
        <option value="custom">Other...</option>
      </select>
      <input v-if="selectedOption === 'custom'" v-model="customCode" type="text"
        maxlength="3" placeholder="USD"
        class="block w-20 rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500 uppercase" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { storeToRefs } from 'pinia'
import { useInvoiceStore } from '@/stores/invoice'

const { invoice } = storeToRefs(useInvoiceStore())

const predefined = ['EUR', 'RSD']
const selectedOption = ref(predefined.includes(invoice.value.currency) ? invoice.value.currency : 'custom')
const customCode = ref(predefined.includes(invoice.value.currency) ? '' : invoice.value.currency)

watch(selectedOption, (val) => {
  if (val !== 'custom') {
    invoice.value.currency = val
    customCode.value = ''
  }
})

watch(customCode, (val) => {
  if (selectedOption.value === 'custom' && val) {
    invoice.value.currency = val.toUpperCase()
  }
})
</script>
