package database

import "github.com/ssst0n3/treemap/test/test_config"

func init() {
	test_config.InitConfig()
	InitConnector()
}
