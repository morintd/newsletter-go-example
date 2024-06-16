package core

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var CoreModule = fx.Options(
	fx.Provide(initialize),
)

func initialize() *gin.Engine {
	return gin.Default()
}
