package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_db"
	"github.com/ssst0n3/lightweight_db/example/mysql"
)

var Conn lightweight_db.Connector

func InitConnector() {
	Conn = mysql.Conn()
	lightweight_api.Conn = Conn
	lightweight_api.InitConnector("mysql", lightweight_db.GetDsnFromEnvNormal())
}
