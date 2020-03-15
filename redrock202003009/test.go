//package redrock20200309
//
//import (
//
//	"fmt"
//
//	"time"
//
//	"github.com/jinzhu/gorm"
//
//	_ "github.com/jinzhu/gorm/dialects/mysql"
//
//)
//
//type Like struct {
//
//	ID  int `gorm:"primary_key"`
//
//	Ip  string `gorm:"type:varchar(20);not null;"`
//
//	Ua  string `gorm:"type:varchar(256);not null;"`
//
//	Title string `gorm:"type:varchar(128);not null;"`
//
//	Hash  string `gorm:"unique_index:hash_idx"`
//
//	CreateAt time.Time
//
//}
//
//var db_1 *gorm.DB
//
//func init() {
//
//	var err error
//
//	db_1, err = gorm.Open("mysql", "root:12345678@/mydatabase?charset=utf8")
//
//	if err != nil {
//
//		panic(err)
//
//	}
//
//}
//
//func main() {
//
//	//创建表
//
//	if err := db_1.CreateTable(&Like{}).Error; err != nil {
//
//		panic(err)
//
//	}
//
//	//插入
//
//	like := &Like{
//
//		ID:  2,
//
//		Ip:  "HHH",
//
//		Ua:  "JJJ",
//
//		Title: "Title",
//
//		Hash:  "999",
//
//		CreateAt: time.Now(),
//
//	}
//
//	if err := db_1.Create(like).Error; err != nil {
//
//		fmt.Println(err)
//
//	}
//
//	//删除
//
//	if err := db_1.Where(&Like{ID: 1}).Delete(Like{}).Error; err != nil {
//
//		fmt.Println(err)
//
//	}
//
//	//查询1
//
//	Query1()
//
//	//查询2
//
//	Query2()
//
//	//遍历
//
//	QueryAll()
//
//	//更新
//
//	Update()
//
//}
//
////查询1
//
//func Query1() (bool, error) {
//
//	var count int //查到的数量
//
//	err := db_1.Model(&Like{}).Where(&Like{ID: 1}).Count(&count).Error
//
//	fmt.Println(count)
//
//	if err != nil {
//
//		return false, err
//
//	}
//
//	return false, err
//
//}
//
////查询ID为1的Like
//
//func Query2() {
//
//	var Likes Like
//
//	db_1.Model(&Like{}).Where(&Like{ID: 1}).Find(&Likes)
//
//	fmt.Println(Likes)
//
//}
//
////查询全部数据
//
//func QueryAll() {
//
//	var tables []Like
//
//	err := db_1.Where(&Like{}).Find(&tables)
//
//	if err != nil {
//
//		fmt.Println(err)
//
//	}
//
//	fmt.Println(tables)
//
//}
//
////更新数据
//
//func Update() {
//
//	db_1.Model(&Like{}).Where(Like{ID: 2}).Updates(Like{Ip: "Hello"}) //Model
//
//}
package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func Client() (*redis.Client) {
	var err error
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = Rdb.Ping().Result()
	if err != nil {
		return nil
	}
	fmt.Println("successful connecting!")
	return Rdb
}

func main()  {
	Client()
	var n string
	fmt.Scanf(n)
	Test(n)
}

func Test(n string) {
	zsetKey := n

	rdb.ZAdd(zsetKey, ).Err()
	//rdb.Set("s2",10,0).Err()
	//rdb.Set("s3",10,0).Err()
	//rdb.Set("s4",10,0).Err()
    //ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
}