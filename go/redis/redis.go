package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client     *redis.Client
	ReqTimeout time.Duration
}

// interfaces for redis client
var _ CacheClient = (*RedisClient)(nil)

func NewRedisClient(clientOptions *redis.Options, requestTimeout time.Duration) (*RedisClient, error) {
	c := redis.NewClient(clientOptions)

	val, err := c.Ping(ctx).Result()
	if err != nil {
		log.Println("Error pinging redis")
		return nil, err
	}

	log.Println(val)
	return &RedisClient{
		Client:     c,
		ReqTimeout: requestTimeout,
	}, nil
}

func (rc *RedisClient) Set(key string, data *UserToken) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, rc.ReqTimeout)
	defer cancel()
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshalling data")
		return err
	}

	return rc.Client.Set(ctxWithTimeout, key, string(bytes), 0).Err()
}

func (rc *RedisClient) Get(key string) (*UserToken, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, rc.ReqTimeout)
	defer cancel()

	val, err := rc.Client.Get(ctxWithTimeout, key).Result()
	if err != nil {
		log.Println("Error getting data")
		return nil, err
	}

	var data UserToken
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		log.Println("Error unmarshalling data")
		return nil, err
	}

	return &data, nil
}
func (rc *RedisClient) Del(keys ...string) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, rc.ReqTimeout)
	defer cancel()

	return rc.Client.Del(ctxWithTimeout, keys...).Err()
}

func (rc *RedisClient) DelOne(key string) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, rc.ReqTimeout)
	defer cancel()

	return rc.Client.Del(ctxWithTimeout, key).Err()
}
