package user_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/utils/jwts"
	"blogx_server/utils/mps"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type UserInfoUpdateRequest struct {
	Username    *string   `json:"username" s-u:"username"`
	Nickname    *string   `json:"nickname" s-u:"nickname"`
	Avatar      *string   `json:"avatar" s-u:"avatar"`
	Abstract    *string   `json:"abstract" s-u:"abstract"`
	LikeTags    *[]string `json:"likeTags" s-u-c:"like_tags"`
	OpenCollect *bool     `json:"openCollect" s-u-c:"open_collect"` //open user collect
	OpenFollow  *bool     `json:"openFollow" s-u-c:"open_follow"`
	OpenFans    *bool     `json:"openFans" s-u-c:"open_fans"`
	HomeStyleID *uint     `json:"homeStyleID" s-u-c:"home_style_id"`
}

func (UserApi) UserInfoUpdateView(c *gin.Context) {
	cr := middleware.BindJson[UserInfoUpdateRequest](c)
	userMap := mps.StructToMap(cr, "s-u")
	userConfMap := mps.StructToMap(cr, "s-u-c")

	claims := jwts.GetClaims(c)

	if len(userMap) > 0 {
		var userModel models.UserModel
		err := global.DB.Preload("UserConfModel").Take(&userModel, claims.UserID).Error
		if err != nil {
			res.FailWithMsg("user not exist1", c)
			return
		}
		if cr.Username != nil {
			var userCount int64
			global.DB.Model(models.UserModel{}).Where("username = ? and id <> ?", *cr.Username, claims.UserID).Count(&userCount)
			if userCount > 0 {
				res.FailWithMsg("this username has been used", c)
				return
			}
			if *cr.Username != claims.UserName {
				var uud = userModel.UserConfModel.UpdateUsernameDate
				if uud != nil {
					if time.Now().Sub(*uud).Hours() < 720 {
						res.FailWithMsg("only can modify once per month", c)
						return
					}
				}
				userConfMap["update_username_date"] = time.Now()
			} else {
				delete(userMap, "username")
			}
		}

		fmt.Println(userMap)
		err = global.DB.Model(&userModel).Updates(userMap).Error
		if err != nil {
			res.FailWithMsg("user info modify error", c)
			return
		}
		res.OkWithMsg("user info modify successfully", c)
	}
	if len(userConfMap) > 0 {
		var userConfModel models.UserConfModel
		err := global.DB.Take(&userConfModel, "user_id = ?", claims.UserID).Error
		if err != nil {
			res.FailWithMsg("user conf info not exist", c)
			return
		}
		fmt.Println(userConfMap)
		err = global.DB.Model(&userConfModel).Updates(userConfMap).Error
		if err != nil {
			res.FailWithMsg("user conf info modify error", c)
			return
		}
		res.OkWithMsg("user conf info modify successfully", c)
	}
}
