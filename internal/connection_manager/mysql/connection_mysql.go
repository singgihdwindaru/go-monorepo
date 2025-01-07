package resources

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlInstance *sql.DB
	once          sync.Once
)

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
