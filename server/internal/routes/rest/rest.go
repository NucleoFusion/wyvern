package rest

import "github.com/gin-gonic/gin"

func AddRestRoutes(r *gin.Engine) {
	grp := r.Group("/rest")

	grp.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "PONG"})
	})
}
