package models

type ParamUser struct {
	UserId   int    `db:"id"`
	Username string `db:"username"`
}
