package main

import (
	"fmt"
	"log"

	"ziglu/models"
	"ziglu/pkg/setting"
	"ziglu/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

func init() {
	setting.Setup()
	models.CreateConnection()
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	host := fmt.Sprintf(":%s", setting.ServerSetting.HttpPort)

	// server := &http.Server{
	// 	Addr:           host,
	// 	Handler:        routers.InitRouter(),
	// 	MaxHeaderBytes: 1 << 20,
	// }

	rssItems, _ := getRSSArticles("http://feeds.bbci.co.uk/news/uk/rss.xml")
	persistRSSArticles(rssItems)

	log.Printf("[info] Started HTTP server on %s", host)

	// server.ListenAndServe()
}

func getRSSArticles(rssFeedUrl string) ([]*gofeed.Item, error) {
	feed, err := utils.ParseRSSFeed(rssFeedUrl)
	if err != nil {
		log.Fatalf("failed to get RSS feed from %s - %s", rssFeedUrl, err)
	}

	return feed.Items, nil
}

func persistRSSArticles(items []*gofeed.Item) {
	for i := range items {
		author := ""
		if len(items[i].Authors) > 0 && items[i].Authors[0].Name != "" {
			author = items[i].Authors[0].Name
		}

		article := models.Article{
			ArticleID:   items[i].GUID,
			Title:       items[i].Title,
			Description: items[i].Description,
			Link:        items[i].Link,
			Published:   items[i].Published,
			Author:      author,
		}

		err := models.AddArticle(&article)
		if err != nil {
			log.Printf("failed to add article (GUID: %s) to DB - %s\n", items[i].GUID, err)
		}
	}
}
