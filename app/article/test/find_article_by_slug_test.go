package article

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"newsletter/app/article"
	"newsletter/app/common"
	"newsletter/app/core"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
)

var _ = Describe("GET /article/:slug", Ordered, func() {
	var r *gin.Engine
	var articleRepository article.IArticleRepository
	var app *fx.App

	articles := []common.Article{
		{
			ID:      "1",
			Title:   "title-1",
			Slug:    "slug-1",
			Content: "content-1",
		},
	}

	BeforeAll(func() {
		app = fx.New(core.Modules, fx.Invoke(func(lifecycle fx.Lifecycle, engine *gin.Engine, a article.IArticleRepository) {
			lifecycle.Append(
				fx.Hook{
					OnStart: func(context.Context) error {
						r = engine
						articleRepository = a
						return nil
					},
					OnStop: func(context.Context) error {
						fmt.Println("Stopping application")
						return nil
					},
				},
			)
		}))

		app.Start(context.Background())

		for _, article := range articles {
			articleRepository.Create(&article)
		}
	})

	AfterAll(func() {
		app.Done()
	})

	It("Should return article", func() {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/article/"+articles[0].Slug, nil)
		r.ServeHTTP(response, request)

		expected := articles[0]

		expectedJson, _ := json.Marshal(expected)

		Expect(response.Code).To(Equal(200))
		Expect(response.Body.String()).To(Equal(string(expectedJson)))
	})

	It("Should return error if article with slug doesn't exist", func() {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/article/wrong-slug", nil)
		r.ServeHTTP(response, request)

		Expect(response.Code).To(Equal(400))
		Expect(response.Body.String()).To(Equal("{}"))
	})
})
