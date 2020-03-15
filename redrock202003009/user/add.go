package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/jinzhu/gorm"
	"net/http"
	"redrock20200309/token"
	"redrock20200309/sql"
)

//报名参赛 OK
func Add(c *gin.Context)  {
	vname := c.PostForm("vote_name")
	tokens := c.GetHeader("token")
	fmt.Println(tokens)

	//检查token
	id, username, err := token.CheckToken(tokens)
	user := string(id)
	if err != nil{
		fmt.Println("err:",err)
		c.JSON(500,gin.H{"status":http.StatusBadRequest,"message":"token验证失败"})
		return
	}

	//MySQL数据库连接
	db := sql.Link()
	defer db.Close()
	var u User
	db.Where("id=?", id).First(&u)
	fmt.Println(id, u)
	if username == u.Username {
		fmt.Println("OK")
		db.Model(User{}).Where("id=?", id).Update(User{Add: true})

		//redis数据库连接
		Rdb := sql.Client()
		users := []redis.Z{
			redis.Z{0,user},
		}
		err := Rdb.ZAdd(vname, users...)
		if err != nil {
			fmt.Println("error:", err)
		}

		//JSON
		c.JSON(200,gin.H{"status":http.StatusOK,"message":"报名参赛成功"})
		return
	}else {
		fmt.Println("err",u.Username)
		c.JSON(400,gin.H{"status:":http.StatusBadRequest,"message":"用户不存在"})
		return
	}
}