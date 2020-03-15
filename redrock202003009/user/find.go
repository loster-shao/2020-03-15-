package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock20200309/sql"
)

//查询排行榜
func Find(c *gin.Context)  {
	//接口
	Vote_id   := c.PostForm("vote_id")
	//Redis数据库连接
	rdb := sql.Client()
	ret, err := rdb.ZRevRangeWithScores(Vote_id, 0, 9).Result()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(ret)
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
	c.JSON(200, gin.H{"status": http.StatusOK, "message": ret})
}
