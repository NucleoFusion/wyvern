package middleware

import (
	"wyvern-server/internal/models"

	"github.com/gin-gonic/gin"
)

func Inject(appCtx *models.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app", appCtx)
		c.Next()
	}
}
