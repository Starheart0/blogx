package user_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models/enum"
	"blogx_server/utils/jwts"
	"blogx_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"oldPwd" binding:"required"`
	NewPwd string `json:"newPwd" binding:"required"`
}

func (UserApi) UpdatePasswordView(c *gin.Context) {
	cr := middleware.BindJson[UpdatePasswordRequest](c)
	if !global.Config.Site.Login.EmailLogin {
		res.FailWithMsg("site hasn't email function", c)
		return
	}
	claims := jwts.GetClaims(c)
	user, err := claims.GetUser()
	if err != nil {
		res.FailWithMsg("user not exist", c)
		return
	}
	if !(user.RegisterSource == enum.RegisterEmailSourceType || user.Email != "") {
		res.FailWithMsg("only users with email can change password", c)
		return
	}

	if !pwd.CompareHashAndPassword(user.Password, cr.OldPwd) {
		res.FailWithMsg("oldPassword error", c)
		return
	}
	hashPwd, _ := pwd.GenerateFromPassword(cr.NewPwd)
	global.DB.Model(&user).Update("password", hashPwd)
	res.FailWithMsg("successfully change password", c)
}
