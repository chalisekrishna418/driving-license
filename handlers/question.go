package handlers

import(
	"github.com/gin-gonic/gin"	
)

// AddQuestions function for reading the questions and prepare for putting into database
func AddQuestions(c *gin.Context){
	question := c.Query("question")
	correctAns := c.Query("correct_answer")
	options := c.Query("options")

	c.JSON(200, gin.H{
		"question": question,
		"answer": correctAns,
		"options": options,
	})
}
