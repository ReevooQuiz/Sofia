package entity

const (
	ADMIN = iota
	USER
	DISABLE
	NOTACTIVE
)

type Users struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     int8   `json:"role"`
}
