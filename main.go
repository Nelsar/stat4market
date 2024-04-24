package main

import (
	"github.com/gin-gonic/gin"
	"stat4market.com/handlers"
)

func main() {
	r := initHttp()
	r.Run(":8090")
}

func initHttp() *gin.Engine {
	route := gin.Default()
	route.POST("/createEvent", handlers.CreateEvent)

	return route
}
