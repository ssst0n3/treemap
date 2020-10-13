package v1

import (
	"bytes"
	"encoding/json"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/treemap/model/node"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListRootNodes(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, NodeResource.BaseRelativePath + "/root", nil)
	w := lightweight_api.ObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateChild(t *testing.T) {
	body, err := json.Marshal(node.CreateBody{
		Node:   node.Node{},
		Parent: 1,
	})
	assert.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPost, NodeResource.BaseRelativePath + "/child", bytes.NewReader(body))
	w := lightweight_api.ObjectOperate(req, router)
	log.Logger.Info(w.Body.String())
}