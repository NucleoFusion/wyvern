package utils

import (
	"errors"

	"wyvern-server/internal/models"

	"github.com/gin-gonic/gin"
)

func GetContext(c *gin.Context) (*AppContext, error) {
	val, ok := c.Get("app")
	if !ok {
		return &AppContext{}, errors.New("could not access app context")
	}
	app := val.(*AppContext)

	return app, nil
}

func GetSessionContext(c *gin.Context) (*models.UserCookie, string, error) {
	val, ok := c.Get("session")
	if !ok {
		return &models.UserCookie{}, "", errors.New("could not access session context")
	}

	val2, ok := c.Get("token")
	if !ok {
		return &models.UserCookie{}, "", errors.New("could not access token context")
	}

	session := val.(*models.UserCookie)
	token := val2.(string)

	return session, token, nil
}
