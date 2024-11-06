package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/dxtym/anon/server/internal/models"
	"github.com/dxtym/anon/server/internal/utils"
)

var cfg, _ = utils.LoadConfig("../..")

func TestUserRepoCreate(t *testing.T) {
	s, down := TestStore(t, cfg)
	defer down()
	user, err := s.User().Create(&models.User{
		Username: "admin",
		Password: "password",
	})
	assert.NoError(t, err)
}

func TestUserRepoFind(t *testing.T) {
	s, down := TestStore(t, cfg)
	defer down()
	user, err := s.User().Find("admin")
	assert.NoError(t, err)
}