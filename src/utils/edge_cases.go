package utils

import (
	"errors"
	"sync"
)

var rewardEvents = struct {
	m map[string]bool
	sync.RWMutex
}{m: make(map[string]bool)}

// CheckDuplicateReward checks if a reward event for a user and stock symbol already exists.
func CheckDuplicateReward(userID string, stockSymbol string) bool {
	rewardEvents.RLock()
	_, exists := rewardEvents.m[userID+stockSymbol]
	rewardEvents.RUnlock()
	if exists {
		return true
	}
	rewardEvents.Lock()
	rewardEvents.m[userID+stockSymbol] = true
	rewardEvents.Unlock()
	return false
}

// HandleStockEvents manages stock splits, mergers, or delisting.
func HandleStockEvents(stockSymbol string, eventType string) error {
	// Implement logic for handling stock events
	switch eventType {
	case "split":
		// Handle stock split logic
	case "merger":
		// Handle merger logic
	case "delist":
		// Handle delisting logic
	default:
		return errors.New("unknown stock event type")
	}
	return nil
}

// HandlePriceAPIDowntime manages scenarios where the price API is down.
func HandlePriceAPIDowntime() error {
	// Implement fallback logic for price API downtime
	return errors.New("price API is currently unavailable")
}
