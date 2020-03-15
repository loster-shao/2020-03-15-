package sql

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

//第二种连接方法
type Cfg struct {
	Addrs    []string
	Pwd      string
	PoolSize int
	DB       int
}


func RedisConnect() *redis.Client {

	c := Cfg{}
	c.Addrs = append(c.Addrs, "localhost:6379")
	c.PoolSize = 10
	c.Pwd = ""
	c.DB = 1

	return redis.NewClient(&redis.Options{
		Addr:     c.Addrs[0],
		Password: c.Pwd,
		PoolSize: c.PoolSize,
		DB:       c.DB,
	})
}

