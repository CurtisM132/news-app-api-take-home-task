package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"ziglu/models"
	"ziglu/pkg/setting"
	"ziglu/pkg/utils"
	"ziglu/routers"
	articleService "ziglu/services/article-service"

	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

func init() {
	setting.Setup()
	models.CreateConnection()
}

func main() {
	go keepArticlesUpToDate()
	startHTTPServer()
}

func startHTTPServer() {
	gin.SetMode(gin.ReleaseMode)

	host := fmt.Sprintf(":%s", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:           host,
		Handler:        routers.InitRouter(),
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[INFO] Started HTTP server on %s\n", host)

	server.ListenAndServe()
}

func keepArticlesUpToDate() {
	rssFeedUrl := "http://feeds.bbci.co.uk/news/uk/rss.xml"

	for {
		rssItems, err := getRSSArticles(rssFeedUrl)
		if err != nil {
			log.Printf("[ERROR] Failed to get articles from RSS feed (%s) - %s\n", rssFeedUrl, err)
		}
		persistRSSArticles(rssItems)

		time.Sleep(5 * time.Minute)
	}
}

func getRSSArticles(rssFeedUrl string) ([]*gofeed.Item, error) {
	feed, err := utils.ParseRSSFeed(rssFeedUrl)
	if err != nil {
		return nil, err
	}

	return feed.Items, nil
}

func persistRSSArticles(items []*gofeed.Item) {
	for i := range items {
		article := articleService.Article{ArticleID: items[i].GUID}

		exists, err := article.ExistByArticleID()
		if err == nil && exists {
			// News articles are defaultly sorted newest to oldest (in an RSS feed) so if we find an article that already
			// exists in our storage then we can safely assume the rest of the articles below this point do too
			break
		}

		author := ""
		if len(items[i].Authors) > 0 && items[i].Authors[0].Name != "" {
			author = items[i].Authors[0].Name
		}
		article = articleService.Article{
			ArticleID:   items[i].GUID,
			Title:       items[i].Title,
			Description: items[i].Description,
			Link:        items[i].Link,
			Published:   items[i].Published,
			Author:      author,
		}

		err = article.Add()
		if err != nil {
			log.Printf("[ERROR] Failed to add article (GUID: %s) - %s\n", items[i].GUID, err)
		}
	}
}
