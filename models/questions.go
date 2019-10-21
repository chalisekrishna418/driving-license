package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Questions struct
type Questions struct {
	QuestionID    bson.ObjectId `bson:"_id" json:"id"`
	Question      string        `bson:"question" json:"question" form:"question" binding:"required"`
	CorrectAnswer string        `bson:"ans" json:"ans" form:"ans" binding:"required"`
	Options       []string      `bson:"options" json:"options" form:"options"`
}
