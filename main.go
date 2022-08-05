package main

import (
	"fmt"
	"net/http"

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

func ServerStatus(c *gin.Context) {
	c.JSON(200, "Ecell backend is up and running!")
}

func main() {
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, "Ecell backend is up and running!")
	})

	err := http.ListenAndServe(fmt.Sprintf(":%v", config.Cfg.PORT), engine)
	if err != nil {
		fmt.Printf("Error in starting server at port %v", config.Cfg.PORT)
	}
}
