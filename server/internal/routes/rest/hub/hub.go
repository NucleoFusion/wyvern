package hub

import "github.com/gin-gonic/gin"

func AddHubRoutes(r *gin.RouterGroup) {
	grp := r.Group("/hub")

	grp.POST("/create", CreateHub)

	grp.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "PONG"})
	})
}
