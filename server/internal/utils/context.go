package utils

import (
	"errors"
	"log"

	"wyvern-server/internal/db/mongodb"
	"wyvern-server/internal/db/pg"
	redisdb "wyvern-server/internal/db/redis"
	"wyvern-server/internal/models"

	"github.com/gin-gonic/gin"
)

func GetContext(c *gin.Context) (*models.AppContext, error) {
	val, ok := c.Get("app")
	if !ok {
		return &models.AppContext{}, errors.New("could not access app context")
	}
	app := val.(*models.AppContext)

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

func CreateContext() *models.AppContext {
	pg, err := pg.ConnectPG()
	if err != nil {
		log.Fatal("Postgres Error:", err)
	}
	defer pg.Close()

	rdb := redisdb.ConnectRedis()
	defer rdb.Close()

	mongo, err := mongodb.ConnectMongo()
	if err != nil {
		log.Fatal("Postgres Error:", err)
	}

	return &models.AppContext{
		Pg:    pg,
		Rdb:   rdb,
		Mongo: mongo,
	}
}
