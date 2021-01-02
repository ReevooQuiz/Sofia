package entity

type Follows struct {
	Uid      int64 `json:"uid"`
	Follower int64 `json:"follower"`
	Time     int64 `json:"time"`
}
