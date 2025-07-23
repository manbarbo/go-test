import type { StockWithScore, StockInformation, StockQueryParams } from '../models/StockInformation'

const API_URL = import.meta.env.VITE_API_BASE_URL

export async function fetchTopRecommendations(top = 5): Promise<StockWithScore[]> {
  const response = await fetch(`${API_URL}/recommendation?top=${top}`)
  if (!response.ok) {
    throw new Error(`Failed to fetch recommendations: ${response.status}`)
  }
  return await response.json()
}

export async function fetchStocks(params: StockQueryParams): Promise<StockInformation[]> {
  const url = new URL(`${import.meta.env.VITE_API_BASE_URL}/stocks`)

  Object.entries(params).forEach(([key, value]) => {
    if (value !== undefined && value !== '') {
      url.searchParams.append(key, value.toString())
    }
  })

  const response = await fetch(url.toString())
  if (!response.ok) {
    throw new Error('Error fetching stocks')
  }

  return await response.json()
}