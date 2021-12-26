package cmd

import (
	"GinSql/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Entrance() {
	router :=gin.Default()
	router.POST("/register",controller.Register)
	router.GET("/login",controller.Login)
	router.POST("/charge",Logger(),controller.Charge)
	router.Run()
}
func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {

		cookie,err:=context.Request.Cookie("user_cookie")
		if err != nil {
			context.Abort()
		}
		fmt.Println(cookie)

	}
}
