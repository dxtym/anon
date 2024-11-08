package store

import (
	"github.com/dxtym/anon/server/internal/models"
	"github.com/dxtym/anon/server/internal/utils"
	"github.com/go-playground/validator/v10"
)

type UserService struct {
	store *Store
}

func (us *UserService) Create(user *models.User) (*models.User, error) {
	validate := validator.New()
	if err := validate.Struct(us); err != nil {
		return nil, err
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	if err := us.store.db.QueryRow(
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id;", 
		user.Username, hashed,
	).Scan(&user.Id); err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) Find(username string) (*models.User, error) {
	user := &models.User{}
	if err := us.store.db.QueryRow(
		"SELECT id, username, password FROM users WHERE username = $1",
		username,
	).Scan(user.Id, user.Username, user.Password); err != nil {
		return nil, err
	}

	return user, nil
}
