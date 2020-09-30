package model

type NodeRelation struct {
	Parent uint `json:"parent"`
	Child  uint `json:"child"`
}

type NodeRelationWithId struct {
	Id uint `json:"id"`
	NodeRelation
}

const (
	TableNameNodeRelation = "node_relation"
	ColumnNameNodeRelationParent = "parent"
	ColumnNameNodeRelationChild = "child"
)