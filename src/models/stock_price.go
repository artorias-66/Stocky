package models

import "time"

type StockPrice struct {
	StockSymbol string    `json:"stockSymbol" gorm:"type:varchar(20);primaryKey"`
	PriceINR    float64   `json:"priceInr" gorm:"type:numeric(18,4)"`
	Timestamp   time.Time `json:"timestamp" gorm:"type:timestamp"`
}


