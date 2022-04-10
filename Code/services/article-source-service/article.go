package article_service

import "ziglu/models"

type ArticleSource struct {
	ID  int
	URL string
}

func (a *ArticleSource) Get() (*models.ArticleSource, error) {
	articleSource, err := models.GetArticleSource(a.ID)
	if err != nil {
		return nil, err
	}

	return articleSource, nil
}

func (a *ArticleSource) GetAll() ([]*models.ArticleSource, error) {
	articleSources, err := models.GetArticleSources()
	if err != nil {
		return nil, err
	}

	return articleSources, nil
}

func (a *ArticleSource) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}
