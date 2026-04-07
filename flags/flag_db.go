package flags

import (
	"blogx_server/global"
	"blogx_server/models"

	"github.com/sirupsen/logrus"
)

func FlagDB() {
	err := global.DB.AutoMigrate(
		&models.UserModel{},
		&models.UserConfModel{},
		&models.ArticleModel{},
		&models.CategoryModel{},
		&models.ArticleDiggModel{},
		&models.CollectModel{},
		&models.UserArticleCollectModel{},
		&models.UserTopArticleModel{},
		&models.ImageModel{},
		&models.UserArticleLookHistoryModel{},
		&models.CommentModel{},
		&models.BannerModel{},
		&models.LogModel{},
		&models.UserLoginModel{},
		&models.GlobalNotificationModel{},
		&models.CommentDiggModel{},
	)
	if err != nil {
		logrus.Errorf("database migrate error : %s", err)
		return
	}
	logrus.Infof("database migrate Ac!")
}
