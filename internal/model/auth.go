package model

type AuthResponse struct {
	Status       int      `json:"Status"`
	Message      string   `json:"Error"`
	AccessToken  string   `json:"AccessToken"`
	RefreshToken string   `json:"RefreshToken"`
	User         AuthUser `json:"User"`
}

type AuthUser struct {
	Id          int    `json:"Id"`
	Username    string `json:"Username"`
	Email       string `json:"Email"`
	IsActivated bool   `json:"IsActivated"`
	Password    []byte `json:"-"`
}
