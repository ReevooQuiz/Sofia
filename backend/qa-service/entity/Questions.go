package entity

import "time"

type Questions struct {
	Qid            string    `json:"qid"`
	Raiser         string    `json:"raiser"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Category       string    `json:"category"`
	AcceptedAnswer string    `json:"accepted_answer"`
	AnswerCount    int64     `json:"answer_count"`
	ViewCount      int64     `json:"view_count"`
	FavoriteCount  int64     `json:"favorite_count"`
	Time           time.Time `json:"time"`
	Labels         []Labels  `json:"labels"`
}
