package models

import (
	"news-app/internal/setting"

	"gorm.io/gorm"
)

type ArticleSource struct {
	Model

	URL string `json:"url"`
}

var activeNewsSource *ArticleSource

// ExistSourceByID Returns if the record exists in the database
func ExistSourceByID(id int) (bool, error) {
	var source ArticleSource
	err := db.Table(setting.DatabaseSetting.ArticleSourceTable).Select("id").Where("id = ?", id).First(&source).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if source.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetArticleSources gets a list of all the possible article RSS sources
func GetArticleSources() ([]*ArticleSource, error) {
	var articleSources []*ArticleSource
	err := db.Table(setting.DatabaseSetting.ArticleSourceTable).Find(&articleSources).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articleSources, nil
}

// GetArticleSource Gets a single article source based on the database record ID
func GetArticleSource(id int) (*ArticleSource, error) {
	var articleSource ArticleSource
	err := db.Table(setting.DatabaseSetting.ArticleSourceTable).Where("id = ?", id).First(&articleSource).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &articleSource, nil
}

// SetActiveArticleSource Sets the active article source URL
func SetActiveArticleSource(id int) error {
	source, err := GetArticleSource(id)
	if err != nil {
		return err
	}

	// TODO: Add validation
	activeNewsSource = source
	return nil
}

// GetActiveArticleSource Gets the active article source URL
func GetActiveArticleSource() *ArticleSource {
	return activeNewsSource
}
