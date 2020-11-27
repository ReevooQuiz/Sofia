package entity

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Comments struct {
	Cmid    bson.ObjectId `bson:"_id"`
	Uid     bson.ObjectId `bson:"uid"`
	Aid     bson.ObjectId `bson:"aid"`
	Content string        `bson:"content"`
	Time    time.Time     `bson:"time"`
}
