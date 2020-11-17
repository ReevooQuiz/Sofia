package entity

import "gopkg.in/mgo.v2/bson"

type KcardAttrs struct {
	Kid    int64         `json:"kid"`
	Name   string        `json:"name"`
	Value  string        `json:"value"`
	Origin bson.ObjectId `json:"origin"`
}
