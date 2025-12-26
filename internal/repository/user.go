package repository

import (
	"database/sql"

	"github.com/Dima5791/go-auth-service/internal/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	query := `
		SELECT id, email, password, role, created_at
		FROM users
		WHERE email = $1
	`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) FindByID(id int64) (*model.User, error) {
	user := &model.User{}

	query := `
		SELECT id, email, password, role, created_at
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) Create(user *model.User) error {
	query := `
		INSERT INTO users (email, password, role)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	return r.db.QueryRow(
		query,
		user.Email,
		user.PasswordHash,
		user.Role,
	).Scan(&user.ID)
}
