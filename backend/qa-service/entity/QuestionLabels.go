package entity

import "gopkg.in/mgo.v2/bson"

type QuestionLabels struct {
	Qid bson.ObjectId `json:"qid"`
	Lid int64         `json:"lid"`
}
