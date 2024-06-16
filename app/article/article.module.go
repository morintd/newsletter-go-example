package article

import (
	"newsletter/app/article/repository"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	// "github.com/gin-gonic/gin"
)

// type ArticleModule struct {
// 	controller ArticleController
// }

// func (module *ArticleModule) Configure(r *gin.Engine) {
// 	article := r.Group("/article")
// 	{
// 		article.GET("", module.controller.FindByPage)
// 		article.POST("", module.controller.Create)
// 		article.GET(":slug", module.controller.FindBySlug)
// 	}
// }

// func NewArticleModule(controller ArticleController) ArticleModule {
// 	return ArticleModule{controller}
// }

func register(r *gin.Engine, controller *ArticleController) {
	article := r.Group("/article")
	{
		article.GET("", controller.FindByPage)
		article.POST("", controller.Create)
		article.GET(":slug", controller.FindBySlug)
	}
}

var ArticleModule = fx.Options(
	fx.Provide(repository.NewInMemoryArticleRepository),
	fx.Provide(NewArticleController),
	fx.Invoke(register),
)
