package v1

import (
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/treemap/model"
)

var NodeRelationResource = lightweight_api.Resource{
	Name:             "node_relation",
	TableName:        model.TableNameNodeRelation,
	BaseRelativePath: "/api/v1/node_relation",
	Model:            model.NodeRelation{},
	GuidFieldJsonTag: "",
}
