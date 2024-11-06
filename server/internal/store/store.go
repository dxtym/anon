package store

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/dxtym/anon/server/internal/utils"
)

type Store struct {
	DB *sql.DB 
	Config utils.Config
	UserRepo *UserRepo
}

func NewStore(cfg utils.Config) *Store {
	return &Store{Config: cfg}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.Config.DatabaseURL)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	s.DB = db
	return nil
}

func (s *Store) Close() {
	if s.DB != nil {
		s.DB.Close()
	}
}

func (s *Store) User() *UserRepo {
	if s.UserRepo != nil {
		return s.UserRepo
	}
	s.UserRepo = &UserRepo{store: s}
	return s.UserRepo
}