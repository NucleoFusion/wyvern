package channels

import "github.com/gin-gonic/gin"

func AddChannelsRoutes(r *gin.RouterGroup) {
	grp := r.Group("/channels")

	grp.POST("/create", CreateChannel)

	grp.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "PONG"})
	})
}
