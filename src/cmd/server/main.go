package main

import (
	"flag"
	"go-clean-arch/src/internal/configs"
	"go-clean-arch/src/internal/infrastructure/server"
	"log"
)

func main() {
	flag.Parse()
	cfg, err := configs.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	server.Start(cfg)
}
