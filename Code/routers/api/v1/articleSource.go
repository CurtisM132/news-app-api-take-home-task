package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	articleSourceService "ziglu/services/article-source-service"
)

// GetArticleSource Gets a single article source
func GetArticleSource(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "request doesn't contain valid id"})
		return
	}

	articleSource := articleSourceService.ArticleSource{ID: id}
	exists, err := articleSource.ExistByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("failed to check for article source - %s", err)})
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "article source not found"})
		return
	}

	article, err := articleSource.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("failed to get article source - %s", err)})
		return
	}

	c.JSON(http.StatusOK, article)
}

// GetArticleSources Gets all available article sources
func GetArticleSources(c *gin.Context) {
	articleSourceService := articleSourceService.ArticleSource{}

	articles, err := articleSourceService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("failed to retrieve article sources - %s", err)})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"sources": articles})
}
