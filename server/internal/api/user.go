package api

import (
	"net/http"

	"github.com/dxtym/anon/server/internal/models"
	"github.com/dxtym/anon/server/internal/utils"
	"github.com/gin-gonic/gin"
)

type registerUserRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}


func (s *Server) registerUser(ctx *gin.Context) {
	var req registerUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	user := &models.User{
		Username: req.Username,
		Password: hashed,
	}
	u, err := s.store.User().Create(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

// func (s *Server) loginUser(ctx *gin.Context) {
	
// }