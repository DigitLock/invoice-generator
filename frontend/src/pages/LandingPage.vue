<template>
  <div class="mx-auto max-w-7xl px-4 py-8 sm:py-16">
    <div class="grid grid-cols-1 items-center gap-8 sm:gap-12 lg:grid-cols-2">
      <!-- Left: invoice preview image -->
      <div class="flex justify-center">
        <div class="h-64 w-48 sm:h-96 sm:w-72 rounded-lg border border-gray-200 bg-gray-50 flex items-center justify-center text-gray-400 text-sm">
          Invoice Preview
        </div>
      </div>

      <!-- Right: action buttons -->
      <div class="space-y-6 text-center lg:text-left">
        <h1 class="text-3xl sm:text-4xl font-bold text-gray-900">Invoice Generator</h1>
        <p class="text-base sm:text-lg text-gray-600">
          Create professional invoices instantly. No account required.
        </p>
        <div class="flex flex-col gap-3 sm:flex-row lg:justify-start justify-center">
          <RouterLink to="/invoice/new"
            class="rounded-md bg-blue-600 px-6 py-3 text-center text-white font-medium hover:bg-blue-700">
            Create Invoice
          </RouterLink>
          <button disabled
            class="rounded-md border border-gray-300 px-6 py-3 text-gray-400 font-medium cursor-not-allowed">
            Sign In
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>
