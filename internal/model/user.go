package model

type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    []byte `json:"-"`
}
