package entity

type AnswerDetails struct {
	Aid        int64  `bson:"_id"`
	Content    string `bson:"content"`
	PictureUrl string `bson:"picture_url"`
	Head       string `json:"head"`
}
