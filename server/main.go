package main

import (
	"wyvern-server/internal/middleware"
	"wyvern-server/internal/routes"
	"wyvern-server/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := utils.CreateContext()

	r := gin.Default()

	r.Use(middleware.Inject(ctx)) // Injecting AppContext
	routes.AddRoutes(r)

	r.Run(":3000")
}
