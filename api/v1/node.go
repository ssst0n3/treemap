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
	Model:            node.Node{},
	GuidFieldJsonTag: "",
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
	NodeResource.CreateResourceTemplate(c, func(modelPtr interface{}) (err error) {
		spew.Dump(modelPtr)
		n := modelPtr.(*node.Node)
		if n.NodeType != nodeType.Root {
			err = errors.New(fmt.Sprintf(consts.ErrorNodeTypeMustBeTypeRoot, nodeType.Root))
		}
		return
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
		NodeResource.CreateResourceTemplate(c, func(modelPtr interface{}) (err error) {
			n := modelPtr.(*node.Node)
			var maxIndex int
			maxIndex, err = database.MaxIndexOfChildren(createBody.Parent)
			n.Index = maxIndex + 1
			return
		}, func(id int64) (err error) {
			nodeRelation := model.NodeRelation{
				Parent: createBody.Parent,
				Child:  uint(id),
			}
			_, err = database.Conn.CreateObject(model.TableNameNodeRelation, nodeRelation)
			return
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
	action := c.Query("action")
	switch action {
	case node.ActionUpdateNodeName:
		NodeResource.UpdateResourceTemplate(c, node.UpdateNodeNameBody{}, nil)
	case node.ActionMoveNode:
		MoveNode(id, action, c)
	case node.ActionMoveNodeFirst:
		MoveNode(id, action, c)
	case node.ActionMoveNodeLast:
		MoveNode(id, action, c)
	case node.ActionUpdateNodeContent:
		NodeResource.UpdateResourceTemplate(c, node.UpdateNodeContentBody{}, nil)
	default:
		lightweight_api.HandleStatusBadRequestError(c, errors.New(fmt.Sprintf("action: %s not valid", action)))
	}
}

func MoveNode(id uint, action string, c *gin.Context) {
	var moveNodeBody node.MoveNodeBody
	if err := c.BindJSON(&moveNodeBody); err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	nodeRelation := model.NodeRelation{
		Parent: moveNodeBody.Parent,
		Child:  id,
	}

	if nodeRelationId, err := database.Conn.QueryIdByGuid(NodeRelationResource.TableName, model.ColumnNameNodeRelationChild, id); err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	} else {
		if err := database.Conn.UpdateObject(nodeRelationId, NodeRelationResource.TableName, &nodeRelation); err != nil {
			lightweight_api.HandleInternalServerError(c, err)
			return
		}
	}

	var err error
	switch action {
	case node.ActionMoveNode:
		err = database.MoveNode(moveNodeBody.BeforeId, id)
	case node.ActionMoveNodeFirst:
		err = database.MoveFirst(moveNodeBody.Parent, id)
	case node.ActionMoveNodeLast:
		err = database.MoveLast(moveNodeBody.Parent, id)
	}
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
}
