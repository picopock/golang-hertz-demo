package redis

import (
	"context"
	"fmt"
	"hertz_demo/conf"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func Init() {
	redisConf := conf.Conf.Redis
	address := fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: redisConf.Username,
		Password: redisConf.Password,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
