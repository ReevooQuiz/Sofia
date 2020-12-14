package entity

const (
	MALE = iota
	FEMALE
	OTHER
)

const (
	ADMIN = iota
	USER
	DISABLE
	NOT_ACTIVE
)

const (
	SOFIA = iota
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
	Role             int8   `json:"role"`
	AccountType      int8   `json:"account_type"`
	ActiveCode       int64  `json:"active_code"`
	PasswdCode       int64  `json:"passwd_code"`
	Exp              int64  `json:"exp"`
	FollowerCount    int64  `json:"follower_count"`
	FollowingCount   int64  `json:"following_count"`
	NotificationTime int64  `json:"notification_time"`
}
