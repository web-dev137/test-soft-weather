package main

import (
	"github.com/web-dev137/test-soft-weather/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	api.Use(handler.Middleware)
	{
		api.GET("/arithmetic-sum/:exp", handler.ArithmeticSumHandler)
	}
	router.Run(":8080")
}
