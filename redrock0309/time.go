package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	voteInfo.Starttime = strings.ReplaceAll(voteInfo.Starttime, "T",      "")
	voteInfo.Starttime = strings.ReplaceAll(voteInfo.Starttime, "+08:00", "")
	voteInfo.Endtime   = strings.ReplaceAll(voteInfo.Endtime,   "T",      "")
	voteInfo.Endtime   = strings.ReplaceAll(voteInfo.Endtime,   "+08:00", "")
	nowTime := time.Now() //获取当前时间
	sTime,_ := time.Parse("2006-01-02 15:04:05", voteInfo.Starttime)
	eTime,_ := time.Parse("2006-01-02 15:04:05", voteInfo.Endtime)
	fmt.Println(voteInfo)
}
