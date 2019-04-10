package config

import "os"

func IsDevelop() bool {
	return os.Getenv("GO_ENV") != "PROD"
}

func PgConfigString() string {
	if IsDevelop() {
		return "host=localhost port=5432 user=root dbname=cache password=pass sslmode=disable"
	}

	return "host=pg port=5432 user=root dbname=cache password=pass sslmode=disable"
}

func RedisConfigAddr() string {
	if IsDevelop() {
		return "localhost:6379"
	}
	return "redis:6379"
}
