package v1

import (
	"github.com/ssst0n3/lightweight_api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListRootNodes(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, NodeResource.BaseRelativePath + "/root", nil)
	w := lightweight_api.ObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
}
