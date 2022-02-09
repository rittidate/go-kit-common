package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	ConnectionString string `mapstructure:"connection_string"`
	Client           RedisClient
}

type RedisClient struct {
	redisClient *redis.Client
	ctx         context.Context
}

func (r *Redis) Binding() error {
	fmt.Println(r.ConnectionString)

	client := redis.NewClient(&redis.Options{
		Addr:     r.ConnectionString,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	client.Ping(context.Background()).Err()
	if err := client.Ping(context.Background()).Err(); err != nil {
		return err
	}
	r.Client = RedisClient{redisClient: client, ctx: context.Background()}
	return nil
}

func (rc RedisClient) saveRedisData(key string, data map[string]interface{}, t time.Duration) error {
	jd, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal data: %v", err)
	}

	return rc.redisClient.Set(rc.ctx, key, jd, t).Err()
}

// SaveDataOnRedis is a function to save data on redis
func (rc RedisClient) SaveDataOnRedis(key string, data map[string]interface{}) error {
	return rc.saveRedisData(key, data, 0)
}

// SaveDataOnRedisTTL is a function to save data on redis ttl
func (rc RedisClient) SaveDataOnRedisTTL(key string, data map[string]interface{}, t time.Duration) error {
	return rc.saveRedisData(key, data, t)
}

// SaveByteDataOnRedisTTL is a function to save data on redis ttl
func (rc RedisClient) SaveByteDataOnRedisTTL(key string, data []byte, t time.Duration) error {
	return rc.redisClient.Set(rc.ctx, key, data, t).Err()
}

// GetByteDataFromRedis is a function to get data from redis
func (rc RedisClient) GetByteDataFromRedis(key string) ([]byte, error) {
	val, err := rc.redisClient.Get(rc.ctx, key).Bytes()
	if err != nil || err == redis.Nil {
		return nil, err
	}
	return val, nil
}

// GetDataFromRedis is a function to get data from redis
func (rc RedisClient) GetDataFromRedis(key string) (map[string]interface{}, error) {
	val, err := rc.redisClient.Get(rc.ctx, key).Bytes()
	if err != nil || err == redis.Nil {
		return nil, err
	}

	rtn := make(map[string]interface{})
	if err := json.Unmarshal(val, &rtn); err != nil {
		return nil, fmt.Errorf("unable to unmarshal data: %v", err)
	}
	return rtn, nil
}

// PublishToRedisChannel is a function to publish data to redis channel
func (rc RedisClient) PublishToRedisChannel(ch string, message []byte) error {
	_, err := rc.redisClient.Publish(rc.ctx, ch, message).Result()
	if err != nil {
		return fmt.Errorf("unable to publish data: %v", err)
	}

	return nil
}

// RemoveDataOnRedis is a function to remove date on redis
func (rc RedisClient) RemoveDataOnRedis(pattern string) error {
	ssc := rc.redisClient.Keys(rc.ctx, pattern)
	if ssc == nil || len(ssc.Val()) == 0 {
		return redis.Nil
	}
	keys := ssc.Val()
	for _, v := range keys {
		if _, err := rc.redisClient.Del(rc.ctx, v).Result(); err != nil {
			return err
		}
	}
	return nil
}
