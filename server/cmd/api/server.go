package api

import (
	"net/http"

	"github.com/dxtym/anon/server/internal/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config utils.Config
	router *gin.Engine
}

func NewServer(cfg utils.Config) *Server {
	server := &Server{
		config: cfg,
	}
	server.setUpRouting()
	return server
}

func (s *Server) Start() error {
	if err := s.router.Run(s.config.Address); err != nil {
		return err
	}
	return nil
}

func (s *Server) setUpRouting() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello, world!")
	})
	s.router = router
}