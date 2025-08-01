package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func AddWsRoutes(r *gin.Engine) {
	grp := r.Group("/ws")

	grp.GET("/connect/:hubID", Connect)
}
