package request

import (
	"fmt"
	"github.com/go-redis/redis"
)

func TestConnect(ch chan string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Omm*D6?v6gU9", // Redis无密码
		DB:       1,  // 使用默认数据库
	})

	// 测试连接
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	res, err := client.HGet("Campaign","317524").Result()
	//fmt.Println(res, err)
	ch <- res
}