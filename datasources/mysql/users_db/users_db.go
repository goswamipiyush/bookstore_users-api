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
	//We have a valid database here to connect to

	// insert, err := sqldb.Query("INSERT INTO test VALUES ( 2, 'TEST' )")

	// // if there is an error inserting, handle it
	// if err != nil {
	// 	panic(err.Error())
	// }
	// be careful deferring Queries if you are using transactions

	//defer insert.Close()
}

func Insert() {
	insert, err := SqlDB.Query("INSERT INTO test VALUES ( 2, 'TEST' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	//be careful deferring Queries if you are using transactions

	defer insert.Close()
}
