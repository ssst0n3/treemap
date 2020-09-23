package node

import "github.com/ssst0n3/treemap/model/node/leafType"

type Leaf struct {
	LeafType leafType.Enum `json:"leaf_type"`
	Content  string        `json:"content"`
}
