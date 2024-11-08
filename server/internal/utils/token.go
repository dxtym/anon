package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const minSecretKeySize = 32

var (
	ErrInvalidKeySize    = errors.New("invalid key size")
	ErrInvalidsignMethod = errors.New("invalid sign method")
	ErrInvalidToken      = errors.New("invalid token")
)

type Maker struct {
	secretKey string
}

func NewMaker(secretKey string) (*Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, ErrInvalidKeySize
	}

	return &Maker{secretKey: secretKey}, nil
}

func (m *Maker) CreateToken(username string, duration time.Duration) (string, error) {
	claim, err := NewClaim(username, duration)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(m.secretKey))
}

func (m *Maker) VerifyToken(token string) (*Claim, error) {
	t, err := jwt.ParseWithClaims(token, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidsignMethod
		}

		return []byte(m.secretKey), nil
	})

	if err != nil {
		if ver, ok := err.(*jwt.ValidationError); ok && errors.Is(ver.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}

		return nil, ErrInvalidToken
	}

	claim, ok := t.Claims.(*Claim)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claim, nil
}
