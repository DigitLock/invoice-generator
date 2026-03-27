<template>
  <span :class="classes" class="inline-block rounded-full px-2.5 py-0.5 text-xs font-medium">
    {{ label }}
  </span>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{ status: string }>()

const statusMap: Record<string, { label: string; classes: string }> = {
  draft:          { label: 'Draft',          classes: 'bg-gray-100 text-gray-700' },
  sent:           { label: 'Sent',           classes: 'bg-blue-100 text-blue-700' },
  partially_paid: { label: 'Partially Paid', classes: 'bg-orange-100 text-orange-700' },
  paid:           { label: 'Paid',           classes: 'bg-green-100 text-green-700' },
  cancelled:      { label: 'Cancelled',      classes: 'bg-red-100 text-red-700' },
}

const label = computed(() => statusMap[props.status]?.label ?? props.status)
const classes = computed(() => statusMap[props.status]?.classes ?? 'bg-gray-100 text-gray-700')
</script>
