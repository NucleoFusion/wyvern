package utils

import (
	"database/sql"
	"errors"
	"log"

	"wyvern-server/internal/db/mongodb"
	"wyvern-server/internal/db/pg"
	redisdb "wyvern-server/internal/db/redis"
	"wyvern-server/internal/managers"
	"wyvern-server/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type AppContext struct {
	Pg       *sql.DB
	Rdb      *redis.Client
	Mongo    *models.Mongo
	Managers *managers.Managers
}

func GetContext(c *gin.Context) (*AppContext, error) {
	val, ok := c.Get("app")
	if !ok {
		return &AppContext{}, errors.New("could not access app context")
	}
	app := val.(*AppContext)

	return app, nil
}

func GetSessionContext(c *gin.Context) (*models.UserCookie, string, error) {
	val, ok := c.Get("session")
	if !ok {
		return &models.UserCookie{}, "", errors.New("could not access session context")
	}

	val2, ok := c.Get("token")
	if !ok {
		return &models.UserCookie{}, "", errors.New("could not access token context")
	}

	session := val.(*models.UserCookie)
	token := val2.(string)

	return session, token, nil
}

func CreateContext() *AppContext {
	pg, err := pg.ConnectPG()
	if err != nil {
		log.Fatal("Postgres Error:", err)
	}

	rdb := redisdb.ConnectRedis()

	mongo, err := mongodb.ConnectMongo()
	if err != nil {
		log.Fatal("Postgres Error:", err)
	}

	return &AppContext{
		Pg:       pg,
		Rdb:      rdb,
		Mongo:    mongo,
		Managers: managers.GetManagers(pg),
	}
}
