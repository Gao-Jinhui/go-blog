package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model
	TagID      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id=?", id).First(&article)
	return article.ID > 0
}

func GetArticleTotal(filter interface{}) int {
	var count int
	db.Model(&Article{}).Find(filter).Count(&count)
	return count
}

func GetArticle(id int) (article Article) {
	db.Preload("Tag").Where("id=?", id).First(&article)
	return
}

func GetArticles(pageNum int, pageSize int, filter interface{}) []Article {
	articles := []Article{}
	db.Preload("Tag").Where(filter).Offset(pageNum).Limit(pageSize).Find(&articles)
	return articles
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id=?", id).Updates(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["ID"].(int),
		Title:     data["Title"].(string),
		Desc:      data["Desc"].(string),
		Content:   data["Content"].(string),
		CreatedBy: data["CreatedBy"].(string),
		State:     data["State"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})
	return true
}
