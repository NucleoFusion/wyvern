package main

import (
	"log"

	"wyvern-server/internal/db/pg"
	redisdb "wyvern-server/internal/db/redis"
	"wyvern-server/internal/middleware"
	"wyvern-server/internal/models"
	"wyvern-server/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	pg, err := pg.ConnectPG()
	if err != nil {
		log.Fatal("Postgres Error:", err)
	}
	defer pg.Close()

	rdb := redisdb.ConnectRedis()
	defer rdb.Close()

	ctx := &models.AppContext{
		Pg:  pg,
		Rdb: rdb,
	}

	r := gin.Default()

	r.Use(middleware.Inject(ctx)) // Injecting AppContext

	routes.AddRoutes(r)

	r.Run(":3000")
}
