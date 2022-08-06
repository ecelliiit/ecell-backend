package main

import (
	"fmt"

	"github.com/ecelliiit/ecell-backend/config"
	"github.com/ecelliiit/ecell-backend/db"
	"github.com/ecelliiit/ecell-backend/routes"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	config.Load()
	db.Connect()
	engine = routes.BuildRoutes()
}

func main() {
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, "Ecell backend is up and running!")
	})

	err := engine.Run(":" + config.Cfg.PORT)
	if err != nil {
		fmt.Printf("Error in starting server at port %v", config.Cfg.PORT)
	}
}
