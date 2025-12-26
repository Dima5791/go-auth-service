package jwt

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessClaims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Manager struct {
	accessSecret  string
	refreshSecret string
	accessTTL     time.Duration
	refreshTTL    time.Duration
}

func NewManager(
	accessSecret string,
	refreshSecret string,
	accessTTL time.Duration,
	refreshTTL time.Duration,
) *Manager {
	return &Manager{
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
		accessTTL:     accessTTL,
		refreshTTL:    refreshTTL,
	}
}

func (m *Manager) GenerateRefreshToken(userID int64) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(userID, 10),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.refreshTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.refreshSecret))
}

func (m *Manager) ParseAccessToken(tokenStr string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(m.accessSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)
	return claims, nil
}

func (m *Manager) ParseRefreshToken(tokenStr string) (int64, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(m.refreshSecret), nil
		},
	)

	if err != nil || !token.Valid {
		return 0, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return 0, errors.New("invalid refresh token")
	}

	userID, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (m *Manager) ValidateAccessToken(tokenString string) (*AccessClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&AccessClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(m.accessSecret), nil
		},
	)

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*AccessClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func (m *Manager) GenerateAccessToken(userID int64, role string) (string, error) {
	claims := AccessClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.accessTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.accessSecret))
}

func (m *Manager) RefreshTTL() time.Duration {
	return m.refreshTTL
}
