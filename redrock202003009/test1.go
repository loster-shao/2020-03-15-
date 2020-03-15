package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func InitClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println("successful connecting!")
	return nil
}


func redisExample() {
	err := rdb.Set("score", 10, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)
}

func redisExample2() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 99.0, Member: "Golang"},
		}
	language := []redis.Z{
		redis.Z{Score: 99.0, Member: "Go"},
	}
	// ZADD
	num, err := rdb.ZAdd(zsetKey, languages... ).Result()
	num1, err := rdb.ZAdd(zsetKey, language... ).Result()
	if err != nil {
		fmt.
			Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n\n", num, num1)
	//查找特定值
	uid := "language_rank"
	vote, err := rdb.Get(uid).Result()
	fmt.Println("s", vote)
	//Del
	//num, err = rdb.Del(zsetKey).Result()
	//rdb.Do("DEL", "Go")

	//// 把Golang的分数加10
	//newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	//if err != nil {
	//	fmt.Printf("zincrby failed, err:%v\n", err)
	//	return
	//}
	//fmt.Printf("Golang's score is %f now.\n", newScore)
	//
	//// 取分数最高的3个
	//ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	//if err != nil {
	//	fmt.Printf("zrevrange failed, err:%v\n", err)
	//	return
	//}
	//for _, z := range ret {
	//	fmt.Println(z.Member, z.Score)
	//}
	//// 取95~100分的
	//op := &redis.ZRangeBy{
	//	Min: "95",
	//	Max: "100",
	//}
	//ret, err = rdb.ZRangeByScoreWithScores(zsetKey, *op ).Result()
	//if err != nil {
	//	fmt.Printf("zrangebyscore failed, err:%v\n", err)
	//	return
	//}
	//for _, z := range ret {
	//	fmt.Println(z.Member, z.Score)
	//}
}

func main(){
	_ = InitClient()
	redisExample()
	redisExample2()
	//num := 1
	//num += num
	//fmt.Println(num)
}

