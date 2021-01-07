package entity

type CommentDetails struct {
	Cmid    int64  `bson:"_id"`
	Content string `bson:"content"`
}
