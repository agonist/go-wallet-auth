package main

import (
	"github.com/gin-gonic/gin"
	"go-block-api/config"
	"go-block-api/controllers"
	"go-block-api/middlewares"

	"github.com/gin-contrib/cors"
)

func main() {
	app := config.Init()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET, POST, OPTIONS, PUT, DELETE"},
		AllowHeaders: []string{"*"},
	}))

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

	users := r.Group("/users").Use(middlewares.Auth())
	{
		users.GET("/me", controllers.GetUser)
	}

	r.Run()
}
