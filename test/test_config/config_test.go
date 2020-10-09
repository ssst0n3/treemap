package test_config

import (
	"github.com/ssst0n3/lightweight_api"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {
	InitConfig()
	assert.Equal(t, "mysql", os.Getenv(lightweight_api.EnvDriverName))
}
