package models

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type AppContext struct {
	Pg  *sql.DB
	Rdb *redis.Client
}
