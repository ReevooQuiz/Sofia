package entity

type UserDetails struct {
	Uid  int64  `bson:"_id"`
	Icon string `bson:"icon"`
}
