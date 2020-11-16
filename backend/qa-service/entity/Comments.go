package entity

import "time"

type Comments struct {
	Cmid    string    `json:"cmid"`
	Uid     string    `json:"uid"`
	Aid     string    `json:"aid"`
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}
