package redisservice

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"site_board_monitor/config"
)

func InitRedis(c config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Pwd,
		DB:       c.Db,
	})
}
