package user_api

import (
	"blogx_server/commom/res"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/utils/jwts"
	"blogx_server/utils/pwd"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type RegisterEmailRequest struct {
	EmailID   string `json:"emailID" binding:"required"`
	EmailCode string `json:"emailCode" binding:"required"`
	Pwd       string `json:"pwd" binding:"required"`
}

func (UserApi) RegisterEmailView(c *gin.Context) {
	var cr RegisterEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	e, _ := c.Get("email")
	email, ok := e.(string)
	if !ok {
		res.FailWithMsg("email verify error", c)
		return
	}
	uname := base64Captcha.RandText(5, "0123456789")
	hashPwd, _ := pwd.GenerateFromPassword(cr.Pwd)

	var user = models.UserModel{
		Username:       fmt.Sprintf("b_%s", uname),
		Nickname:       "邮箱用户",
		RegisterSource: enum.RegisterEmailSourceType,
		Password:       hashPwd,
		Email:          email,
		Role:           enum.UserRole,
	}
	err = global.DB.Create(&user).Error
	if err != nil {
		res.FailWithMsg("email register error", c)
		logrus.Errorf("create user error")
		return
	}
	token, err := jwts.GetToken(jwts.Claims{
		UserID:   user.ID,
		UserName: user.Username,
		Role:     user.Role,
	})
	if err != nil {
		res.FailWithMsg("email login error", c)
		return
	}
	res.OkWithData(token, c)
}
