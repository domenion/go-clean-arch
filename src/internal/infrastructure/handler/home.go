package handler

import (
	"fmt"
	"go-clean-arch/src/internal/configs"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	cfg, _ := configs.GetConfig()
	ctx.Writer.WriteString(fmt.Sprintf("%s\nVersion: %s\nEnvironment: %s", cfg.App.Name, cfg.App.Version, cfg.App.Environment))
}
