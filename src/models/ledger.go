package models

import (
	"time"
	"github.com/google/uuid"
)

type LedgerEntry struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	RewardID    uuid.UUID `json:"rewardId" gorm:"type:uuid;index"`
	EntryType   string    `json:"entryType" gorm:"type:varchar(20)"` // STOCK, CASH, FEE
	StockSymbol string    `json:"stockSymbol" gorm:"type:varchar(20);index"`
	Quantity    float64   `json:"quantity" gorm:"type:numeric(18,6)"`
	AmountINR   float64   `json:"amountInr" gorm:"type:numeric(18,4)"`
	Description string    `json:"description" gorm:"type:varchar(100)"`
	Timestamp   time.Time `json:"timestamp" gorm:"type:timestamp"`
}


