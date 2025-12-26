package repository

import (
	"database/sql"

	"github.com/Dima5791/go-auth-service/internal/model"
)

type RefreshTokenRepo struct {
	db *sql.DB
}

func NewRefreshTokenRepo(db *sql.DB) *RefreshTokenRepo {
	return &RefreshTokenRepo{db: db}
}

func (r *RefreshTokenRepo) Create(rt *model.RefreshToken) error {
	return nil
}

func (r *RefreshTokenRepo) Find(token string) (*model.RefreshToken, error) {
	return nil, nil
}

func (r *RefreshTokenRepo) Delete(token string) error {
	return nil
}
