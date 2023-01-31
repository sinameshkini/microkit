package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var c *CacheConn

type CacheConn struct {
	Active      bool
	RedisClient *redis.Client
	ExpireTime  time.Duration
}

func GetConnection() *CacheConn {
	return c
}

func Connect(cacheConn *redis.Options, expDur time.Duration) (err error) {
	c = &CacheConn{
		Active:      viper.GetBool("cache.active"),
		RedisClient: redis.NewClient(cacheConn),
		ExpireTime:  expDur,
	}

	return nil
}

func Get(ctx context.Context, key string, val interface{}) (err error) {
	if err = c.RedisClient.Get(ctx, key).Scan(val); err != nil {
		return
	}

	return
}

func Set(ctx context.Context, key string, val interface{}) (err error) {
	if err = c.RedisClient.Set(ctx, key, val, c.ExpireTime).Err(); err != nil {
		return
	}

	return
}

func Active() bool {
	return c != nil && c.Active
}
