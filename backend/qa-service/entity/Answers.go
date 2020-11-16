package entity

import "time"

type Answers struct {
	Aid            string    `json:"aid"`
	Answerer       string    `json:"answerer"`
	Qid            string    `json:"qid"`
	Content        string    `json:"content"`
	CommentCount   int64     `json:"comment_count"`
	CriticismCount int64     `json:"criticism_count"`
	LikeCount      int64     `json:"like_count"`
	ApprovalCount  int64     `json:"approval_count"`
	Time           time.Time `json:"time"`
}
