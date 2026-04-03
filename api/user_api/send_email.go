package user_api

import (
	"blogx_server/commom/res"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/server/email_service"
	"blogx_server/utils/email_store"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type SendEmailRequest struct {
	Type  int8   `json:"type"`
	Email string `json:"email"`
}

type SendEmailResponse struct {
	EmailID string `json:"emailID"`
}

func (UserApi) SendEmailView(c *gin.Context) {
	var cr SendEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	code := base64Captcha.RandText(4, "0123456789")
	id := base64Captcha.RandomId()
	switch cr.Type {
	case 1:
		var user models.UserModel
		err = global.DB.Take(&user, "email = ?", cr.Email).Error
		if err == nil {
			res.FailWithMsg("this email has already been used", c)
			return
		}
		err = email_service.SendRegisterCode(cr.Email, code)
	case 2:
		err = email_service.SendResetPwdCode(cr.Email, code)
	}
	if err != nil {
		logrus.Errorf("email send error %s", err)
		res.FailWithMsg("email send error", c)
		return
	}
	email_store.Set(id, cr.Email, code)
	res.OkWithData(SendEmailResponse{
		EmailID: id,
	}, c)
}
