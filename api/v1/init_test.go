package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/treemap/database"
	"github.com/ssst0n3/treemap/test/test_config"
)

var router *gin.Engine

func init() {
	test_config.InitConfig()
	database.InitConnector()
	router = gin.Default()
	InitRouter(router)
}
