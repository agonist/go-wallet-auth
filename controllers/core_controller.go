package controllers

import (
	"go-block-api/evm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	client := c.MustGet("ethClient").(*evm.Client)
	address := c.Param("address")

	balance, err := client.GetBalance(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, balance)
}

func GetGasPrice(c *gin.Context) {
	client := c.MustGet("ethClient").(*evm.Client)

	gasPrice, err := client.GetGasPrice()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gasPrice)
}
