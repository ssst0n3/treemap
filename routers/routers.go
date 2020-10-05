package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/log"
	v1 "github.com/ssst0n3/treemap/api/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	{
		// cors
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowCredentials = true
		corsConfig.AllowOrigins = append(corsConfig.AllowOrigins, "http://127.0.0.1:12100")
		log.Logger.Info(corsConfig.AllowOrigins)
		router.Use(cors.New(corsConfig))
	}
	{
		// frontend
		router.Use(static.Serve("/", static.LocalFile("./dist", false)))
	}
	{
		// ping
		router.GET("/ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
	}
	{
		v1.InitRouter(router)
	}
	return router
}
