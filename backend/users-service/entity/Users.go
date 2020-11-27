package entity

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	MALE = iota
	FEMALE
	OTHER
)

const (
	ADMIN = iota
	USER
	DISABLE
	NOTACTIVE
)

const (
	SOFIA = iota
	QQ
	WECHAT
	GITHUB
)

type Users struct {
	Uid              bson.ObjectId `bson:"_id"`
	Oid              string        `bson:"oid"`
	Name             string        `bson:"name"`
	Nickname         string        `bson:"nickname"`
	Password         string        `bson:"password"`
	Email            string        `bson:"email"`
	Icon             string        `bson:"icon"`
	Gender           int8          `bson:"gender"`
	Role             int8          `bson:"role"`
	AccountType      int8          `bson:"account_type"`
	NotificationTime time.Time     `bson:"notification_time"`
}
