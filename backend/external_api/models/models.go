package models

type User struct {
	Id       int32  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
