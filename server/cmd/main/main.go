package main

import (
	"log"

	"github.com/dxtym/anon/server/internal/api"
	"github.com/dxtym/anon/server/internal/utils"
)

func main() {
	cfg, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatalf("cannot load config: %s", err.Error())
	}

	server := api.NewServer(cfg)
	if err := server.Start(); err != nil {
		log.Fatalf("cannot start server: %s", err.Error())
	}
}
