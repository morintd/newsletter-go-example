package article

import (
	"net/http"
	"newsletter/app/article/model"
	repositorypkg "newsletter/app/article/repository"
	commonservicepkg "newsletter/app/common/service"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	repository    repositorypkg.IArticleRepository
	slugGenerator commonservicepkg.ISlugGenerator
	idGenerator   commonservicepkg.IIDGenerator
}

func (controller *ArticleController) FindByPage(c *gin.Context) {
	var input struct {
		Page *int `form:"page" binding:"required"`
	}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	pages, articles, err := controller.repository.FindByPage(*input.Page)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{
		"pages":    pages,
		"articles": articles,
	})
}

func (controller *ArticleController) Create(c *gin.Context) {
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	article := model.Article{ID: controller.idGenerator.Generate(), Title: input.Title, Slug: controller.slugGenerator.Generate(input.Title), Content: input.Content}
	err := controller.repository.Create(&article)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(201, article)
}

func (controller *ArticleController) FindBySlug(c *gin.Context) {
	var input struct {
		Slug string `uri:"slug" binding:"required"`
	}

	if err := c.ShouldBindUri(&input); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	article, err := controller.repository.FindBySlug(input.Slug)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	if article != nil {
		c.JSON(http.StatusOK, article)
	} else {
		c.Status(http.StatusBadRequest)
	}

}

func NewArticleController(repository repositorypkg.IArticleRepository, slugGenerator commonservicepkg.ISlugGenerator, idGenerator commonservicepkg.IIDGenerator) *ArticleController {
	return &ArticleController{repository, slugGenerator, idGenerator}
}
