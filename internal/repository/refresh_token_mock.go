package repository

import (
	"errors"
	"time"

	"github.com/Dima5791/go-auth-service/internal/model"
)

type refreshTokenMockRepo struct {
	data map[string]*model.RefreshToken
}

func NewRefreshTokenMockRepo() RefreshTokenRepository {
	return &refreshTokenMockRepo{
		data: make(map[string]*model.RefreshToken),
	}
}

func (r *refreshTokenMockRepo) Create(token *model.RefreshToken) error {
	r.data[token.Token] = token
	return nil
}

func (r *refreshTokenMockRepo) Find(token string) (*model.RefreshToken, error) {
	t, ok := r.data[token]
	if !ok || t.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("refresh token invalid")
	}
	return t, nil
}

func (r *refreshTokenMockRepo) Delete(token string) error {
	delete(r.data, token)
	return nil
}

func (r *refreshTokenMockRepo) DeleteByUser(userID int64) error {
	for k, v := range r.data {
		if v.UserID == userID {
			delete(r.data, k)
		}
	}
	return nil
}
