package tools

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"hmdp/config"
	"hmdp/global"
	"time"
)

func InitializeRedis() *redis.Client {
	rdConf := config.Conf.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     rdConf.Host + ":" + rdConf.Port,
		Password: rdConf.Password,
		DB:       rdConf.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logrus.Debugf("Redis connect ping failed, err: %s", err)
		return nil
	}
	return client
}

const (
	COUNT_BITS      = 32
	BEGIN_TIMESTAMP = 1640995200
)

func NextId(keyPrefix string) uint64 {
	client := global.App.Redis
	// 生成时间戳
	now := time.Now()
	nowSecond := now.Unix()
	timestamp := nowSecond - BEGIN_TIMESTAMP

	// 生成序列号
	// 获取日期
	date := now.Format("2006:01:02")
	// 自增
	ctx := context.Background()
	count, _ := client.Incr(ctx, "icr:"+keyPrefix+":"+date).Uint64()

	return (uint64(timestamp) << COUNT_BITS) | count
}
