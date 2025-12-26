package handler

import "github.com/Dima5791/go-auth-service/pkg/jwt"

type AuthService interface {
	Login(email, password string) (*jwt.TokenPair, error)
	Refresh(refreshToken string) (string, error)
}
