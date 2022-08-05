package controllers

import (
	"github.com/ecelliiit/ecell-backend/models"
	"github.com/ecelliiit/ecell-backend/repository"
	"github.com/ecelliiit/ecell-backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubscriberController struct {
}

func NewSubscriberController() *SubscriberController {
	return &SubscriberController{}
}

func (*SubscriberController) Create(c *gin.Context) {
	subscriber := &models.Subscriber{}
	err := c.BindJSON(subscriber)
	if err != nil {
		utils.SendResponse(c, 422, "Invalid payload", nil)
		return
	}

	result, err := repository.CreateSubscriber(subscriber)
	if err != nil {
		utils.SendResponse(c, 500, "Error in adding to database", nil)
		return
	}

	utils.SendResponse(c, 201, "Susbcriber created", result)
}

func (*SubscriberController) GetAll(c *gin.Context) {
	result, err := repository.GetAllSubscribers()
	if err != nil {
		utils.SendResponse(c, 500, "Error in fetching from database", nil)
		return
	}

	utils.SendResponse(c, 200, "All subscribers fetched successfully", result)
}

func (*SubscriberController) GetById(c *gin.Context) {
	idString := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		utils.SendResponse(c, 422, "invalid object id", nil)
		return
	}

	result, err := repository.GetSubscriberById(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.SendResponse(c, 400, "No subscriber found with given id", nil)
			return
		}
		utils.SendResponse(c, 500, "Error in fetching data from db", nil)
		return
	}

	utils.SendResponse(c, 200, "Susbcriber fetched successfully", result)
}

func (*SubscriberController) Delete(c *gin.Context) {
	idString := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		utils.SendResponse(c, 422, "invalid object id", nil)
		return
	}

	err = repository.DeleteSubscriberByID(id)
	if err != nil {
		utils.SendResponse(c, 500, "Error in deleting data from db", nil)
		return
	}

	utils.SendResponse(c, 200, "Susbcriber deleted successfully", nil)
}

func (*SubscriberController) MarkAsContacted(c *gin.Context) {
	idString := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		utils.SendResponse(c, 422, "invalid object id", nil)
		return
	}

	result, err := repository.MarkAsContacted(id)
	if err != nil {
		utils.SendResponse(c, 500, "Error in marking subscriber as complete from db", nil)
		return
	}

	utils.SendResponse(c, 200, "Susbcriber marked as contacted successfully", result)
}
