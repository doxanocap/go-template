package service

import "app/internal/model"

type IAuthService interface {
	NewPairTokens(user model.UserDTO) (result *model.Tokens, err error)
}
