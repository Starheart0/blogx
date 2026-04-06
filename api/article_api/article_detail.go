package article_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/service/redis_service/redis_article"
	"blogx_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

type ArticleDetailResponse struct {
	models.ArticleModel
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	UserAvatar string `json:"userAvatar"`
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	cr := middleware.BindUri[models.IDRequest](c)

	// 未登录的用户，只能看到发布成功的文章

	// 登录用户，能看到自己的所有文章

	// 管理员，能看到全部的文章
	var article models.ArticleModel
	err := global.DB.Preload("UserModel").Take(&article, cr.ID).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		// 没登录的
		if article.Status != enum.ArticleStatusPublished {
			res.FailWithMsg("文章不存在", c)
			return
		}
	}
	switch claims.Role {
	case enum.UserRole:
		if claims.UserID != article.UserID {
			// 登录的人看到不是自己的
			if article.Status != enum.ArticleStatusPublished {
				res.FailWithMsg("文章不存在", c)
				return
			}
		}
	}
	collectCount := redis_article.GetCacheCollect(article.ID)
	diggCount := redis_article.GetCacheDigg(article.ID)
	lookCount := redis_article.GetCacheLook(article.ID)

	article.CollectCount = article.CollectCount + collectCount
	article.DiggCount = article.DiggCount + diggCount
	article.LookCount = article.LookCount + lookCount
	res.OkWithData(ArticleDetailResponse{
		ArticleModel: article,
		Username:     article.UserModel.Username,
		Nickname:     article.UserModel.Nickname,
		UserAvatar:   article.UserModel.Avatar,
	}, c)
}
