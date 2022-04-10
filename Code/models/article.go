package models

import (
	"news-app/internal/setting"

	"gorm.io/gorm"
)

type Article struct {
	Model

	ArticleID   string `json:"article_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Published   string `json:"published"`
	Author      string `json:"author"`
}

// ExistArticleByID Returns if the record exists in the database
func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Table(setting.DatabaseSetting.ArticleTable).Select("id").Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

// ExistArticleByArticleID Returns if the record exists in the database using the articles id
func ExistArticleByArticleID(id string) (bool, error) {
	var article Article
	err := db.Table(setting.DatabaseSetting.ArticleTable).Select("id").Where("article_id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

// AddArticle Adds an article to the database
func AddArticle(article *Article) error {
	err := db.Table(setting.DatabaseSetting.ArticleTable).Create(article).Error
	if err != nil {
		return err
	}

	return nil
}

// GetArticleTotal gets the total number of articles
func GetArticleTotal() (int64, error) {
	var count int64
	err := db.Table(setting.DatabaseSetting.ArticleTable).Model(&Article{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetArticles gets a list of articles based on paging constraints
func GetArticles() ([]*Article, error) {
	var articles []*Article
	err := db.Table(setting.DatabaseSetting.ArticleTable).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

// GetArticle Get a single article based on ID
func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Table(setting.DatabaseSetting.ArticleTable).Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}
