package database

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

func init() {
	RedisManager := NewRedisTool()
	if _, err := RedisManager.Ping(ctx).Result(); err != nil {
		panic("redis 连接出错，" + err.Error())
	}
}

type RedisTool struct {
	*redis.Client
}

var redisTool *RedisTool

var redisOnce sync.Once

func GetRedisHelper() *RedisTool {
	return redisTool
}

func NewRedisTool() *redis.Client {

	var rdb *redis.Client

	redisOnce.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:         "localhost:6379",
			Password:     "123456",
			DB:           1,
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			PoolSize:     10,
			PoolTimeout:  30 * time.Second,
		})

		fmt.Println("redisOnce")
		rdh := new(RedisTool)
		rdh.Client = rdb
		redisTool = rdh
	})

	//
	// fmt.Println("NewRedisTool")

	return rdb
}
