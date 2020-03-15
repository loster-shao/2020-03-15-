package sql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "strconv"
)

var DB *gorm.DB

func Link() (*gorm.DB){
	DB, error := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if error != nil {
		fmt.Println("err:", error)
		return nil
	}
	return DB
}
//连接数据库（尽量保证在登录注册的接口里可以不用再写一遍，redis也是的）