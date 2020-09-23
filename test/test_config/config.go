package test_config

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/treemap/config"
	"os"
)

func InitConfig() {
	pathExamplePassword := config.ProjectDir + "/secrets/treemap_db_password"
	awesome_error.CheckFatal(os.Setenv("DB_NAME", "treemap"))
	awesome_error.CheckFatal(os.Setenv("DB_HOST", "172.31.0.2"))
	awesome_error.CheckFatal(os.Setenv("DB_USER", "treemap"))
	awesome_error.CheckFatal(os.Setenv("DB_PASSWORD_FILE", pathExamplePassword))
}
