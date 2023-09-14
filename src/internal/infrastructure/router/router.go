package router

import (
	"go-clean-arch/src/internal/infrastructure/handler"
	"go-clean-arch/src/internal/interface/controllers"

	"github.com/gin-gonic/gin"
)

func Setup(ge *gin.Engine, c controllers.Controllers) *gin.Engine {
	ge.GET("/", func(ctx *gin.Context) { handler.Home(ctx) })
	api := ge.Group("api")
	api.GET("/sample", func(ctx *gin.Context) { c.GetSample(ctx) })
	return ge
}
