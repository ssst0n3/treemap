package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api/middleware"
)

func InitRouter(router *gin.Engine) {
	nodeGroup := router.Group(NodeResource.BaseRelativePath)
	{
		nodeGroup.GET("/root", ListRootNodes)
		nodeGroup.GET("/root/:id", TreeNodes)
		nodeGroup.GET("/children/:id", NodeResource.ShowResource)
		nodeGroup.POST("/root", CreateRoot)
		nodeGroup.POST("/child", CreateChild)
		nodeGroup.DELETE("/child/:id", NodeResource.DeleteResource)
		nodeGroup.PUT("/node/:id", UpdateNode)
	}
}

func InitRouterWithJwtUser(router *gin.Engine) {
	nodeGroup := router.Group(NodeResource.BaseRelativePath, middleware.JwtUser())
	{
		nodeGroup.GET("/root", ListRootNodes)
		nodeGroup.GET("/root/:id", TreeNodes)
		nodeGroup.GET("/children/:id", NodeResource.ShowResource)
		nodeGroup.POST("/root", CreateRoot)
		nodeGroup.POST("/child", CreateChild)
		nodeGroup.DELETE("/child/:id", NodeResource.DeleteResource)
		nodeGroup.PUT("/node/:id", UpdateNode)
	}
}
