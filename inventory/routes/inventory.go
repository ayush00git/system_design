package routes

import (
	"inventory/handlers"

	"github.com/gin-gonic/gin"
)

func InventoryRoutes (route *gin.Engine, invHandler *handlers.InventoryHandler) {
	route.POST("/api/inventory", invHandler.PostAProduct)
}
