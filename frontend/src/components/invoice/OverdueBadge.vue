<template>
  <span v-if="isOverdue" class="inline-block rounded-full bg-red-100 text-red-700 px-2.5 py-0.5 text-xs font-medium">
    Overdue
  </span>
</template>

<script setup lang="ts">
defineProps<{ isOverdue: boolean }>()
</script>
