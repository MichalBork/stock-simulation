package model

type Wallet struct {
	UserID   int64   `json:"user_id"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
}
