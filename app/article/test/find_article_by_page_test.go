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

var _ = Describe("GET /article?page=:page", Ordered, func() {
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
		{
			ID:      "2",
			Title:   "title-2",
			Slug:    "slug-2",
			Content: "content-2",
		},
		{
			ID:      "3",
			Title:   "title-3",
			Slug:    "slug-3",
			Content: "content-3",
		},
		{
			ID:      "4",
			Title:   "title-4",
			Slug:    "slug-4",
			Content: "content-4",
		},
		{
			ID:      "5",
			Title:   "title-5",
			Slug:    "slug-5",
			Content: "content-5",
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

	It("Should return articles", func() {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/article?page=0", nil)
		r.ServeHTTP(response, request)

		expected := map[string]interface{}{
			"pages":    1,
			"articles": articles[0:5],
		}

		expectedJson, _ := json.Marshal(expected)

		Expect(response.Code).To(Equal(200))
		Expect(response.Body.String()).To(Equal(string(expectedJson)))
	})
})
