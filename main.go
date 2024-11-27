package main

import (
	"github.com/devasherr/exam_hub/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.POST("/login", routes.Login)

	router.GET("/exams", routes.GetExams)
	router.GET("/exam/:id", routes.GetExam)

	router.Run(":8800")
}
