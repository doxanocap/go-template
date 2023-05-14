package model

type AuthResponseModel struct {
	Status       int    `json:"Status"`
	Message      string `json:"Error"`
	AccessToken  string `json:"AccessToken"`
	RefreshToken string `json:"RefreshToken"`
	User         User   `json:"User"`
}

type Error struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
}
