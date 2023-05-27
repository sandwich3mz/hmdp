package tools

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"hmdp/config"
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
