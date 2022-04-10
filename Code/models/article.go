package models

import (
	"fmt"

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

// ExistArticleByID checks if an article exists based on ID
func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

func AddArticle(article *Article) error {
	fmt.Printf("%+v\n", *article)

	if err := db.Create(article).Error; err != nil {
		return err
	}

	return nil
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetArticleTotal() (int64, error) {
	var count int64
	if err := db.Model(&Article{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetArticles gets a list of articles based on paging constraints
func GetArticles() ([]*Article, error) {
	var articles []*Article
	err := db.Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

// GetArticle Get a single article based on ID
func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}
