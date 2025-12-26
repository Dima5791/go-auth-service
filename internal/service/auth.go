package service

import (
	"errors"
	"time"

	"github.com/Dima5791/go-auth-service/internal/model"
	"github.com/Dima5791/go-auth-service/internal/repository"
	"github.com/Dima5791/go-auth-service/pkg/jwt"
)

type authService struct {
	userRepo    repository.UserRepository
	refreshRepo repository.RefreshTokenRepository
	jwtMgr      *jwt.Manager
}

func NewAuthService(
	userRepo repository.UserRepository,
	refreshRepo repository.RefreshTokenRepository,
	jwtMgr *jwt.Manager,
) AuthService

func (s *authService) Refresh(refreshToken string) (*jwt.TokenPair, error) {

	rt, err := s.refreshRepo.Find(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	if time.Now().After(rt.ExpiresAt) {
		_ = s.refreshRepo.Delete(refreshToken)
		return nil, errors.New("refresh token expired")
	}

	user, err := s.userRepo.FindByID(rt.UserID)
	if err != nil {
		return nil, err
	}

	if err := s.refreshRepo.Delete(refreshToken); err != nil {
		return nil, err
	}

	access, err := s.jwtMgr.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	newRefresh, err := s.jwtMgr.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	if err := s.refreshRepo.Create(&model.RefreshToken{
		UserID:    user.ID,
		Token:     newRefresh,
		ExpiresAt: time.Now().Add(s.jwtMgr.RefreshTTL()),
	}); err != nil {
		return nil, err
	}

	return &jwt.TokenPair{
		AccessToken:  access,
		RefreshToken: newRefresh,
	}, nil
}
