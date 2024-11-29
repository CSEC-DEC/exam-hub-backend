package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/devasherr/exam_hub/helpers"
	"github.com/devasherr/exam_hub/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection *mongo.Collection = OpenCollection(Client, "users")
var examsCollection *mongo.Collection = OpenCollection(Client, "exams")

func Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var loginRequest models.LoginRequest
	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var foundUser models.User
	err := usersCollection.FindOne(ctx, bson.M{"username": loginRequest.UserName}).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	if !helpers.CheckHashedPassword(foundUser.Password, loginRequest.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	// generate access token
	token, err := helpers.GenerateJWT(foundUser.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate access token"})
		return
	}

	c.SetCookie("access_token", token, 3600, "/", "", true, false)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func GetExams(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var data []bson.M
	cursor, err := examsCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := cursor.All(ctx, &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetExam(c *gin.Context) {
	id := c.Params.ByName("id")
	docId, _ := primitive.ObjectIDFromHex(id)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result bson.M
	if err := examsCollection.FindOne(ctx, bson.M{"_id": docId}).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
