package channels

import (
	"database/sql"
	"net/http"
	"strconv"

	"wyvern-server/internal/models"
	"wyvern-server/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateChannel(c *gin.Context) {
	// Getting Contexts
	app, err := utils.GetContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sess, _, err := utils.GetSessionContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Getting Params
	hubID := c.Query("hubID")
	name := c.Query("name")
	privateStr := c.Query("is_private")
	isPrivate, err := strconv.ParseBool(privateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid/missing parameters"})
		return
	}

	if hubID == "" || name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid/missing parameters"})
		return
	}

	// Checking if HUB with id exists
	var ownerID int64
	var repo string
	err = app.Pg.QueryRow(`SELECT owner_id, repo FROM hub WHERE id = $1`, hubID).Scan(&ownerID, &repo)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "hub not found with given id"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Checking if owner and request maker are same
	if sess.GithubID != int(ownerID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "you do not have permission to create channels"})
		return
	}

	var chanRow models.Channel
	err = app.Pg.QueryRow(`INSERT INTO channel(hub_id, name, private) VALUES ($1, $2, $3) RETURNING id, hub_id, name, private`, hubID, name, isPrivate).
		Scan(&chanRow.ID, &chanRow.HubID, &chanRow.Name, &chanRow.Private)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, chanRow)
}
