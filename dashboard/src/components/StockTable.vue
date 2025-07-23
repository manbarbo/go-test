<template>
  <div class="space-y-4">
    <div class="flex flex-wrap items-center justify-between gap-2 glass p-1 w-full">
      <div class="flex flex-wrap gap-2">
        <input v-model="filters.ticker" placeholder="Ticker" class="input glass p-1" />
        <input v-model="filters.company" placeholder="Company" class="input glass p-1" />
        <input v-model="filters.brokerage" placeholder="Brokerage" class="input glass p-1" />
      </div>
      <div>
        <button @click="loadStocks" class="btn">Apply Filters</button>
      </div>
    </div>

    <DataTable
      :value="stocks"
      :loading="loading"
      :paginator="true"
      :rows="limit"
      :totalRecords="totalRecords"
      :lazy="true"
      @page="onPage"
      @sort="onSort"
      responsiveLayout="scroll"
      stripedRows
      class="rounded"
    >
      <Column field="ticker" header="Ticker" style="width: 10%" sortable></Column>
      <Column field="company" header="Company" style="width: 20%" sortable></Column>
      <Column field="brokerage" header="Brokerage" style="width: 20%" sortable></Column>
      <Column header="Target" style="width: 20%">
        <template #body="slotProps">
          {{ slotProps.data.target_from }} → {{ slotProps.data.target_to }}
        </template>
      </Column>
      <Column field="rating_from" header="Rating From" style="width: 10%" sortable></Column>
      <Column field="rating_to" header="Rating To" style="width: 10%" sortable></Column>
      <Column field="time" header="Time" style="width: 10%" sortable>
        <template #body="slotProps">
          {{ formatDate(slotProps.data.time) }}
      </template>
      </Column>
    </DataTable>
  </div>
</template>

<script setup lang="ts">
  import { ref, reactive, onMounted } from 'vue'
  import DataTable from 'primevue/datatable'
  import Column from 'primevue/column'
  import { fetchStocks } from '../services/stockService'
  import type { StockInformation } from '../model/StockInformation'

  const stocks = ref<StockInformation[]>([])
  const totalRecords = ref(0)
  const loading = ref(false)
  const offset = ref(0)
  const limit = ref(5)

  const sort_by = ref('time')
  const order = ref<'ASC' | 'DESC'>('DESC')

  const filters = reactive({
    company: '',
    brokerage: '',
    ticker: '',
    action: ''
  })

  async function loadStocks() {
    loading.value = true
    try {
      const result = await fetchStocks({
        ...filters,
        limit: limit.value,
        offset: offset.value,
        sort_by: sort_by.value,
        order: order.value
      })
      stocks.value = result
      totalRecords.value = 1000 // Puedes mejorar esto si tu API retorna `totalCount`
    } catch (err) {
      console.error(err)
    } finally {
      loading.value = false
    }
  }

  function onPage(event: any) {
    offset.value = event.first
    limit.value = event.rows
    loadStocks()
  }

  function onSort(event: any) {
    sort_by.value = event.sortField
    order.value = event.sortOrder === 1 ? 'ASC' : 'DESC'
    loadStocks()
  }

  function formatDate(dateString: string): string {
    const date = new Date(dateString)
    return date.toLocaleString('en-GB', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit'
    }).replace(',', '')
  }
  onMounted(loadStocks)
</script>

<style scoped>
  .input {
    @apply p-2 border rounded text-sm;
  }
  .btn {
    @apply bg-blue-600 text-white px-3 py-1 rounded;
  }
</style>