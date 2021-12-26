package models

import (
	"GinSql/dao"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"time"
)

type user struct {
	username string
	password string
	Balance int
}

func Register(Username string,Password string) bool{
	stm ,err:=dao.DB.Prepare("insert into table_test values (?,?,?) ")
	if err != nil {
		fmt.Println("insert failed ",err)
		return false
	}
	defer  stm.Close()
	_,err=stm.Exec(Username,Password,)
	if err != nil {
		fmt.Println("insert failed ",err)
		return false
	}
	return true
}
func Login(Username string,Password string)  bool{
	stm ,err :=dao.DB.Query("select * from table_test where username = ?",Username)
	if err != nil {
		fmt.Println("err ",err)
		return false
	}
	defer stm.Close()
	var u user
	for stm.Next(){
	err = stm.Scan(&u.username,&u.password)
	if err !=nil{
		fmt.Println("select err",err)
		}
	}
	if u.password==Password {
		return true

	}
	return false
}
func Charge(username string,money int,remark string) bool {
	stm,err:=dao.DB.Prepare("update table_test set balance=? where username=?;")
	if err !=nil {
		return false
	}
	_,err=stm.Exec(money,username)
	if err !=nil {
		return false
	}
	_,err=dao.DB.Exec("insert into table_logs values (?,?,?,?) ",username,money,time.Now(),remark)
	if err != nil {
		return false
	}
	return true
}
