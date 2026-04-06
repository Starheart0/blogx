package user_api

import (
	"blogx_server/common"
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/utils/jwts"
	"time"

	"github.com/gin-gonic/gin"
)

type UserLoginListRequest struct {
	common.PageInfo
	UserID    uint   `form:"userID"`
	Ip        string `form:"ip"`
	Addr      string `form:"addr"`
	StartTime string `form:"startTime"` // format: 2006-01-02 15:04:05
	EndTime   string `form:"endTime"`
	Type      int8   `form:"type" binding:"required,oneof=1 2"` // 1 -> only self   2 -> check all
}

type UserLoginListResponse struct {
	models.UserLoginModel
	UserNickname string `json:"userNickname,omitempty"`
	UserAvatar   string `json:"userAvatar,omitempty"`
}

func (UserApi) UserLoginListView(c *gin.Context) {
	cr := middleware.BindQuery[UserLoginListRequest](c)
	claims := jwts.GetClaims(c)
	if cr.Type == 1 {
		cr.UserID = claims.UserID
	}

	var query = global.DB.Where("")
	if cr.StartTime != "" {
		_, err := time.Parse("2006-01-02 15:04:05", cr.StartTime)
		if err != nil {
			res.FailWithMsg("startTime format error", c)
			return
		}
		query.Where("created_at >= ?", cr.StartTime)
	}
	if cr.EndTime != "" {
		_, err := time.Parse("2006-01-02 15:04:05", cr.EndTime)
		if err != nil {
			res.FailWithMsg("endTime format error", c)
			return
		}
		query.Where("created_at >= ?", cr.EndTime)
	}
	var preloads []string
	if cr.Type == 2 {
		preloads = []string{"UserModel"}
	}

	_list, count, _ := common.ListQuery(models.UserLoginModel{
		UserID: cr.UserID,
		IP:     cr.Ip,
		Addr:   cr.Addr,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Where:    query,
		Preloads: preloads,
	})
	var list = make([]UserLoginListResponse, 0)
	for _, model := range _list {
		list = append(list, UserLoginListResponse{
			UserLoginModel: model,
			UserNickname:   model.UserModel.Nickname,
			UserAvatar:     model.UserModel.Avatar,
		})
	}
	res.OkWithList(list, count, c)
}
