package database

import (
	"encoding/json"
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/treemap/model/node"
	"github.com/ssst0n3/treemap/model/node/nodeType"
)

func ListRootNodes() ([]node.WithId, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=%d", node.TableNameNode, node.ColumnNameNodeNodeType, nodeType.Root)
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

/*
TODO: generate cache after update
*/
func TreeNodes(id uint) (node.NodeRecursive, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", node.TableNameNode)
	var nodeRecursive node.NodeRecursive
	if err := Conn.OrmQueryRowBind(&nodeRecursive, query, id); err != nil {
		awesome_error.CheckErr(err)
		return node.NodeRecursive{}, err
	} else {
		var children []uint
		if err := json.Unmarshal([]byte(nodeRecursive.Children), &children); err != nil {
			awesome_error.CheckErr(err)
			return node.NodeRecursive{}, err
		} else {
			for _, id := range children {
				if subNode, err := TreeNodes(id); err != nil {
					awesome_error.CheckErr(err)
					return node.NodeRecursive{}, err
				} else {
					nodeRecursive.Sub = append(nodeRecursive.Sub, subNode)
				}
			}
			return nodeRecursive, nil
		}
	}
}
