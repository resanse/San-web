package service

import (
	"GinSql/models"
	"github.com/gin-gonic/gin"
)
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Balance int `json:"balance"`
}
var U User
func Register(ctx *gin.Context) bool {
	err:=ctx.ShouldBindJSON(&U)
	if err != nil {
		return false
	}
	res:=models.Register(U.Username,U.Password)
	return res
}
func Login(ctx *gin.Context) bool {
	err:=ctx.ShouldBindJSON(&U)
	if err != nil {
		return false
	}
	res:=models.Login(U.Username,U.Password)
	//return res
	if res{
		ctx.SetCookie("user_cookie", U.Username, 1000, "/", "localhost", false, true)
	}
	return res
}
func Charge(ctx *gin.Context)bool  {
	err:=ctx.ShouldBindJSON(&U)
	mark:=ctx.PostForm("mark")
	if err!=nil {
		return false
	}
	res :=models.Charge(U.Username,U.Balance,mark)
	return res
}
