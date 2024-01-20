package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
	Expires      int    `json:"expires"`
	ExpiresAt    int    `json:"expires_at"`
}

type TokenManager interface {
	GenerateTokens(userID string, userSign string) (Tokens, error)
	Parse(token string) (string, string, error)
	NewToken(userID string, userSign string, ttl time.Duration) (string, error)
}

type Manager struct {
	secretKey  string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewManager(secretKey string, accessTTL time.Duration, refreshTTL time.Duration) (*Manager, error) {
	if secretKey == "" {
		return nil, errors.New("empty secret key")
	}

	return &Manager{
		secretKey:  secretKey,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}, nil
}

func (m *Manager) NewToken(userID string, userSign string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Issuer:    userSign,
		Subject:   userID,
	})

	return token.SignedString([]byte(m.secretKey))
}

func (m *Manager) Parse(t string) (string, string, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.secretKey), nil
	})
	if err != nil {
		return "", "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("cant get claims from token")
	}
	return claims["sub"].(string), claims["iss"].(string), nil
}

func (m *Manager) GenerateTokens(userID string, userSign string) (Tokens, error) {
	expiresAt := time.Now().Add(m.accessTTL)
	accessToken, err := m.NewToken(userID, userSign, m.accessTTL)
	if err != nil {
		return Tokens{}, err
	}
	refreshToken, err := m.NewToken(userID, userSign, m.refreshTTL)
	if err != nil {
		return Tokens{}, err
	}
	return Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expires:      int(m.accessTTL.Seconds()),
		ExpiresAt:    int(expiresAt.Unix()),
		UserID:       userID,
	}, nil
}
