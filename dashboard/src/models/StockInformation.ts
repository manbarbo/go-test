export interface StockInformation {
  id: string
  ticker: string
  target_from: string
  target_to: string
  company: string
  action: string
  brokerage: string
  rating_from: string
  rating_to: string
  time: string
}

export interface StockWithScore {
    stock_information: StockInformation
    score: number
}

export interface StockQueryParams {
  company?: string
  brokerage?: string
  ticker?: string
  action?: string
  sort_by?: string
  order?: 'ASC' | 'DESC'
  limit?: number
  offset?: number
}