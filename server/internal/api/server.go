package api

import (
	"net/http"

	"github.com/dxtym/anon/server/internal/store"
	"github.com/dxtym/anon/server/internal/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config utils.Config
	router *gin.Engine
	store  *store.Store
}

func NewServer(cfg utils.Config) *Server {
	server := &Server{
		config: cfg,
		router: gin.Default(),
		store:  store.NewStore(cfg),
	}
	server.setUpRouting()
	if err := server.setUpStore(); err != nil {
		return nil
	}
	return server
}

func (s *Server) Start() error {
	if err := s.router.Run(s.config.Address); err != nil {
		return err
	}
	return nil
}

func (s *Server) setUpRouting() {
	s.router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello, world!")
	})
}

func (s *Server) setUpStore() error {
	if err := s.store.Open(); err != nil {
		return err
	}
	return nil
}
