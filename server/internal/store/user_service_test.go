package store

import (
	"testing"

	"github.com/dxtym/anon/server/internal/models"
	"github.com/dxtym/anon/server/internal/utils"
	"github.com/stretchr/testify/assert"
)

var cfg, _ = utils.LoadConfig("../..")

func testUser() *models.User {
	return &models.User{
		Username: "admin",
		Password: "password",
	}
}

func TestUserRepoCreate(t *testing.T) {
	s, down := TestStore(t, cfg)
	defer down()

	user, err := s.User().Create(testUser())
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepoFind(t *testing.T) {
	s, down := TestStore(t, cfg)
	defer down()

	user, err := s.User().Find(testUser().Username)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
