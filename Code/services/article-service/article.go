package article_service

import "news-app/models"

type Article struct {
	ID          int
	ArticleID   string
	Title       string
	Description string
	Link        string
	Published   string
	Author      string
}

func (a *Article) Add() error {
	article := models.Article{
		ArticleID:   a.ArticleID,
		Title:       a.Title,
		Description: a.Description,
		Link:        a.Link,
		Published:   a.Published,
		Author:      a.Author,
	}

	err := models.AddArticle(&article)
	if err != nil {
		return err
	}

	return nil
}

func (a *Article) Get() (*models.Article, error) {
	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (a *Article) GetAll() ([]*models.Article, error) {
	articles, err := models.GetArticles()
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

func (a *Article) ExistByArticleID() (bool, error) {
	return models.ExistArticleByArticleID(a.ArticleID)
}

func (a *Article) Count() (int64, error) {
	return models.GetArticleTotal()
}
