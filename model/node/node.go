package node

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/treemap/model/node/nodeType"
)

type Node struct {
	Index    uint          `json:"index"`
	Name     string        `json:"name"`
	NodeType nodeType.Enum `json:"node_type"`

	// common:
	Description string `json:"description"`

	// only for leaf:
	Leaf
}

type WithId struct {
	Id uint `json:"id"`
	Node
}

type Recursive struct {
	WithId
	Sub []Recursive
}

type CreateBody struct {
	Node
	Parent uint `json:"parent"`
}

type MoveNodeBody struct {
	Parent uint `json:"parent"`
	Index  uint `json:"index"`
}

const (
	TableNameNode          = "node"
	ColumnNameNodeName     = "name"
	ColumnNameNodeNodeType = "node_type"
	ColumnNameNodeIndex    = "index"
)

func (r *Recursive) Tree(getChildrenFunc func(id uint) ([]Recursive, error)) (err error) {
	r.Sub, err = getChildrenFunc(r.Id)
	for index, sub := range r.Sub {
		if err := sub.Tree(getChildrenFunc); err != nil {
			awesome_error.CheckErr(err)
			return err
		}
		r.Sub[index] = sub
	}
	return nil
}
