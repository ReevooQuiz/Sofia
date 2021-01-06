package entity

type Answers struct {
	Aid            int64 `json:"aid"`
	Answerer       int64 `json:"answerer"`
	Qid            int64 `json:"qid"`
	ViewCount      int64 `json:"view_count"`
	CommentCount   int64 `json:"comment_count"`
	CriticismCount int64 `json:"criticism_count"`
	LikeCount      int64 `json:"like_count"`
	ApprovalCount  int64 `json:"approval_count"`
	Time           int64 `json:"time"`
	Scanned        int8  `json:"scanned"`
}
