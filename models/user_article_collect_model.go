package models

import (
	"blogx_server/service/redis_service/redis_article"

	"gorm.io/gorm"
)

type UserArticleCollectModel struct {
	Model
	UserID       uint         `gorm:"uniqueIndex:idx_name" json:"userID"`
	UserModel    UserModel    `gorm:"foreignKey:UserID" json:"-"`
	ArticleID    uint         `gorm:"uniqueIndex:idx_name" json:"articleID"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID" json:"-"`
	CollectID    uint         `gorm:"uniqueIndex:idx_name" json:"collectID"`
	CollectModel CollectModel `gorm:"foreignKey:CollectID" json:"collectModel"`
}

func (u *UserArticleCollectModel) BeforeDelete(tx *gorm.DB) (err error) {
	redis_article.SetCacheCollect(u.ArticleID, false)
	return
}
