package models

type StockInformationResponse struct {
	Items    []StockInformation `json:"items"`
	NextPage string             `json:"next_page"`
}
