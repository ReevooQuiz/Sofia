package entity

import (
	"database/sql"
)

type Questions struct {
	Qid            int64         `bson:"_id"`
	Raiser         int64         `bson:"raiser"`
	Title          string        `bson:"title"`
	Category       string        `bson:"category"`
	AcceptedAnswer sql.NullInt64 `bson:"accepted_answer"`
	AnswerCount    int64         `bson:"answer_count"`
	ViewCount      int64         `bson:"view_count"`
	FavoriteCount  int64         `bson:"favorite_count"`
	Time           int64     `bson:"time"`
	Labels         []string      `bson:"labels"`
	Scanned        bool          `bson:"scanned"`
}
