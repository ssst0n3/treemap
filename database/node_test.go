package database

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/lightweight_db"
	"github.com/ssst0n3/treemap/model/node"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListRootNodes(t *testing.T) {
	spew.Dump(ListRootNodes())
}

func TestNodeChildren(t *testing.T) {
	lightweight_db.Logger.Level = logrus.DebugLevel
	nodes, err := NodeChildren(1)
	assert.NoError(t, err)
	spew.Dump(nodes)
}

func TestTreeNodes(t *testing.T) {
	lightweight_db.Logger.Level = logrus.DebugLevel
	n, err := TreeNodes(1)
	assert.NoError(t, err)
	spew.Dump(n)
}

func TestMaxIndexOfChildren(t *testing.T) {
	lightweight_db.Logger.Level = logrus.DebugLevel
	max, err := MaxIndexOfChildren(0)
	assert.NoError(t, err)
	log.Logger.Info(max)
}

func TestMoveNode(t *testing.T) {
	beforeId := 17
	nodeId := 18
	var beforeIndex uint
	var nodeIndex uint
	assert.NoError(t, Conn.OrmShowObjectOnePropertyBydIdByReflectBind(node.TableNameNode, node.ColumnNameNodeIndex, int64(beforeId), &beforeIndex))
	assert.NoError(t, Conn.OrmShowObjectOnePropertyBydIdByReflectBind(node.TableNameNode, node.ColumnNameNodeIndex, int64(nodeId), &nodeIndex))
	log.Logger.Info(beforeIndex, nodeIndex)
	assert.NoError(t, MoveNode(uint(beforeId), uint(nodeId)))
	assert.NoError(t, Conn.OrmShowObjectOnePropertyBydIdByReflectBind(node.TableNameNode, node.ColumnNameNodeIndex, int64(beforeId), &beforeIndex))
	assert.NoError(t, Conn.OrmShowObjectOnePropertyBydIdByReflectBind(node.TableNameNode, node.ColumnNameNodeIndex, int64(nodeId), &nodeIndex))
	log.Logger.Info(beforeIndex, nodeIndex)
}
