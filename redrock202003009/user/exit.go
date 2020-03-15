package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"net/http"
	"redrock20200309/token"
	"redrock20200309/sql"
)

func Exit(c *gin.Context)  {
	tokens := c.GetHeader("token")
	fmt.Println(tokens)
	//检查token
	id, username, err := token.CheckToken(tokens)
	user := string(id)
	fmt.Println(user)
	if err != nil{
		fmt.Println("err:",err)
		c.JSON(500,gin.H{"status":http.StatusBadRequest,"message":"token验证失败"})
		return
	}

	//数据库连接
	db := sql.Link()
	defer db.Close()
	var u User
	db.Where("id=?", id).First(&u)
	fmt.Println(id, u)
	if username == u.Username {
		fmt.Println("OK")
		db.Model(User{}).Where("id=?", id).Update(User{Add: false})

		//Redis数据库连接
		Rdb := sql.Client()
		err := Rdb.Del(user)
		if err != nil {
			fmt.Println("error:", err)
		}
		c.JSON(200,gin.H{"status":http.StatusOK,"message":"退出比赛成功"})
		return
	}else {
		fmt.Println("err",u.Username)
		c.JSON(400,gin.H{"status:":http.StatusBadRequest,"message":"用户不存在"})
		return
	}
}
