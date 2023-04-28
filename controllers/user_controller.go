package controllers

import (
	"go-block-api/config"
	"go-block-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	app := c.MustGet("app").(*config.App)
	address := c.MustGet("address")

	var user model.User
	app.Db.Where("address = ?", address).First(&user)
	c.JSON(http.StatusOK, user)
}
