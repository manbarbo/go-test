<template>
  <div
    class="
      glass shadow-md rounded-xl p-4 border border-gray-200 cursor-pointer transition transform duration-300 ease-in-out
      hover:shadow-xl hover:-translate-y-1 hover:scale-105"
  >
    <h2 class="text-xl font-semibold text-fuchsia-500 min-h-[3.5rem] line-clamp-2 ">{{ stock.stock_information.company }}</h2>
    <p class="text-sm text-gray-200"><span class="text-green-500">Ticker:</span> {{ stock.stock_information.ticker }}</p>
    <p class="text-sm text-gray-200"><span class="text-green-500">Rating:</span> {{ stock.stock_information.rating_from }} → {{ stock.stock_information.rating_to }}</p>
    <p class="text-sm text-gray-200"><span class="text-green-500">Target:</span> {{ stock.stock_information.target_from }} → {{ stock.stock_information.target_to }}</p>
  </div>
</template>

<script setup lang="ts">
import type { StockInformation } from '@/models/StockInformation'

defineProps<{
  stock: StockInformation
}>()
</script>