package ws

import (
	"net/http"
	"strconv"

	"wyvern-server/internal/log"
	"wyvern-server/internal/managers"
	"wyvern-server/internal/managers/hub"
	"wyvern-server/internal/models"
	"wyvern-server/internal/utils"

	"github.com/gin-gonic/gin"
)

func Connect(c *gin.Context) {
	// Getting Contexts
	app, err := utils.GetContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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

	app.Managers.HubChan <- &hub.HubMessage{
		MsgType: hub.RegisterClient,
		Param:   &client,
	}

	go managers.HandleClientReads(&client)
}
