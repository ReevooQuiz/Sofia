package entity

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Answers struct {
	Aid            bson.ObjectId `bson:"_id"`
	Answerer       bson.ObjectId `bson:"answerer"`
	Qid            bson.ObjectId `bson:"qid"`
	Content        string        `bson:"content"`
	CommentCount   int64         `bson:"comment_count"`
	CriticismCount int64         `bson:"criticism_count"`
	LikeCount      int64         `bson:"like_count"`
	ApprovalCount  int64         `bson:"approval_count"`
	Time           time.Time     `bson:"time"`
}
