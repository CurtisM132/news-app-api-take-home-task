package article_service

import "ziglu/models"

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string
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

func (a *Article) Count() (int64, error) {
	return models.GetArticleTotal()
}
