package user_api

import (
	"blogx_server/commom/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

type ResetPasswordRequest struct {
	Pwd string `json:"pwd" binding:"required"`
}

func (UserApi) ResetPasswordView(c *gin.Context) {
	cr := middleware.BindJson[ResetPasswordRequest](c)
	if !global.Config.Site.Login.EmailLogin {
		res.FailWithMsg("site hasn't email function", c)
		return
	}
	e, _ := c.Get("email")
	email, _ := e.(string)
	var user models.UserModel
	err := global.DB.Take(&user, "email = ?", email).Error
	if err != nil {
		res.FailWithMsg("user not exist", c)
		return
	}
	if !(user.RegisterSource == enum.RegisterEmailSourceType) {
		res.FailWithMsg("Users who have not registered through the email cannot reset their passwords", c)
		return
	}
	hashPwd, _ := pwd.GenerateFromPassword(cr.Pwd)
	global.DB.Model(&user).Update("password", hashPwd)
	res.OkWithMsg("successfully reset password", c)
}
