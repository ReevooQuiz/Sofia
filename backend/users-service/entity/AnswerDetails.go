package entity

type AnswerDetails struct {
	Aid     int64  `bson:"aid"`
	Content string `bson:"content"`
}
