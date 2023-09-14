package server

import (
	"fmt"
	"go-clean-arch/src/internal/configs"
	"go-clean-arch/src/internal/infrastructure/registry"
	"go-clean-arch/src/internal/infrastructure/router"
	"log"

	"github.com/gin-gonic/gin"
)

func Start(cfg *configs.Config) {
	gin.SetMode(cfg.Server.Mode)
	ge := gin.New()
	ge.Use(gin.Recovery())

	r := registry.NewRegistry()
	router.Setup(ge, r.NewAppControllers())

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Panic(ge.Run(addr))
}
