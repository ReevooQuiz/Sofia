package entity

type Comments struct {
	Cmid int64 `json:"cmid"`
	Uid  int64 `json:"uid"`
	Aid  int64 `json:"aid"`
	Time int64 `json:"time"`
}
