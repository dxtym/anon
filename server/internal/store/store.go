package store

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/dxtym/anon/server/internal/utils"
)

type Store struct {
	config utils.Config
	db     *sql.DB
}

func NewStore(cfg utils.Config) *Store {
	return &Store{
		config: cfg,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	if s.db != nil {
		s.db.Close()
	}
}
