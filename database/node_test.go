package database

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestListRootNodes(t *testing.T) {
	spew.Dump(ListRootNodes())
}

func TestTreeNodes(t *testing.T) {
	spew.Dump(TreeNodes(1))
}
