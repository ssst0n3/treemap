package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/treemap/database"
	"github.com/ssst0n3/treemap/model/node"
	"net/http"
)

var NodeResource = lightweight_api.Resource{
	Name:             "node",
	TableName:        node.TableNameNode,
	BaseRelativePath: "/api/v1/node",
	Model:            node.WithId{},
}

func ListRootNodes(c *gin.Context) {
	if nodeWithIdList, err := database.ListRootNodes(); err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, nodeWithIdList)
	}
}

func TreeNodes(c *gin.Context) {
	id, err := NodeResource.MustResourceExistsByIdAutoParseParam(c)
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	if nodeRecursive, err := database.TreeNodes(uint(id)); err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, nodeRecursive)
	}
}
