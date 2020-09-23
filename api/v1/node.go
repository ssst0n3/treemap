package v1

import (
	"encoding/json"
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

func CreateNode(c *gin.Context) {
	var body node.CreateBody
	if err := c.BindJSON(&body); err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	if err := NodeResource.MustResourceExistsById(c, int64(body.Parent)); err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	NodeResource.CreateResource(c, &node.Node{}, "", nil, func(id int64) {
		var children string
		if err := database.Conn.OrmShowObjectOnePropertyByIdUsingJsonBind(node.TableNameNode, node.ColumnNameNodeChildren, int64(body.Parent), &children); err != nil {
			lightweight_api.HandleInternalServerError(c, err)
			return
		} else {
			var childrenIdList []uint
			if err := json.Unmarshal([]byte(children), &childrenIdList); err != nil {
				lightweight_api.HandleInternalServerError(c, err)
				return
			} else {
				childrenIdList = append(childrenIdList, uint(id))
				if newChildren, err := json.Marshal(childrenIdList); err != nil {
					lightweight_api.HandleInternalServerError(c, err)
					return
				} else {
					if err := database.Conn.UpdateObjectSingleColumnById(int64(body.Parent), node.TableNameNode, node.ColumnNameNodeChildren, newChildren); err != nil {
						lightweight_api.HandleInternalServerError(c, err)
						return
					}
				}
			}
		}
	})
}
