package repository

import "newsletter/app/article/model"

type IArticleRepository interface {
	FindByPage(page int) (int, *[]model.Article, error)
	FindBySlug(slug string) (*model.Article, error)
	Create(article *model.Article) error
}
