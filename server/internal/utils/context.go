package utils

import (
	"errors"

	"wyvern-server/internal/models"

	"github.com/gin-gonic/gin"
)

func GetContext(c *gin.Context) (*models.AppContext, error) {
	val, ok := c.Get("app")
	if !ok {
		return &models.AppContext{}, errors.New("could not access app context")
	}
	app := val.(*models.AppContext)

	return app, nil
}
