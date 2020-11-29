package models

type User struct {
	UserID     int64  `db:"user_id"`
	Username   string `db:"username"`
	Password   string `db:"password"`
	Token      string
	CreateTime int64 `db:"create_time"`
}
