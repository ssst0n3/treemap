package database

import (
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_db"
)

var Conn lightweight_db.Connector

func InitConnector() {
	//Conn = mysql.Conn()
	//lightweight_api.Conn = Conn
	//lightweight_api.InitConnector("mysql", lightweight_db.GetDsnFromEnvNormal())
	Conn = lightweight_api.Conn
}

func init() {
	InitConnector()
}
