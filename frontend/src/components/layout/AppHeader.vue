<template>
  <header class="border-b border-gray-200 bg-white">
    <div class="mx-auto max-w-7xl px-4 py-4 flex items-center justify-between">
      <RouterLink to="/" class="text-xl font-semibold text-gray-900">
        Invoice Generator
      </RouterLink>
      <nav>
        <!-- Sign In placeholder for Stage 3 -->
        <button class="text-sm text-gray-500 cursor-not-allowed" disabled>
          Sign In
        </button>
      </nav>
    </div>
  </header>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>
