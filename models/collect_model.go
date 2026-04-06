package models

type CollectModel struct {
	Model
	Title       string                    `gorm:"size:32" json:"title"`
	Abstract    string                    `gorm:"size:256" json:"abstract"`
	Cover       string                    `gorm:"size:256" json:"cover"`
	ArticleList []UserArticleCollectModel `gorm:"foreignKey:collectID" json:"-"`
	UserID      uint                      `json:"userID"`
	IsDefault   bool                      `json:"isDefault"`
	UserModel   UserModel                 `gorm:"foreignKey:UserID" json:"-"`
}
