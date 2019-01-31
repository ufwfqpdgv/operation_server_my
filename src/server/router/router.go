package routers

import (
	"server/api"
	"server/custom_middleware"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(custom_middleware.ServerHeader)
	// router.Use(custom_middleware.Log())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.GET("/api/v1/operation/activity", api.ActivityApi)
	return router
}
