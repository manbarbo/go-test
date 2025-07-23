<template>
  <div class="grid min-h-full grid-rows-[auto_1fr_auto] gap-4">
    <h1 class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-green-500 to-purple-500">
      Invest Now!
    </h1>

    <div v-if="loading" class="grid place-items-center text-gray-500">
      <p>Loading recommendations...</p>
    </div>

    <div v-else-if="error" class="grid place-items-center text-red-500">
      <p>Failed to load recommendations. Please try again later.</p>
    </div>

    <div v-else class="grid grid-cols-5 gap-4">
      <RecommendationCard
        v-for="stock in recommendations"
        :key="stock.stock_information.id"
        :stock="stock"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import RecommendationCard from './RecommendationCard.vue'
import type { StockWithScore } from '../models/StockInformation'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL
const recommendations = ref<StockWithScore[]>([])
const loading = ref(true)
const error = ref(false)

onMounted(async () => {
  try {
    const res = await fetch(`${API_BASE_URL}/recommendation?top=5`)
    recommendations.value = await res.json()
    console.log(recommendations)
  } catch (err) {
    console.error(err)
    error.value = true
  } finally {
    loading.value = false
  }
})

</script>