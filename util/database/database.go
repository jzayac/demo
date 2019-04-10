package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"demo/util/config"
	"github.com/go-redis/redis"
)

func PgNewClient() (*gorm.DB, error) {
	dbCfg := config.PgConfigString()

	con, err := gorm.Open("postgres", dbCfg)
	if err != nil {
		return nil, ErrSqlConnectionProblem
	}
	return con, err
}

func RedisNewClient() (*redis.Client, error) {
	addr := config.RedisConfigAddr()
	client := redis.NewClient(&redis.Options{
		// Addr:     "localhost:6379",
		Addr:     addr,
		Password: "",
		DB:       0, // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, ErrRedisConnectionProblem
	}

	return client, nil
}

var ErrSqlConnectionProblem = errors.New("Connect| sql connection problem")
var ErrSqlInitDbDriver = errors.New("Initialize| Postgres Drive Error")

var ErrRedisConnectionProblem = errors.New("Connect| redis connection problem")
