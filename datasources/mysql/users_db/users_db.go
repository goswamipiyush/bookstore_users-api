package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	SqlDB *sql.DB
)

func GetConn() *sql.DB {
	return SqlDB
}

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", "root", "goswamip1yush", "users_db")
	var OpenErr error
	SqlDB, OpenErr = sql.Open("mysql", dataSourceName)
	if OpenErr != nil {
		panic(OpenErr)
	}
	//defer SqlDB.Close()
	log.Println("Database connection successful")
}
