package routes

import (
	"github.com/ecelliiit/ecell-backend/controllers"
	"github.com/gin-gonic/gin"
)

func InitSubscriberRoutes(engine *gin.Engine) {
	subscriberController := controllers.NewSubscriberController()
	engineGrouped := engine.Group("/subscriber")

	engineGrouped.GET("", subscriberController.GetAll)
	engineGrouped.GET("/:id", subscriberController.GetById)
	engineGrouped.POST("", subscriberController.Create)
	engineGrouped.DELETE("/:id", subscriberController.Delete)
	engineGrouped.PATCH("/contacted/:id", subscriberController.MarkAsContacted)
}
