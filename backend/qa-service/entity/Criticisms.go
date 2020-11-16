package entity

import "time"

type Criticisms struct {
	Ctid    string    `json:"ctid"`
	Uid     string    `json:"uid"`
	Aid     string    `json:"aid"`
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}
