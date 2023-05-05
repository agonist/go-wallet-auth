package main

import (
	"github.com/agonist/goblockapi"
	"github.com/agonist/goblockapi/api"
	"github.com/agonist/goblockapi/api/middleware"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	app := goblockapi.Init()

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
		core.GET("/gasPrice", api.GetGasPrice)
		core.GET("/balance/:address", api.GetBalance)
	}
	auth := r.Group("/auth")
	{
		auth.GET("/nonce/:address", api.Nonce)
		auth.POST("/signin", api.Signin)
	}

	users := r.Group("/users").Use(middleware.Auth())
	{
		users.GET("/me", api.GetUser)
	}

	r.Run()
}
