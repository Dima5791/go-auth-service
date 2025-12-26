package repository

import (
	"errors"

	"github.com/Dima5791/go-auth-service/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type userMockRepo struct {
	users map[string]*model.User
}

func NewUserMockRepo() UserRepository {
	hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	return &userMockRepo{
		users: map[string]*model.User{
			"test@test.com": {
				ID:           1,
				Email:        "test@test.com",
				PasswordHash: string(hash),
			},
		},
	}
}

func (r *userMockRepo) FindByEmail(email string) (*model.User, error) {
	user, ok := r.users[email]
	if !ok {
		return nil, errors.New("not found")
	}
	return user, nil
}

func (r *userMockRepo) FindByID(id int64) (*model.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("not found")
}

func (r *userMockRepo) Create(user *model.User) error {
	r.users[user.Email] = user
	return nil
}
