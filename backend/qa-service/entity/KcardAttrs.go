package entity

type KcardAttrs struct {
	Kid    int64  `json:"kid"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	Origin string `json:"origin"`
}
