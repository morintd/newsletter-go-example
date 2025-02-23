package article

import (
	"net/http"
	"newsletter/app/common"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	repository    IArticleRepository
	slugGenerator common.ISlugGenerator
	idGenerator   common.IIDGenerator
}

func (controller *ArticleController) FindByPage(c *gin.Context) {
	var input struct {
		Page *int `form:"page" binding:"required"`
	}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	pages, articles, err := controller.repository.FindByPage(*input.Page)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
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
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	article := common.Article{ID: controller.idGenerator.Generate(), Title: input.Title, Slug: controller.slugGenerator.Generate(input.Title), Content: input.Content}
	err := controller.repository.Create(&article)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	c.JSON(201, article)
}

func (controller *ArticleController) FindBySlug(c *gin.Context) {
	var input struct {
		Slug string `uri:"slug" binding:"required"`
	}

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	article, err := controller.repository.FindBySlug(input.Slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	if article != nil {
		c.JSON(http.StatusOK, article)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{})
	}
}

func (controller *ArticleController) Configure(r *gin.Engine) {
	article := r.Group("/article")
	{
		article.GET("", controller.FindByPage)
		article.POST("", controller.Create)
		article.GET(":slug", controller.FindBySlug)
	}
}

func NewArticleController(repository IArticleRepository, slugGenerator common.ISlugGenerator, idGenerator common.IIDGenerator) *ArticleController {
	return &ArticleController{repository, slugGenerator, idGenerator}
}
