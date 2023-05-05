package api

import (
	"net/http"

	"github.com/agonist/goblockapi"
	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	app := c.MustGet("app").(*goblockapi.App)
	address := c.Param("address")

	balance, err := app.Rpc.GetBalance(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, balance)
}

func GetGasPrice(c *gin.Context) {
	app := c.MustGet("app").(*goblockapi.App)

	gasPrice, err := app.Rpc.GetGasPrice()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gasPrice)
}
