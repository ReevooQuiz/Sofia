package entity

type QuestionDetails struct {
	Qid     int64  `bson:"qid"`
	Title   string `bson:"title"`
	Content string `bson:"content"`
}
