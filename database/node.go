package database

import (
	"github.com/ssst0n3/awesome_libs"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/treemap/model"
	"github.com/ssst0n3/treemap/model/node"
	"github.com/ssst0n3/treemap/model/node/nodeType"
)

func ListRootNodes() ([]node.WithId, error) {
	query := awesome_libs.Format(
		"SELECT * FROM {.node} WHERE {.node_type}={.type}", awesome_libs.Dict{
			"node":      node.TableNameNode,
			"node_type": node.ColumnNameNodeNodeType,
			"type":      nodeType.Root,
		},
	)
	var nodeWithIdList []node.WithId
	if err := Conn.OrmQueryRowsBind(&nodeWithIdList, query); err != nil {
		awesome_error.CheckErr(err)
		return nil, err
	}
	if len(nodeWithIdList) == 0 {
		nodeWithIdList = []node.WithId{}
	}
	return nodeWithIdList, nil
}

func NodeChildren(id uint) ([]node.Recursive, error) {
	var nodeRecursive []node.Recursive

	query := awesome_libs.Format(
		"SELECT {.node}.* FROM {.node} JOIN {.node_relation} ON `{.parent}`=? AND `{.child}`={.node}.id ORDER BY `{.index}`",
		awesome_libs.Dict{
			"node":          node.TableNameNode,
			"node_relation": model.TableNameNodeRelation,
			"parent":        model.ColumnNameNodeRelationParent,
			"child":         model.ColumnNameNodeRelationChild,
			"index":         node.ColumnNameNodeIndex,
		},
	)

	if err := Conn.OrmQueryRowsBind(&nodeRecursive, query, id); err != nil {
		awesome_error.CheckErr(err)
		return nil, err
	}
	return nodeRecursive, nil
}

/*
TODO: generate cache after update
*/

func TreeNodes(id uint) (node.Recursive, error) {
	query := awesome_libs.Format(
		"select * from {.node} where id=?", awesome_libs.Dict{
			"node": node.TableNameNode,
		},
	)
	var nodeRecursive node.Recursive
	if err := Conn.OrmQueryRowBind(&nodeRecursive, query, id); err != nil {
		awesome_error.CheckErr(err)
		return node.Recursive{}, err
	} else {
		err := nodeRecursive.Tree(NodeChildren)
		awesome_error.CheckErr(err)
		return nodeRecursive, err
	}
}

func ExtremumIndexOfChildren(id uint, function string) (int, error) {
	query := awesome_libs.Format(
		"SELECT COALESCE({.func}(`{.index}`),0) FROM {.node} JOIN {.node_relation} ON {.node}.id={.child} AND {.parent}=?",
		awesome_libs.Dict{
			"func":          function,
			"index":         node.ColumnNameNodeIndex,
			"node":          node.TableNameNode,
			"node_relation": model.TableNameNodeRelation,
			"child":         model.ColumnNameNodeRelationChild,
			"parent":        model.ColumnNameNodeRelationParent,
		},
	)
	var maxIndex int
	return maxIndex, Conn.QueryRow(query, &maxIndex, id)
}

func MaxIndexOfChildren(id uint) (int, error) {
	return ExtremumIndexOfChildren(id, "MAX")
}

func MinIndexOfChildren(id uint) (int, error) {
	return ExtremumIndexOfChildren(id, "MIN")
}

func MoveNode(beforeId uint, nodeId uint) error {
	var beforeIndex int
	if err := Conn.OrmShowObjectOnePropertyByIdByReflectBind(node.TableNameNode, node.ColumnNameNodeIndex, int64(beforeId), &beforeIndex); err != nil {
		awesome_error.CheckErr(err)
		return err
	}
	var nodeIndex uint
	if err := Conn.OrmShowObjectOnePropertyByIdByReflectBind(node.TableNameNode, node.ColumnNameNodeIndex, int64(nodeId), &nodeIndex); err != nil {
		awesome_error.CheckErr(err)
		return err
	}
	if err := Conn.UpdateObjectSingleColumnById(int64(nodeId), node.TableNameNode, node.ColumnNameNodeIndex, beforeIndex); err != nil {
		awesome_error.CheckErr(err)
		return err
	}
	if err := Conn.UpdateObjectSingleColumnById(int64(beforeId), node.TableNameNode, node.ColumnNameNodeIndex, nodeIndex); err != nil {
		awesome_error.CheckErr(err)
		return err
	}
	return nil
}

func MoveExtremum(parent, nodeId uint, extremum string) (err error) {
	var extremumIndex int
	switch extremum {
	case "max":
		extremumIndex, err = MaxIndexOfChildren(parent)
		extremumIndex += 1
	case "min":
		extremumIndex, err = MinIndexOfChildren(parent)
		extremumIndex -= 1
	}
	if err != nil {
		awesome_error.CheckErr(err)
		return err
	}
	if err := Conn.UpdateObjectSingleColumnById(int64(nodeId), node.TableNameNode, node.ColumnNameNodeIndex, extremumIndex); err != nil {
		awesome_error.CheckErr(err)
		return err
	}
	return nil
}

func MoveFirst(parent, nodeId uint) error {
	return MoveExtremum(parent, nodeId, "min")
}

func MoveLast(parent uint, nodeId uint) error {
	return MoveExtremum(parent, nodeId, "max")
}
