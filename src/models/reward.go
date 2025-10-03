package models

import (
	"time"

	"github.com/google/uuid"
)

type Reward struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID `json:"userId" gorm:"type:uuid;index"`
	StockSymbol string    `json:"stockSymbol" gorm:"type:varchar(20);index"`
	Quantity    float64   `json:"quantity" gorm:"type:numeric(18,6)"`
	Timestamp   time.Time `json:"timestamp" gorm:"type:timestamp"`
}

type RewardRequest struct {
	UserID      string    `json:"userId"`
	StockSymbol string    `json:"stockSymbol"`
	Quantity    float64   `json:"quantity"`
	Timestamp   time.Time `json:"timestamp"`
}