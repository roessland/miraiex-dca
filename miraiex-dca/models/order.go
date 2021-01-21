package models

type Order struct {
	ID         string  `json:"id"`
	ExternalID string  `json:"external_id"`
	OrderType  string  `json:"type"`
	Price      float64 `json:"price"`
	Amount     float64 `json:"amount"`
	Market     string  `json:"market"`
}
