package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.Default()

	wTools := router.Group("/wTools")
	{
		InitWatermark(wTools)
	}
	return router
}
