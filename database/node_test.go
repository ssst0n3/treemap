package database

import (
	"github.com/ssst0n3/treemap/test/test_config"
	"testing"
)

func TestTreeNodes(t *testing.T) {
	test_config.InitConfig()
	TreeNodes(1)
}
