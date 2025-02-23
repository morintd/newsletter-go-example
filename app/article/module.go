package article

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func configure(r *gin.Engine, controller *ArticleController) {
	controller.Configure(r)
}

var ArticleModule = fx.Options(
	fx.Provide(NewInMemoryArticleRepository),
	fx.Provide(NewArticleController),
	fx.Invoke(configure),
)
