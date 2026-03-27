<template>
  <div>
    <label class="block text-sm font-medium text-gray-700">Change Status</label>
    <select v-model="selected" @change="handleChange"
      class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-1 focus:ring-blue-500">
      <option :value="currentStatus" disabled>{{ currentLabel }} (current)</option>
      <option v-for="next in allowedTransitions" :key="next" :value="next">
        {{ statusLabels[next] }}
      </option>
    </select>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

const props = defineProps<{ currentStatus: string }>()
const emit = defineEmits<{ change: [status: string] }>()

const selected = ref(props.currentStatus)

const statusLabels: Record<string, string> = {
  draft: 'Draft',
  sent: 'Sent',
  partially_paid: 'Partially Paid',
  paid: 'Paid',
  cancelled: 'Cancelled',
}

const transitions: Record<string, string[]> = {
  draft:          ['sent', 'cancelled'],
  sent:           ['partially_paid', 'paid', 'cancelled'],
  partially_paid: ['paid', 'cancelled'],
}

const currentLabel = computed(() => statusLabels[props.currentStatus] ?? props.currentStatus)
const allowedTransitions = computed(() => transitions[props.currentStatus] ?? [])

function handleChange() {
  if (selected.value !== props.currentStatus) {
    emit('change', selected.value)
  }
}
</script>
