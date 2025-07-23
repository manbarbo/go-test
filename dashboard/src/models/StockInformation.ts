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