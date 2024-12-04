package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string
	DB       int
	Username string
	Password string
}

type cli struct {
	client *redis.Client
}

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error)
	Get(ctx context.Context, key string, value interface{}) (err error)
	RedisClient() *redis.Client
}

func New(conf Config) Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		DB:       conf.DB,
		Username: conf.Username,
		Password: conf.Password,
	})

	return &cli{
		client: client,
	}
}

func (r *cli) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error) {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, p, expiration).Err()
}

func (r *cli) Get(ctx context.Context, key string, value interface{}) (err error) {
	var p []byte
	if err = r.client.Get(ctx, key).Scan(&p); err != nil {
		return err
	}
	return json.Unmarshal(p, &value)
}

func (r *cli) RedisClient() *redis.Client {
	return r.client
}
