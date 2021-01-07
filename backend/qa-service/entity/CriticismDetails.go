package entity

type CriticismDetails struct {
	Ctid    int64  `bson:"_id"`
	Content string `bson:"content"`
}
