package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
	"wm-take-out/global"
)

var client *redis.Client

func InitRedis() *redis.Client {
	redisOpt := global.Config.Redis
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisOpt.Host, redisOpt.Port),
		Password: redisOpt.Password,
		DB:       redisOpt.DataBase,
	})
	ping := client.Ping()
	err := ping.Err()
	if err != nil {
		panic(err)
	}
	return client
}
