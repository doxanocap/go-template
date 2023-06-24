package model

type AuthResponse struct {
	User   UserDTO `json:"user"`
	Tokens Tokens  `json:"tokens"`
}

type AuthUser struct {
	Id          int    `json:"Id"`
	Username    string `json:"Username"`
	Email       string `json:"Email"`
	IsActivated bool   `json:"IsActivated"`
	Password    []byte `json:"-"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignInRequest struct {
}

type SignUp struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Password    []byte `json:"-"`
}
