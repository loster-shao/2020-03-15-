package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"redrock20200309/model"
	"redrock20200309/sql"
)

var Vote_id string

//创建比赛 OK
func Createvote(c *gin.Context)  {
	//接口
	vname     := c.PostForm("vote_name")
	starttime := c.PostForm("starttime")
	endtime   := c.PostForm("endtime")

	//MySQL数据库连接
	db := sql.Link()
	db.AutoMigrate(&model.Vote{})
	db.Create(&model.Vote{ Vname: vname, Starttime: starttime, Endtime: endtime})
	c.JSON(200, gin.H{"status": http.StatusOK, "message": "赛事创建成功"})

	//Redis数据库连接
	rdb := sql.Client()
	zsetKey := vname
	languages := []redis.Z{}
	num, err := rdb.ZAdd(zsetKey, languages... ).Result()
	if err != nil{
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)
	fmt.Println("创建成功")
	c.JSON(http.StatusOK, gin.H{"status":200,"message":"赛事"+vname+"创办成功"})
}
