package main

import (
	"wyvern-server/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.AddRoutes(r)

	r.Run(":3000")
}
