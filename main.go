package main

import (
	"github.com/gin-gonic/gin"
	"go-block-api/controllers"
	"go-block-api/evm"
)

func main() {
	r := gin.Default()

	client := evm.New("https://eth-mainnet.g.alchemy.com/v2/pt26dRBVnOVRXaLmzCB7oIZ0o1PfIAcu")

	r.Use(func(c *gin.Context) {
		c.Set("ethClient", client)
	})

	core := r.Group("/core")
	{
		core.GET("/gasPrice", controllers.GetGasPrice)
		core.GET("/balance/:address", controllers.GetBalance)
	}
	r.Run()
}
