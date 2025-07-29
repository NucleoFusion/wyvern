package middleware

import (
	"wyvern-server/internal/utils"

	"github.com/gin-gonic/gin"
)

func Inject(appCtx *utils.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app", appCtx)
		c.Next()
	}
}
