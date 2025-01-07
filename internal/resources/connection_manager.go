package resources

import (
	"context"
	"database/sql"
	"sync"

	"cloud.google.com/go/firestore"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var (
	mysqlInstance     *sql.DB
	redisInstance     *redis.Client
	firestoreInstance *firestore.Client
	once              sync.Once
)

// GetMySQLConnection ensures a single MySQL connection pool
func GetMySQLConnection() (*sql.DB, error) {
	var err error
	once.Do(func() {
		mysqlInstance, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
		if err == nil {
			mysqlInstance.SetMaxOpenConns(10) // Example: limit max connections
			mysqlInstance.SetMaxIdleConns(5)
		}
	})
	return mysqlInstance, err
}

// GetRedisClient ensures a single Redis client
func GetRedisClient() *redis.Client {
	once.Do(func() {
		redisInstance = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})
	})
	return redisInstance
}

// GetFirestoreClient ensures a single Firestore client
func GetFirestoreClient(ctx context.Context) *firestore.Client {
	once.Do(func() {
		firestoreInstance, _ = firestore.NewClient(ctx, "project-id")
	})
	return firestoreInstance
}
