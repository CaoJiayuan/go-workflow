package model

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis"
)

var redisClusterClient *redis.ClusterClient
var redisClient *redis.Client
var clusterIsOpen = false

// RedisOpen 是否连接 redis
var RedisOpen = false

// SetRedis 设置redis
func SetRedis() {
	fmt.Println("-------启动redis--------")
	ctx := context.Background()
	if conf.RedisCluster == "true" {
		clusterIsOpen = true
		redisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    []string{conf.RedisHost + ":" + conf.RedisPort},
			Password: conf.RedisPassword,
		})
		pong, err := redisClusterClient.Ping(ctx).Result()
		if err != nil {
			fmt.Printf("------------连接 redis cluster：%s 失败,原因：%v\n", conf.RedisHost+":"+conf.RedisPort, err)
		}
		RedisOpen = true
		fmt.Printf("---------连接 redis cluster 成功, %v\n", pong)
	} else {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     conf.RedisHost + ":" + conf.RedisPort,
			Password: conf.RedisPassword,
		})
		pong, err := redisClient.Ping(ctx).Result()
		if err != nil {
			fmt.Printf("------------连接 redis：%s 失败,原因：%v\n", conf.RedisHost+":"+conf.RedisPort, err)
		}
		RedisOpen = true
		fmt.Printf("---------连接 redis  成功, %v\n", pong)
	}
}

// RedisSetVal 将值保存到redis
func RedisSetVal(key, value string, expiration time.Duration) error {
	ctx := context.Background()

	if clusterIsOpen {
		return redisClusterClient.Set(ctx, key, value, expiration).Err()
	}
	return redisClient.Set(ctx, key, value, expiration).Err()
}

// RedisGetVal 从redis获取值
func RedisGetVal(key string) (string, error) {
	ctx := context.Background()

	if clusterIsOpen {
		return redisClusterClient.Get(ctx, key).Result()
	}
	return redisClient.Get(ctx, key).Result()
}
