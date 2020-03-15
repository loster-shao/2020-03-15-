package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"redrock0309/sql"
	"redrock0309/models"
	"redrock0309/token"
	"strconv"
	"fmt"
	"encoding/json"
	"time"
	"strings"
)

//用户给参赛选手投票
func Tp(c *gin.Context)  {
	
	tokenValue := c.GetHeader("token")
	vidStr := c.PostForm("vid")
	vidInt,_ := strconv.Atoi(vidStr)

	xsUserIdStr := c.PostForm("xsuserid")
	xsUserIdInt, _ := strconv.Atoi(xsUserIdStr)

	uid, username, _ := token.CheckToken(tokenValue)

	redisClient := sql.RedisConnect()//连接数据库
	voteInfoStr,_ := redisClient.Get(vidStr).Result()
	var voteInfo models.Vote

	if len(voteInfoStr)>0 {//redis中有值
		
		err := json.Unmarshal([]byte(voteInfoStr), &voteInfo)//

		if err!=nil{
			fmt.Println("反序列化失败3", err)
			return
		}
	}else{

		db := sql.DbConn()//连接数据库
		defer db.Close()
		db.Where("id=?", vidInt).First(&voteInfo)

		dataByte, err := json.Marshal(voteInfo)
		if err!=nil{
			fmt.Println("序列化失败",err)
			return
		}

		dataStr := string(dataByte)
		redisClient.Set(vidStr, dataStr, 60*time.Second).Err()//redis 缓存过期时间60s
	}

	voteInfo.Starttime = strings.ReplaceAll(voteInfo.Starttime, "T",      "")
	voteInfo.Starttime = strings.ReplaceAll(voteInfo.Starttime, "+08:00", "")
	voteInfo.Endtime   = strings.ReplaceAll(voteInfo.Endtime,   "T",      "")
	voteInfo.Endtime   = strings.ReplaceAll(voteInfo.Endtime,   "+08:00", "")
	nowTime := time.Now() //获取当前时间
	sTime,_ := time.Parse("2006-01-02 15:04:05", voteInfo.Starttime)
	eTime,_ := time.Parse("2006-01-02 15:04:05", voteInfo.Endtime)
	fmt.Println(voteInfo)

	if nowTime.Unix() < sTime.Unix(){
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "投票时间未开始"})
		return
	}

	if nowTime.Unix() > eTime.Unix(){
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "投票已结束"})
		return
	}

	nowTimeStr := nowTime.Format("2006-01-02")
	voteCountKey := vidStr + strconv.Itoa(uid) + nowTimeStr
	result, _ := redisClient.Incr(voteCountKey).Result()

	if result > 3{
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "今日3次投票机会已用完"})
		return
	}

	go saveResult(vidInt, xsUserIdInt, uid, username)
	c.JSON(200, gin.H{"status": http.StatusOK, "message": "投票成功"})
}

//保存
func saveResult(voteId int,xsuserid int ,uid int,tpusername string){

	db := sql.DbConn()
	defer db.Close()

	var voteUser models.Voteuser
	db.Where("voteid = ? AND xsuserid >= ?", voteId, xsuserid).First(&voteUser)
	voteUser.Votetotalcount=voteUser.Votetotalcount+1

	db.Model(voteUser).Updates(models.Voteuser{Votetotalcount:voteUser.Votetotalcount})

	var u models.User
	db.Where("id=?", xsuserid).First(&u)

	db.AutoMigrate(&models.Voteresult{})
	db.Create(&models.Voteresult{ Voteid: voteId, Xsuserid: xsuserid, Tpuserid: uid, Xsusername: u.Username, Tpusername: tpusername})
}

