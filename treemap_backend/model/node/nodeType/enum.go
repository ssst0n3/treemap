package nodeType

type Enum uint

const (
	Root Enum = iota // default value
	ChildNode
	LeafNode
)
