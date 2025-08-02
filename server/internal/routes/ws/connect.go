package ws

import (
	"net/http"
	"strconv"

	"wyvern-server/internal/dbctx/mgrctx"
	"wyvern-server/internal/handlers"
	"wyvern-server/internal/log"
	"wyvern-server/internal/managers/hub"
	"wyvern-server/internal/models"

	"github.com/gin-gonic/gin"
)

func Connect(c *gin.Context) {
	hubIDStr := c.Param("hubID")
	userIDStr := c.Query("userID")
	if hubIDStr == "" || userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid/missing parameters"})
	}

	hubID, err := strconv.Atoi(hubIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameters"})
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameters"})
	}

	conn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Log(log.Fatal, err.Error())
	}

	client := models.Client{
		Conn:   conn,
		HubID:  int(hubID),
		UserID: int(userID),
	}

	mgrctx.GetCtx().HubChan <- &hub.HubMessage{
		MsgType: hub.RegisterClient,
		Param:   &client,
	}

	go handlers.HandleClientReads(&client)
}
