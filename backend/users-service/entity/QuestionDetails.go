package entity

type QuestionDetails struct {
	Qid        int64  `bson:"_id"`
	Title      string `bson:"title"`
	Content    string `bson:"content"`
	PictureUrl string `bson:"picture_url"`
	Head       string `json:"head"`
}
