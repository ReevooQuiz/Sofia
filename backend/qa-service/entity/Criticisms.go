package entity

import (
	"gopkg.in/mgo.v2/bson"
)

type Criticisms struct {
	Ctid    bson.ObjectId `bson:"_id"`
	Uid     bson.ObjectId `bson:"uid"`
	Aid     bson.ObjectId `bson:"aid"`
	Content string        `bson:"content"`
	Time    int64     `bson:"time"`
}
