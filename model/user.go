package model

var UserArray []User

type User struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"username"`
	Surname  string `json:"surname"`
	IsActive bool   `json:"isactive"`
}

type Users struct {
	Users []User `json:"users"`
}
