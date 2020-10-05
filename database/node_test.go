package database

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/lightweight_db"
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
	node, err := TreeNodes(1)
	assert.NoError(t, err)
	spew.Dump(node)
}
