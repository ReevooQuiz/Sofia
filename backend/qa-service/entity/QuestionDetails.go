package entity

type QuestionDetails struct {
	Qid        int64  `bson:"_id"`
	Content    string `bson:"content"`
	PictureUrl string `bson:"picture_url"`
	Head       string `bson:"head"`
}
