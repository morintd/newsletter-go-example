package core

import (
	"context"
	"fmt"
	"newsletter/app/article"
	"newsletter/app/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	article.ArticleModule,
	common.CommonModule,
	CoreModule,
)

func Bootstrap() *fx.App {
	return fx.New(
		Modules,
		fx.Invoke(runnerHook),
		fx.NopLogger,
	)
}

func runnerHook(lifecycle fx.Lifecycle, r *gin.Engine) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go r.Run()
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
