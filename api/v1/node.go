package v1

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/treemap/consts"
	"github.com/ssst0n3/treemap/database"
	"github.com/ssst0n3/treemap/model"
	"github.com/ssst0n3/treemap/model/node"
	"github.com/ssst0n3/treemap/model/node/nodeType"
	"io/ioutil"
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

func CreateRoot(c *gin.Context) {
	NodeResource.CreateResource(c, &node.Node{}, "", func(modelPtr interface{}) {
		spew.Dump(modelPtr)
		n := modelPtr.(*node.Node)
		if n.NodeType != nodeType.Root {
			lightweight_api.HandleStatusBadRequestError(c, errors.New(fmt.Sprintf(consts.ErrorNodeTypeMustBeTypeRoot, nodeType.Root)))
			return
		}
	}, nil)
}

func CreateChild(c *gin.Context) {
	if body, err := ioutil.ReadAll(c.Request.Body); err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	} else {
		var createBody node.CreateBody
		if err := json.Unmarshal(body, &createBody); err != nil {
			lightweight_api.HandleStatusBadRequestError(c, err)
			return
		}
		if err := NodeResource.MustResourceExistsById(c, int64(createBody.Parent)); err != nil {
			lightweight_api.HandleStatusBadRequestError(c, err)
			return
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		NodeResource.CreateResource(c, &node.Node{}, "", func(modelPtr interface{}) {
			n := modelPtr.(*node.Node)
			maxIndex, err := database.MaxIndexOfChildren(createBody.Parent)
			if err != nil {
				lightweight_api.HandleInternalServerError(c, err)
				return
			}
			n.Index = maxIndex + 1
		}, func(id int64) {
			nodeRelation := model.NodeRelation{
				Parent: createBody.Parent,
				Child:  uint(id),
			}
			if _, err := database.Conn.CreateObject(model.TableNameNodeRelation, nodeRelation); err != nil {
				lightweight_api.HandleInternalServerError(c, err)
				return
			}
		})
	}
}

func UpdateNode(c *gin.Context) {
	idInt64, err := NodeResource.MustResourceExistsByIdAutoParseParam(c)
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	id := uint(idInt64)
	action := c.Param("action")
	switch action {
	case consts.ActionUpdateName:
		UpdateName(id, c)
	case consts.ActionMoveNode:
		MoveNode(id, c)
	}
}

func UpdateName(id uint, c *gin.Context) {
	var n node.Node
	if err := c.BindJSON(&n); err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	if err := database.Conn.UpdateObjectSingleColumnById(int64(id), node.TableNameNode, node.ColumnNameNodeName, n.Name); err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
}

func MoveNode(id uint, c *gin.Context) {
	var moveNodeBody node.MoveNodeBody
	if err := c.BindJSON(&moveNodeBody); err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	nodeRelation := model.NodeRelation{
		Parent: moveNodeBody.Parent,
		Child:  id,
	}
	NodeRelationResource.UpdateResource(c, &nodeRelation, "", nil)
	if err := database.Conn.UpdateObjectSingleColumnById(int64(id), node.TableNameNode, node.ColumnNameNodeIndex, moveNodeBody.Index); err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
}
