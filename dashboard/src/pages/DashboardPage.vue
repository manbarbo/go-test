<template>
  <div class="flex flex-col min-h-full min-w-full p-6 gap-6 text-white">
    <div class="h-1/3 glass rounded-xl shadow p-4">
      <TopRecommendations />
    </div>
    <div class="h-2/3 glass rounded-xl shadow p-4">
      <StockTable />
    </div>
  </div>
</template>

<script setup>
    import StockTable from '../components/StockTable.vue'
    import TopRecommendations from '../components/TopRecommendations.vue'
</script>