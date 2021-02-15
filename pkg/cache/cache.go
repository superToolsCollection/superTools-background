package cache

import (
	"github.com/go-redis/redis/v8"

	"superTools-background/pkg/setting"

	"time"
)

/**
* @Author: super
* @Date: 2020-11-18 11:35
* @Description: 根据配置创建redis连接池
**/

func NewRedisEngine(cacheSetting *setting.CacheSettingS) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:cacheSetting.Host,
		Password:cacheSetting.Password,
		MaxRetries:5,
		IdleTimeout:300 * time.Second,
	})
	return client, nil
}
