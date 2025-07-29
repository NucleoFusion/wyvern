package main

import (
	"wyvern-server/internal/managers/channel"
	"wyvern-server/internal/managers/hub"
	"wyvern-server/internal/middleware"
	"wyvern-server/internal/models"
	"wyvern-server/internal/routes"
	"wyvern-server/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := utils.CreateContext()

	ctx.Managers.HubChan <- &hub.HubMessage{
		MsgType: hub.AddHub,
		Param:   models.Mongo{},
	}

	ctx.Managers.ChanChan <- &channel.ChannelMessage{
		MsgType: channel.AddChannel,
		Param:   models.Mongo{},
	}

	r := gin.Default()

	r.Use(middleware.Inject(ctx)) // Injecting AppContext
	routes.AddRoutes(r)

	r.Run(":3000")
}
