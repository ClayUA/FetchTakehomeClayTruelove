package main

import (
	"github.com/ClayUA/FetchTakehomeClayTruelove/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/add", endpoints.AddHandler)
	router.GET("/balance", endpoints.BalanceHandler)
	router.POST("/spend", endpoints.SpendHandler)
	router.Run(":8000")
}
