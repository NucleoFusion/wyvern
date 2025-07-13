package ws

import "github.com/gin-gonic/gin"

func AddWsRoutes(r *gin.Engine) {
	grp := r.Group("/ws")

	grp.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "PONG"})
	})
}
