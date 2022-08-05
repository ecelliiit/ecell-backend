package utils

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SendResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{"message": message, "data": data})
}

func PrettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}

func GenerateFindOptions(page_number int64, limit int64, sortBy string, sortOrder int64) *options.FindOptions {
	findOptions := options.Find()
	if page_number == 0 {
		page_number = 1
	}
	if limit == 0 {
		limit = 10
	}
	findOptions.SetSkip((page_number - 1) * limit)
	findOptions.SetLimit(limit)
	if sortBy != "" {
		if sortOrder == 0 {
			findOptions.SetSort(bson.M{sortBy: 1})
		} else {
			findOptions.SetSort(bson.M{sortBy: sortOrder})
		}
	}
	return findOptions
}
