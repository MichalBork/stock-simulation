package model

import (
	"time"
)

type Transaction struct {
	ID             int64     `json:"id"`
	UserID         int64     `json:"user_id"`
	BaseCurrency   string    `json:"base_currency"`
	TargetCurrency string    `json:"target_currency"`
	Amount         float64   `json:"amount"`
	ExchangeRate   float64   `json:"exchange_rate"` // Kurs wymiany waluty bazowej na walutę docelową
	Timestamp      time.Time `json:"timestamp"`
	Type           string    `json:"type"` // "buy" lub "sell"
}
