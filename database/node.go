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
		"SELECT {.node}.* from {.node} JOIN {.node_relation} ON {.parent}=? and {.child}={.node}.id",
		awesome_libs.Dict{
			"node":          node.TableNameNode,
			"node_relation": model.TableNameNodeRelation,
			"parent":        model.ColumnNameNodeRelationParent,
			"child":         model.ColumnNameNodeRelationChild,
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
