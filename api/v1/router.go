package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/treemap/model/node"
)

func InitRouter(router *gin.Engine) {
	nodeGroup := router.Group(NodeResource.BaseRelativePath)
	{
		nodeGroup.GET("/", ListRootNodes)
		nodeGroup.GET("/:id", NodeResource.ShowResource)
		nodeGroup.POST("", func(context *gin.Context) {
			NodeResource.CreateResource(context, &node.Node{}, "", nil, nil)
		})
	}
}
