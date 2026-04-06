package user_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"

	"github.com/gin-gonic/gin"
)

type UserBaseInfoResponse struct {
	UserID       uint   `json:"userID"`
	CodeAge      int    `json:"codeAge"`
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
	LookCount    int    `json:"lookCount"`
	ArticleCount int    `json:"articleCount"`
	// todo
	FansCount   int    `json:"fansCount"`
	FollowCount int    `json:"followCount"`
	Place       string `json:"place"`
}

func (UserApi) UserBaseInfoView(c *gin.Context) {
	cr := middleware.BindQuery[models.IDRequest](c)
	var user models.UserModel
	err := global.DB.Take(&user, cr.ID).Error
	if err != nil {
		res.FailWithMsg("user not exist", c)
		return
	}

	data := UserBaseInfoResponse{
		UserID:       user.ID,
		CodeAge:      user.CodeAge(),
		Avatar:       user.Avatar,
		Nickname:     user.Nickname,
		LookCount:    1,
		ArticleCount: 1,
		FansCount:    1,
		FollowCount:  1,
		Place:        user.Addr,
	}
	res.OkWithData(data, c)
}
