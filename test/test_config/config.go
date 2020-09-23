package test_config

import (
	"github.com/ssst0n3/treemap/config"
	"os"
)

func InitConfig() {
	pathExamplePassword := config.ProjectDir + "/secrets/treemap_db_password"
	os.Setenv("DB_NAME", "treemap")
	os.Setenv("DB_HOST", "172.26.0.3")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD_FILE", pathExamplePassword)
}
