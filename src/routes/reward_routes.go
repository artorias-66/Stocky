package routes

import (
	"github.com/gin-gonic/gin"
	"stock-reward-system/src/controllers"
)

func SetupRoutes(router *gin.Engine, rewardController controllers.RewardController) {
	router.POST("/reward", rewardController.PostReward)
	router.GET("/today-stocks/:userId", rewardController.GetTodayStocks)
	router.GET("/historical-inr/:userId", rewardController.GetHistoricalINR)
	router.GET("/stats/:userId", rewardController.GetStats)
}

func RegisterRewardRoutes(r *gin.Engine) {
	rc := &controllers.RewardController{}
	r.POST("/reward", rc.PostReward)
	r.GET("/today-stocks/:userId", rc.GetTodayStocks)
	r.GET("/historical-inr/:userId", rc.GetHistoricalINR)
	r.GET("/stats/:userId", rc.GetStats)
	r.GET("/portfolio/:userId", rc.GetPortfolio)
}