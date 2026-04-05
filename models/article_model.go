package models

import (
	"blogx_server/models/ctype"
	"blogx_server/models/enum"
	_ "embed"
)

type ArticleModel struct {
	Model
	Title        string             `gorm:"size:32" json:"title"`
	Abstract     string             `gorm:"size:256" json:"abstract"`
	Content      string             `json:"content"`
	CategoryID   uint               `json:"categoryID"`
	TagList      ctype.List         `gorm:"type:longtext;" json:"tagList"`
	Cover        string             `gorm:"size:256" json:"cover"`
	UserID       uint               `json:"userID"`
	UserModel    UserModel          `gorm:"foreignKey:UserID" json:"-"`
	LookCount    int                `json:"lookCount"`
	DiggCount    int                `json:"diggCount"`
	CommentCount int                `json:"commentCount"`
	CollectCount int                `json:"collectCount"`
	OpenComment  bool               `json:"openComment"`
	Status       enum.ArticleStatus `json:"status"`
}

//go:embed mappings/article_mapping.json
var articleMapping string

func (ArticleModel) Mapping() string {
	return articleMapping
}

func (ArticleModel) Index() string {
	return "article_index"
}
