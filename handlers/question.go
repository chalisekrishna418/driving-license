package handlers

import (
	"fmt"

	"github.com/zero2her0/driving-license/models"
	"github.com/zero2her0/driving-license/utils"

	"github.com/gin-gonic/gin"
)

// AddQuestions function for reading the questions and prepare for putting into database
func AddQuestions(c *gin.Context) {

	ques := models.Questions{}
	bindErr := c.ShouldBind(&ques)
	if bindErr != nil {
		fmt.Printf("Error: %s", bindErr)
	}

	insErr := utils.Insert(ques, "questions")
	if insErr != nil {
		fmt.Printf("Error: %s", insErr)
	}
	c.JSON(200, gin.H{
		"status": "question successfully inserted !!!",
	})
}
