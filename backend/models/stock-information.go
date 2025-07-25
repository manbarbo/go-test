package models

type StockInformation struct {
	ID         string `json:"id"`
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type StockWithScore struct {
	StockInformation `json:"stock_information"`
	Score            float64 `json:"score"`
}
