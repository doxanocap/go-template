package model

type User struct {
	Id          int    `json:"id" db:"id"`
	UUID        string `json:"uuid" db:"uuid"`
	Email       string `json:"email" db:"email"`
	Username    string `json:"username" db:"username"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Password    string `json:"-" db:"password"`
}

type UserDTO struct {
	UUID        string `json:"uuid" db:"uuid"`
	Email       string `json:"email" db:"email"`
	Username    string `json:"username" db:"username"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Password    string `json:"-" db:"password"`
}
