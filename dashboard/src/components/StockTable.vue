<template>
  <DataTable
    :value="stocks"
    :paginator="true"
    :rows="5"
    :loading="loading"
    responsiveLayout="scroll"
    stripedRows
    class="rounded-lg"
  >
    <Column field="company" header="Company" style="width: 20%"></Column>
    <Column field="ticker" header="Ticker" style="width: 20%"></Column>
    <Column field="target_from" header="Target From" style="width: 5%"></Column>
    <Column field="target_to" header="Target To" style="width: 5%"></Column>
    <Column field="action" header="Action" style="width: 10%"></Column>
    <Column field="brokerage" header="Brokerage" style="width: 20%"></Column>
    <Column field="rating_from" header="Rating From" style="width: 5%"></Column>
    <Column field="rating_to" header="Rating To" style="width: 5%"></Column>
  </DataTable>
</template>

<script setup lang="ts">
  import { ref, onMounted } from 'vue'
  import DataTable from 'primevue/datatable'
  import Column from 'primevue/column'
  import type { StockInformation } from '@/interfaces/stock'
  import { fetchStocks } from '../services/StockService'

  const stocks = ref<StockInformation[]>([])
  const loading = ref(true)

  onMounted(async () => {
    try {
      stocks.value = await fetchStocks()
    } catch (err) {
      console.error(err)
    } finally {
      loading.value = false
    }
  })
</script>