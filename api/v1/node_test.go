package v1

import (
	"testing"
)

func TestListRootNodes(t *testing.T) {
	NodeResource.TestResourceListResource(t, router)
}
