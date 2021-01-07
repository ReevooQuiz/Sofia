package rpc

type UserInfo struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
}

type UsersRPC interface {
	GetUserInfos(uids []int64) (result []UserInfo, err error)
	ParseToken(token string) (successful bool, uid int64, role int8)
}
