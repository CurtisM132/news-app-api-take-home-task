package cache_service

import (
	"strconv"
	"strings"
)

const CACHE_ARTICLE = "ARTICLE"

type Article struct {
	ID        int
	ArticleID string
}

func (a *Article) GetArticleKey() string {
	return CACHE_ARTICLE + "_" + strconv.Itoa(a.ID)
}

func (a *Article) GetArticlesKey() string {
	keys := []string{
		CACHE_ARTICLE,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}
	if a.ArticleID != "" {
		keys = append(keys, a.ArticleID)
	}

	return strings.Join(keys, "_")
}
