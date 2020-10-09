package database

import (
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/middleware"
	"github.com/ssst0n3/treemap/test/test_config"
)

func init() {
	test_config.InitConfig()
	lightweight_api.InitConnector()
	InitConnector()
	middleware.InitJwtKey()
}
