package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"redrock20200309/middleware"
	"redrock20200309/user"
)

func SetupRouter(app *gin.Engine)  {
	app.POST("/register", user.Register)//OK
	app.POST("/login", user.Login)//OK
	app.POST("/Create", user.Createvote)//OK

	app.Use(middleware.User)//token 拦截验证

	app.GET("/Find",  user.Find)//排行榜 OK
	app.POST("/Add",   user.Add)//OK
	app.POST("/Exit", user.Exit)//暂时不会
	app.POST("/Vote", user.Vote)//未设置每日三次机会
	fmt.Println("start")
}


