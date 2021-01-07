package entity

import (
	"database/sql"
)

type Questions struct {
	Qid            int64         `json:"qid"`
	Raiser         int64         `json:"raiser"`
	Title          string        `json:"title"`
	Category       string        `json:"category"`
	AcceptedAnswer sql.NullInt64 `json:"accepted_answer"`
	AnswerCount    int64         `json:"answer_count"`
	ViewCount      int64         `json:"view_count"`
	FavoriteCount  int64         `json:"favorite_count"`
	Time           int64         `json:"time"`
	Labels         []string      `json:"labels"`
	Closed         bool          `json:"closed"`
	Scanned        bool          `json:"scanned"`
}
