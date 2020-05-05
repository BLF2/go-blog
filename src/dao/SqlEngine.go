package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DbCon *sqlx.DB

func init() {
	if DbCon == nil {
		driverName := "mysql"
		dataSourceName := "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=true"

		DbCon = sqlx.MustConnect(driverName, dataSourceName)
	}
}
