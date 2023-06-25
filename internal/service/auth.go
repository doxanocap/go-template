package service

import (
	"app/internal/manager/interfaces"
	"app/internal/model"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthService struct {
	manager interfaces.IManager
}

func InitAuthService(manager interfaces.IManager) *AuthService {
	return &AuthService{
		manager: manager,
	}
}

func (as *AuthService) NewPairTokens(user model.UserDTO) (result *model.Tokens, err error) {
	payload, err := json.Marshal(user)
	if err != nil {
		return
	}

	accessClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		Issuer:    string(payload),
	})
	refreshClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 20)),
		Issuer:    string(payload),
	})
	accessToken, _ := accessClaim.SignedString([]byte("JWT_SECRET_KEY"))
	refreshToken, _ := refreshClaim.SignedString([]byte("JWT_SECRET_KEY"))

	result = &model.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return result, nil
}
