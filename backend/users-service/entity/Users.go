package entity

const (
	MALE int8 = iota
	FEMALE
	OTHER
)

const (
	ADMIN int8 = iota
	USER
	DISABLE
	NOT_ACTIVE
)

const (
	SOFIA int8 = iota
	QQ
	WECHAT
	GITHUB
)

type Users struct {
	Uid              int64  `json:"uid"`
	Oid              string `json:"oid"`
	Name             string `json:"name"`
	Nickname         string `json:"nickname"`
	Salt             string `json:"salt"`
	HashPassword     string `json:"hash_password"`
	Email            string `json:"email"`
	Gender           int8   `json:"gender"`
	Profile          string `json:"profile"`
	Role             int8   `json:"role"`
	AccountType      int8   `json:"account_type"`
	ActiveCode       int64  `json:"active_code"`
	PasswdCode       int64  `json:"passwd_code"`
	Exp              int64  `json:"exp"`
	FollowerCount    int64  `json:"follower_count"`
	FollowingCount   int64  `json:"following_count"`
	QuestionCount    int64  `json:"question_count"`
	AnswerCount      int64  `json:"answer_count"`
	LikeCount        int64  `json:"like_count"`
	ApprovalCount    int64  `json:"approval_count"`
	NotificationTime int64  `json:"notification_time"`
}
