package entity

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Questions struct {
	Qid            bson.ObjectId `bson:"_id"`
	Raiser         bson.ObjectId `bson:"raiser"`
	Title          string        `bson:"title"`
	Content        string        `bson:"content"`
	Category       string        `bson:"category"`
	AcceptedAnswer bson.ObjectId `bson:"accepted_answer"`
	AnswerCount    int64         `bson:"answer_count"`
	ViewCount      int64         `bson:"view_count"`
	FavoriteCount  int64         `bson:"favorite_count"`
	Time           time.Time     `bson:"time"`
}
