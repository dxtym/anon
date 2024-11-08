package utils

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredToken = errors.New("token expired")

type Claim struct {
	Id        uuid.UUID
	Username  string
	CreatedAt time.Time
	ExpireAt  time.Time
}

func NewClaim(username string, duration time.Duration) (*Claim, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Claim{
		Id:        id,
		Username:  username,
		CreatedAt: time.Now(),
		ExpireAt:  time.Now().Add(duration),
	}, nil
}

func (c *Claim) Valid() error {
	if time.Now().After(c.ExpireAt) {
		return ErrExpiredToken
	}

	return nil
}
