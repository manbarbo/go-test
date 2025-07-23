import type { StockWithScore, StockInformation } from '../models/StockInformation'

const API_URL = import.meta.env.VITE_API_BASE_URL

export async function fetchTopRecommendations(top = 5): Promise<StockWithScore[]> {
  const response = await fetch(`${API_URL}/recommendation?top=${top}`)
  if (!response.ok) {
    throw new Error(`Failed to fetch recommendations: ${response.status}`)
  }
  return await response.json()
}

export async function fetchStocks(): Promise<StockInformation[]> {
  const res = await fetch(`${API_URL}/stocks`)
  if (!res.ok) {
    throw new Error('Failed to fetch stocks')
  }
  return res.json()
}