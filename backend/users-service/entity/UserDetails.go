package entity

type UserDetails struct {
	Uid  int64  `bson:"uid"`
	Icon string `bson:"icon"`
}
