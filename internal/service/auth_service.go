package service

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/salesforceanton/files-portal/internal/config"
	"github.com/salesforceanton/files-portal/internal/repository"
)

type AuthService struct {
	cfg  config.Config
	repo repository.Authorization
}

type TokenClaims struct {
	*jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&TokenClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Invalid Signing Method")
			}
			return []byte(s.cfg.TokenSecret), nil
		},
	)

	if err != nil {
		return 0, errors.New("Error with parsing Access Token")
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("Error with parsing Access Token")
	}

	return claims.UserId, nil
}
