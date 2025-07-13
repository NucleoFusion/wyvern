package routes

import (
	"wyvern-server/internal/routes/rest"
	"wyvern-server/internal/routes/ws"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {
	rest.AddRestRoutes(r)
	ws.AddWsRoutes(r)
}
