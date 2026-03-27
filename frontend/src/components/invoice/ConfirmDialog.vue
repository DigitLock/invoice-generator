<template>
  <Teleport to="body">
    <div v-if="visible" class="fixed inset-0 z-50 flex items-center justify-center">
      <div class="fixed inset-0 bg-black/50" @click="$emit('cancel')" />
      <div class="relative z-10 w-full max-w-md rounded-lg bg-white p-6 shadow-xl">
        <h2 class="text-lg font-semibold text-gray-900">Confirm PDF Generation</h2>
        <p class="mt-2 text-sm text-gray-600">
          Have you verified all invoice details? This will generate a PDF document.
        </p>
        <div class="mt-6 flex justify-end gap-3">
          <button type="button" @click="$emit('cancel')"
            class="rounded-md border border-gray-300 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">
            Cancel
          </button>
          <button type="button" @click="$emit('confirm')"
            class="rounded-md bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700">
            Generate PDF
          </button>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
defineProps<{
  visible: boolean
}>()

defineEmits<{
  confirm: []
  cancel: []
}>()
</script>
