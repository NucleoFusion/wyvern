package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"wyvern-server/internal/models"
	"wyvern-server/internal/utils"

	"github.com/gin-gonic/gin"
)

func ParseAuth(c *gin.Context) {
	// TODO: Add any paths that do not require this here (if condns)

	// Getting Context
	app, err := utils.GetContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Handling no uuid
	session := c.Query("session")
	if session == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "session data not provided or invalid"})
		c.Abort()
		return
	}

	// Parsing Cookie
	var cookieData models.UserCookie
	if err = json.Unmarshal([]byte(session), &cookieData); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid/dirty session data"})
		c.Abort()
		return
	}

	// Getting token
	token, err := app.Rdb.Get(context.Background(), "sess:"+cookieData.UUID).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.Set("token", token)
	c.Set("session", &cookieData)

	c.Next()
}
