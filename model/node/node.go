package node

import "github.com/ssst0n3/treemap/model/node/nodeType"

type Node struct {
	Name     string        `json:"name"`
	NodeType nodeType.Enum `json:"node_type"`

	// common:
	Description string `json:"description"`
	Children    string `json:"children"` // [1,2,3]

	// only for leaf:
	Leaf
}

type WithId struct {
	Id uint `json:"id"`
	Node
}

const (
	TableNameNode = "node"
	ColumnNameNodeNodeType = "node_type"
)
