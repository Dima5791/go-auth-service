package repository

import "github.com/Dima5791/go-auth-service/internal/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id int64) (*model.User, error)
}

type RefreshTokenRepository interface {
	Create(token *model.RefreshToken) error
	Find(token string) (*model.RefreshToken, error)
	Delete(token string) error
}
