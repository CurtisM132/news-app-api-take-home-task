package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	articleService "ziglu/services/article-service"
)

// @Summary Get a single article
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles/{id} [get]
func GetArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "request doesn't contain valid id"})
		return
	}

	articleService := articleService.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("failed to check for article - %s", err)})
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "article not found"})
		return
	}

	article, err := articleService.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("failed to get article - %s", err)})
		return
	}

	c.JSON(http.StatusOK, article)
}

// @Summary Get multiple articles
// @Produce  json
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles [get]
func GetArticles(c *gin.Context) {
	articleService := articleService.Article{
		// TagID: tagId,
	}

	total, err := articleService.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("failed to count articles - %s", err)})
		return
	}

	articles, err := articleService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("failed to retrieve articles - %s", err)})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"articles": articles, "total": total})
}
