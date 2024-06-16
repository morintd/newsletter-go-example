package repository

import (
	"newsletter/app/article/model"
	repositorypkg "newsletter/app/article/repository"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("InMemoryArticleRepository", func() {
	Describe("FindByPage", func() {
		articles := []model.Article{
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
			{
				ID:      "6",
				Title:   "title-5",
				Slug:    "slug-5",
				Content: "content-5",
			},
		}

		repository := repositorypkg.InMemoryArticleRepository{
			Articles: articles,
		}

		It("Should get articles by page with complete page", func() {
			pageCount, actual, err := repository.FindByPage(0)

			if err != nil {
				Fail(err.Error())
			}

			expected := articles[0:5]

			Expect(pageCount).To(Equal(2))
			Expect(actual).To(Equal(&expected))
		})

		It("Should get articles by page with partial page", func() {
			pageCount, actual, err := repository.FindByPage(1)

			if err != nil {
				Fail(err.Error())
			}

			expected := articles[5:6]

			Expect(pageCount).To(Equal(2))
			Expect(actual).To(Equal(&expected))
		})

		It("Should get articles by page with empty page", func() {
			pageCount, actual, err := repository.FindByPage(2)

			if err != nil {
				Fail(err.Error())
			}

			expected := make([]model.Article, 0)

			Expect(pageCount).To(Equal(2))
			Expect(actual).To(Equal(&expected))
		})
	})

	Describe("FindBySlug", func() {
		It("Should return article", func() {
			article := model.Article{
				ID:      "1",
				Title:   "title-1",
				Slug:    "slug-1",
				Content: "content-1",
			}

			repository := repositorypkg.InMemoryArticleRepository{
				Articles: []model.Article{article},
			}

			actual, err := repository.FindBySlug("slug-1")

			if err != nil {
				Fail(err.Error())
			}

			Expect(actual).To(Equal(&article))
		})
	})

	Describe("Create", func() {
		actual := model.Article{
			ID:      "1",
			Title:   "title-1",
			Slug:    "slug-1",
			Content: "content-1",
		}

		repository := repositorypkg.InMemoryArticleRepository{
			Articles: []model.Article{},
		}

		err := repository.Create(&actual)

		if err != nil {
			Fail(err.Error())
		}

		Expect(actual).To(Equal(repository.Articles[0]))
	})
})
