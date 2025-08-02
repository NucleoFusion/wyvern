package dbctx

import (
	"database/sql"
	"log"

	"wyvern-server/internal/db/mongodb"
	"wyvern-server/internal/db/pg"
	redisdb "wyvern-server/internal/db/redis"
	"wyvern-server/internal/models"

	"github.com/redis/go-redis/v9"
)

type DBContext struct {
	Pg    *sql.DB
	Rdb   *redis.Client
	Mongo *models.Mongo
}

var ctx *DBContext

func SetCtx() {
	pg, err := pg.ConnectPG()
	if err != nil {
		log.Fatal("Postgres Error:", err)
	}

	rdb := redisdb.ConnectRedis()

	mongo, err := mongodb.ConnectMongo()
	if err != nil {
		log.Fatal("Postgres Error:", err)
	}

	ctx = &DBContext{
		Pg:    pg,
		Rdb:   rdb,
		Mongo: mongo,
	}
}

func GetCtx() *DBContext {
	return ctx
}
