package model

import "time"

type User struct {
	Id          int64  `json:"id" db:"id"`
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

type UserParams struct {
	TokenID      int64     `json:"token_id" db:"token_id"`
	RefreshToken string    `json:"refresh_token" db:"refresh_token"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

func ParseUserDTO(user User) UserDTO {
	return UserDTO{
		UUID:        user.UUID,
		Email:       user.Email,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
	}
}
