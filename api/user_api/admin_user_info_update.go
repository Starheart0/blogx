package user_api

import (
	"blogx_server/commom/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/utils/mps"

	"github.com/gin-gonic/gin"
)

type AdminUserInfoUpdateRequest struct {
	UserID   uint           `json:"userID" binding:"required"`
	Username *string        `json:"username" s-u:"username"`
	Nickname *string        `json:"nickname" s-u:"nickname"`
	Avatar   *string        `json:"avatar" s-u:"avatar"`
	Abstract *string        `json:"abstract" s-u:"abstract"`
	Role     *enum.RoleType `json:"role" s-u:"role"`
}

func (UserApi) AdminUserInfoUpdateView(c *gin.Context) {
	cr := middleware.BindJson[AdminUserInfoUpdateRequest](c)
	userMap := mps.StructToMap(cr, "s-u")
	var user models.UserModel
	err := global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMsg("user not exist", c)
		return
	}
	err = global.DB.Model(&user).Updates(userMap).Error
	if err != nil {
		res.FailWithMsg("user info modify error", c)
		return
	}
	res.OkWithMsg("user info modify successfully", c)
}
