package main

import (
	"github.com/chalisekrishna418/driving-license/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/questions", handlers.AddQuestions )
	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
