<template>
  <div class="mx-auto max-w-7xl px-4 py-8 sm:py-16">
    <div class="grid grid-cols-1 items-center gap-8 sm:gap-12 lg:grid-cols-2">
      <!-- Left: invoice preview image -->
      <div class="flex flex-col items-center gap-2">
        <a href="/invoice-preview.pdf" target="_blank" rel="noopener"
          class="inline-block transition-transform duration-200 hover:scale-[1.02] hover:shadow-xl rounded-lg">
          <img
            src="/invoice-preview.png"
            alt="Sample invoice preview"
            class="max-h-[500px] w-auto rounded-lg border border-gray-200 shadow-lg cursor-pointer"
          />
        </a>
        <p class="text-sm text-gray-400">Click to preview sample</p>
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
          <RouterLink to="/login"
            class="rounded-md border border-gray-300 px-6 py-3 text-center text-gray-700 font-medium hover:bg-gray-50">
            Sign In
          </RouterLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>
