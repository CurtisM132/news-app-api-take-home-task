package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	articleService "ziglu/services/article-service"
)

// GetArticle Gets a single article
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

// GetArticles Gets all current articles
func GetArticles(c *gin.Context) {
	articleService := articleService.Article{}

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
