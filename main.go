package main

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/treemap/database"
	"github.com/ssst0n3/treemap/routers"
)

func main() {
	database.InitConnector()
	router := routers.InitRouter()
	awesome_error.CheckFatal(router.Run(":12000"))
}
