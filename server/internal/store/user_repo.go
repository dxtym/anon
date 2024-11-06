package store

import "github.com/dxtym/anon/server/internal/models"

type UserRepo struct {
	store *Store
}

func (ur *UserRepo) Create(user *models.User) (*models.User, error) {
	if err := ur.store.DB.QueryRow(
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id;",
		user.Username,
		user.Password,
	).Scan(&user.Id); err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepo) Find(username string) (*models.User, error) {
	user := &models.User{}
	if err := ur.store.DB.QueryRow(
		"SELECT id, username, password FROM users WHERE username = $1",
		username,
	).Scan(user.Id, user.Username, user.Password); err != nil {
		return nil, err
	}
	return user, nil
}