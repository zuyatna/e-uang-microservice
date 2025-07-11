package model

import "time"

type Transaction struct {
	ID              string `gorm:"primaryKey"`
	FromAccountID   string
	ToAccountID     string
	Amount          float64   `gorm:"not null"`
	TransactionType string    `gorm:"not null"` // e.g., "credit", "debit"
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
