package routers

import (
	"github.com/gin-gonic/gin"

	v1 "news-app/routers/api/v1"
)

// InitRouter Initialises API routing endpoints
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/article/:id", v1.GetArticle)

		apiv1.GET("/article-sources", v1.GetArticleSources)
		apiv1.GET("/article-source/:id", v1.GetArticleSource)
		apiv1.POST("/article-source/:id", v1.SetActiveArticleSource)
	}

	return r
}
