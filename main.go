package main

import (
	"go-block-api/config"
	"go-block-api/controllers"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	app := config.Init()

	r := gin.Default()
	r.Use(cors.Default())

	r.Use(func(c *gin.Context) {
		c.Set("app", app)
	})

	core := r.Group("/core")
	{
		core.GET("/gasPrice", controllers.GetGasPrice)
		core.GET("/balance/:address", controllers.GetBalance)
	}
	auth := r.Group("/auth")
	{
		auth.GET("/nonce/:address", controllers.Nonce)
		auth.POST("/signin", controllers.Signin)
	}

	r.Run()
}
