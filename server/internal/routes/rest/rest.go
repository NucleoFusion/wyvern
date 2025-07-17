package rest

import (
	"wyvern-server/internal/middleware"
	"wyvern-server/internal/routes/rest/channels"
	"wyvern-server/internal/routes/rest/hub"

	"github.com/gin-gonic/gin"
)

func AddRestRoutes(r *gin.Engine) {
	grp := r.Group("/rest")
	grp.Use(middleware.ParseAuth)

	hub.AddHubRoutes(grp)
	channels.AddChannelsRoutes(grp)

	grp.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "PONG"})
	})
}
