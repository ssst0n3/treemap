package database

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/treemap/model/node"
	"github.com/ssst0n3/treemap/model/node/nodeType"
)

func ListRootNodes() ([]node.WithId, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=%d", node.TableNameNode, node.ColumnNameNodeNodeType, nodeType.Root)
	nodeWithIdList := []node.WithId{}
	if err := Conn.OrmQueryRowsBind(&nodeWithIdList, query); err != nil {
		awesome_error.CheckErr(err)
		return nil, err
	}
	return nodeWithIdList, nil
}
