package test_config

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/secret/consts"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_db"
	"github.com/ssst0n3/treemap/config"
	"os"
)

func InitConfig() {
	pathExamplePassword := config.ProjectDir + "/secrets/treemap_db_password"
	awesome_error.CheckFatal(os.Setenv(lightweight_api.EnvDriverName, "mysql"))
	awesome_error.CheckFatal(os.Setenv(lightweight_db.EnvDbName, "treemap"))
	awesome_error.CheckFatal(os.Setenv(lightweight_db.EnvDbHost, "127.0.0.1"))
	awesome_error.CheckFatal(os.Setenv(lightweight_db.EnvDbPort, "12306"))
	awesome_error.CheckFatal(os.Setenv(lightweight_db.EnvDbUser, "treemap"))
	awesome_error.CheckFatal(os.Setenv(lightweight_db.EnvDbPasswordFile, pathExamplePassword))
	awesome_error.CheckFatal(os.Setenv(consts.EnvDirSecret, "/tmp/secret"))
}
