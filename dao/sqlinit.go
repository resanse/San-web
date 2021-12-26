package dao

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)
 var DB *sql.DB
func MysqlInit() *sql.DB {
	db, err := sql.Open("mysql", "root:Simple2002@tcp(127.0.0.1:3306)/sql_test")
	if err != nil {
		fmt.Println("open sql failed",err)
	}
	err =db.Ping()
	if err != nil {
		fmt.Println("open failed ",err)
	}
	DB=db
	return DB
}
