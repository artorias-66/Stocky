package controllers

import (
	"net/http"
	"time"

	"stock-reward-system/src/models"
	"stock-reward-system/src/utils"

	"github.com/gin-gonic/gin"
)

type RewardController struct{}

func (rc *RewardController) PostReward(c *gin.Context) {
	var req models.RewardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if utils.CheckDuplicateReward(req.UserID, req.StockSymbol) {
		c.JSON(http.StatusConflict, gin.H{"error": "Duplicate reward event"})
		return
	}
	// TODO: Save to DB, update ledger, handle fees
	c.JSON(http.StatusOK, gin.H{"status": "success", "rewardId": "mock-uuid"})
}

func (rc *RewardController) GetTodayStocks(c *gin.Context) {
	userId := c.Param("userId")
	// TODO: Query today's rewards for user from DB
	c.JSON(http.StatusOK, gin.H{
		"userId": userId,
		"date":   time.Now().Format("2006-01-02"),
		"rewards": []gin.H{
			{"stockSymbol": "RELIANCE", "quantity": 5.0, "timestamp": "2025-10-03T09:00:00Z"},
		},
	})
}

func (rc *RewardController) GetHistoricalINR(c *gin.Context) {
	userId := c.Param("userId")
	// TODO: Query historical INR values for user from DB
	c.JSON(http.StatusOK, gin.H{
		"userId": userId,
		"history": []gin.H{
			{"date": "2025-10-01", "inrValue": 15000.25},
			{"date": "2025-10-02", "inrValue": 12000.75},
		},
	})
}

func (rc *RewardController) GetStats(c *gin.Context) {
	userId := c.Param("userId")
	// TODO: Query stats for user from DB
	c.JSON(http.StatusOK, gin.H{
		"userId": userId,
		"todayRewards": []gin.H{
			{"stockSymbol": "TCS", "quantity": 2.5},
		},
		"portfolioValueINR": 32000.50,
	})
}

func (rc *RewardController) GetPortfolio(c *gin.Context) {
	userId := c.Param("userId")
	// TODO: Query portfolio for user from DB
	c.JSON(http.StatusOK, gin.H{
		"userId": userId,
		"holdings": []gin.H{
			{"stockSymbol": "INFY", "quantity": 3.75, "currentINRValue": 6200.00},
		},
	})
}
