package handlers

import (
	"fmt"

	"github.com/chalisekrishna418/driving-license/config"

	"github.com/gin-gonic/gin"
)

// AddQuestions function for reading the questions and prepare for putting into database
func AddQuestions(c *gin.Context) {
	db, err := config.MongoInit()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	question := c.Query("question")
	correctAns := c.Query("correct_answer")
	option1 := c.Query("option1")
	option2 := c.Query("option2")
	option3 := c.Query("option3")

	coll := db.C("questions")
	data := gin.H{
		"question":       question,
		"correct_answer": correctAns,
		"option1":        option1,
		"option2":        option2,
		"option3":        option3,
	}
	insErr := coll.Insert(data)
	if insErr != nil {
		fmt.Printf("Error: %s", insErr)
	}
	c.JSON(200, gin.H{
		"status": "question successfully inserted !!!",
	})
}
