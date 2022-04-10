package utils

import (
	"context"
	"time"

	"github.com/mmcdole/gofeed"
)

func ParseRSSFeed(rssFeedUrl string) (*gofeed.Feed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	return gofeed.NewParser().ParseURLWithContext(rssFeedUrl, ctx)
}
