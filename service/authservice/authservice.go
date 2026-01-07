package authservice

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"q-game-app/entity"
	"strings"
	"time"
)

// AuthParser interface
type AuthParser interface {
	ParseAccessToken(token string) (entity.User, error)
}

type Config struct {
	SignKey               string
	AccessExpirationTime  time.Duration
	RefreshExpirationTime time.Duration
	AccessSubject         string
	RefreshSubject        string
}

type Service struct {
	config Config
}

// New now returns pointer
func New(cfg Config) Service {
	return Service{
		config: cfg,
	}
}

// CreateAccessToken creates JWT access token
func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.config.AccessSubject, s.config.AccessExpirationTime)
}

// CreateRefreshToken creates JWT refresh token
func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.config.RefreshSubject, s.config.RefreshExpirationTime)
}

// ParseToken parses bearer token and validates it
func (s Service) ParseToken(bearerToken string) (*Claims, error) {
	if bearerToken == "" {
		return nil, fmt.Errorf("token is empty")
	}

	// Remove "Bearer " prefix
	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if token == nil || token.Claims == nil {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Check expiration
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}

// createToken generates a JWT token
func (s Service) createToken(userID uint, subject string, expireDuration time.Duration) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.SignKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// Claims struct
// Claims struct
type Claims struct {
	jwt.RegisteredClaims
	UserID uint
}

// Valid validates the claims (expiration, issuedAt, etc.)
