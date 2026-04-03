<template>
  <div class="mx-auto max-w-7xl px-4 py-8 sm:py-16">
    <div class="flex flex-col items-center gap-8 sm:gap-12">
      <!-- Top: heading and action buttons -->
      <div class="space-y-6 text-center">
        <h1 class="text-3xl sm:text-4xl font-bold text-gray-900">Invoice Generator</h1>
        <p class="text-base sm:text-lg text-gray-600">
          Create professional invoices instantly. No account required.
        </p>
        <div class="flex flex-col gap-3 sm:flex-row justify-center">
          <RouterLink to="/invoice/new"
            class="rounded-md bg-blue-600 px-6 py-3 text-center text-white font-medium hover:bg-blue-700">
            Create Invoice
          </RouterLink>
          <RouterLink to="/login"
            class="rounded-md border border-gray-300 px-6 py-3 text-center text-gray-700 font-medium hover:bg-gray-50">
            Sign In
          </RouterLink>
        </div>
      </div>

      <!-- Bottom: invoice preview image -->
      <div class="flex flex-col items-center gap-2">
        <button type="button" @click="showModal = true"
          class="inline-block transition-transform duration-200 hover:scale-[1.02] hover:shadow-xl rounded-lg">
          <img
            src="/invoice-preview.png"
            alt="Sample invoice preview"
            class="max-h-[500px] w-auto rounded-lg border border-gray-200 shadow-lg cursor-pointer"
          />
        </button>
        <p class="text-sm text-gray-400">Click to preview sample</p>
      </div>
    </div>
  </div>

  <!-- Preview modal -->
  <Teleport to="body">
    <div v-if="showModal" class="fixed inset-0 z-50 flex items-center justify-center p-4"
      @click.self="showModal = false">
      <div class="fixed inset-0 bg-black/60" @click="showModal = false"></div>
      <div class="relative z-10 max-h-[90vh] max-w-3xl">
        <button type="button" @click="showModal = false"
          class="absolute -top-3 -right-3 flex h-8 w-8 items-center justify-center rounded-full bg-white text-gray-600 shadow-lg hover:text-gray-900">
          &times;
        </button>
        <img
          src="/invoice-preview.png"
          alt="Sample invoice preview"
          class="max-h-[90vh] w-auto rounded-lg shadow-2xl"
        />
      </div>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { RouterLink } from 'vue-router'

const showModal = ref(false)
</script>
