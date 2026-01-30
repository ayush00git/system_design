package routes

import (
	"github.com/gin-gonic/gin"
	"tinyurl/handlers"
)

func URLRoute(router *gin.Engine, urlHandler *handlers.URLHandler) {
	api := router.Group("/api/v1")
	{
		api.POST("/tinyurl", urlHandler.ToTinyURL)
	}
	router.GET("/:tinyId", urlHandler.HitTinyURL)
}
