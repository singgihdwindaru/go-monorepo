package resources

import (
	"context"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var (
	redisInstance *redis.Client
	once          sync.Once
)

// GetRedisClient ensures a single Redis client
func GetRedisClient() (*redis.Client, error) {
	var err error
	once.Do(func() {
		redisInstance = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})
		err = redisInstance.Ping(context.Background()).Err()
	})
	return redisInstance, err
}
