package models

import (
	"database/sql"

	"wyvern-server/internal/managers"

	"github.com/redis/go-redis/v9"
)

type AppContext struct {
	Pg       *sql.DB
	Rdb      *redis.Client
	Mongo    *Mongo
	Managers *managers.Managers
}
