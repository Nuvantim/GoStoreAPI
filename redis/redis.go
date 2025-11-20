package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"

	"api/config"
)

var RDS *redis.Client

var ctx = context.Background()

type noopLogger struct{}

func (noopLogger) Printf(_ context.Context, _ string, _ ...interface{}) {}

var once sync.Once

func InitRedis() {
	// Disable log
	redis.SetLogger(noopLogger{})
	once.Do(func() {
		rdsConfig, err := config.GetRedisConfig()
		if err != nil {
			log.Fatalf("Failed to get database config: %v", err)
		}

		// Initialization Redia Connection
		rdb := redis.NewClient(&redis.Options{
			Addr:            fmt.Sprintf("%s:%d", rdsConfig.Host, rdsConfig.Port),
			Password:        rdsConfig.Password,
			DB:              0,
			ReadBufferSize:  1024 * 1024,
			WriteBufferSize: 1024 * 1024,
		})

		// Test Connection
		if err := rdb.Ping(ctx).Err(); err != nil {
			log.Println("Failed connecting to redis:", err)
			return
		}
		fmt.Println("Redis connected...")

		RDS = rdb
	})

}

func RedisClose() {
	if err := RDS.Close(); err != nil {
		log.Println("Failed close Redis : ", err)
	} else {
		log.Println("Success close Redis conncetion...")
	}
}

func SetData[T any](key string, data T) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return RDS.Set(ctx, key, b, 10*time.Minute).Err()
}

func GetData[T any](key string) (*T, error) {
	val, err := RDS.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var data T

	if err := json.Unmarshal(val, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func DelData(key string) error {
	if err := RDS.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil

}
