package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"redrock20200309/model"
	"redrock20200309/sql"
	"redrock20200309/token"
	"strconv"
)

//投票
func Vote(c *gin.Context) {
	//token
	tokens := c.GetHeader("token")
	id, username, err := token.CheckToken(tokens)

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(id, username)

	//接口
	vid := c.PostForm("vote_id")
	vidInt, _ := strconv.Atoi(vid) //string->int好像这个没啥用
	uid := c.PostForm("user_id")
	uidInt, _ := strconv.Atoi(uid) //string->int
	fmt.Println(vidInt, uidInt)

	//Redis数据库连接
	rdb := sql.Client()
	vote, err := rdb.ZIncrBy(vid, 1, uid).Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Println(vote)
	num, err := rdb.Get(id).Result()
	num += num
	num, err = rdb.Set(id, 1,0).Result()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	//MySQL数据库连接
	db := sql.Link()
	defer db.Close()
	db.AutoMigrate(&model.Voteuser{})
	var u User
	db.Where("id=?", uid).First(&u)
}

