package routes

import (
	"inventory/handlers"

	"github.com/gin-gonic/gin"
)

func OrderRoute(route *gin.Engine, orderHandler *handlers.OrderCollection) {
	route.POST("/api/order", orderHandler.PlaceOrder)
	route.GET("/", orderHandler.HealthRoute)
}
