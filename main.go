package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zero2her0/driving-license/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/questions", handlers.AddQuestions)
	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
