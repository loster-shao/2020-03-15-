package user

import (
	"fmt"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"net/http"
	"redrock20200309/sql"
	"redrock20200309/token"
	"strconv"
)

//type LoginForm struct {
//	ID       int    `json:"id"       binding:"required"`
//	Username string `json:"username" binding:"required"`
//	Password string `json:"password" binding:"required"`
//}

func Login(c *gin.Context){
	id0     := c.PostForm("user_id")
	id, err := strconv.Atoi(id0)
	if  err != nil {
		fmt.Println("err:",err)
	}
	password := c.PostForm("password")
	//MySQL数据库连接
	db :=sql.Link()
	defer db.Close()

	var u User
	db.Where("id=?", id).First(&u)
	tokens := token.Create(u.Username, id)//创建token

	//Redis数据库连接
	rdb := sql.Client()
	num, err := rdb.Set(id0,0,86400).Result()
	fmt.Println(num)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	if password == u.Password {
		c.JSON(200,gin.H{"status:":http.StatusOK,"message":"登录成功","token":tokens})
	}else {
		c.JSON(400,gin.H{"status:":http.StatusBadRequest,"message":"用户名或密码错误"})
	}
}



