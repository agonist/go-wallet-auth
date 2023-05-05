package api

import (
	"net/http"

	"github.com/agonist/goblockapi"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	app := c.MustGet("app").(*goblockapi.App)
	address := c.MustGet("address")

	var user goblockapi.User
	app.Db.Where("address = ?", address).First(&user)
	c.JSON(http.StatusOK, user)
}
