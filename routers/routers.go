package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ssst0n3/treemap/api/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	v1.InitRouter(router)
	return router
}
