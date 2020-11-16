package entity

type Kcards struct {
	Kid   int64        `json:"kid"`
	Title string       `json:"title"`
	Attrs []KcardAttrs `json:"attrs"`
}
