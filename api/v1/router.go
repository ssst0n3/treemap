package v1

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	nodeGroup := router.Group(NodeResource.BaseRelativePath)
	{
		nodeGroup.GET("/", ListRootNodes)
		nodeGroup.GET("/:id", NodeResource.ShowResource)
		nodeGroup.POST("", CreateNode)
	}
}
