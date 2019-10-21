package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zero2her0/driving-license/handlers"
)

// InitRoutes defines routes
func InitRoutes() {
	r := gin.Default()

	r.POST("/questions", handlers.AddQuestions)

	r.Run(":8000") // listens and serve on 0.0.0.0:8080
}
