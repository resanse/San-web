package main

import (
	"GinSql/cmd"
	"GinSql/dao"
)
func main() {
	dao.MysqlInit()
	cmd.Entrance()
}