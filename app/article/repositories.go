package article

import (
	"newsletter/app/common"
)

/*********************/
/* ArticleRepository */
/*********************/

type IArticleRepository interface {
	FindByPage(page int) (int, *[]common.Article, error)
	FindBySlug(slug string) (*common.Article, error)
	Create(article *common.Article) error
}

type InMemoryArticleRepository struct {
	Articles []common.Article
}

func (articleRepository *InMemoryArticleRepository) FindByPage(page int) (int, *[]common.Article, error) {
	articleCount := len(articleRepository.Articles)

	if page >= 0 {
		min := 5 * page
		max := 5 * (page + 1)

		if min < articleCount {
			if max > articleCount {
				max = articleCount
			}

			articles := articleRepository.Articles[min:max]
			return articleRepository.calculatePageCount(), &articles, nil
		}
	}

	articles := make([]common.Article, 0)
	return articleRepository.calculatePageCount(), &articles, nil
}

func (articleRepository *InMemoryArticleRepository) FindBySlug(slug string) (*common.Article, error) {
	for _, article := range articleRepository.Articles {
		if article.Slug == slug {
			return &article, nil
		}
	}

	return nil, nil
}

func (articleRepository *InMemoryArticleRepository) Create(a *common.Article) error {
	articleRepository.Articles = append(articleRepository.Articles, *a)
	return nil
}

func (articleRepository *InMemoryArticleRepository) calculatePageCount() int {
	return (len(articleRepository.Articles) + 5 - 1) / 5
}

func NewInMemoryArticleRepository() IArticleRepository {
	return &InMemoryArticleRepository{[]common.Article{}}
}
