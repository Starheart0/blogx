package user_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/service/email_service"
	"blogx_server/utils/email_store"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type SendEmailRequest struct {
	Type  int8   `json:"type"` // 1 -> send email   2 -> reset pwd
	Email string `json:"email"`
}

type SendEmailResponse struct {
	EmailID string `json:"emailID"`
}

func (UserApi) SendEmailView(c *gin.Context) {
	cr := middleware.BindJson[SendEmailRequest](c)
	if !global.Config.Site.Login.EmailLogin {
		res.FailWithMsg("site hasn't email function", c)
		return
	}
	code := base64Captcha.RandText(4, "0123456789")
	id := base64Captcha.RandomId()
	var err error = nil
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
		var user models.UserModel
		err = global.DB.Take(&user, "email = ?", cr.Email).Error
		if err != nil {
			res.FailWithMsg("this email not exist", c)
			return
		}
		err = email_service.SendResetPwdCode(cr.Email, code)
	case 3:
		var user models.UserModel
		err = global.DB.Take(&user, "email = ?", cr.Email).Error
		if err == nil {
			res.FailWithMsg("this email has already been used", c)
			return
		}
		err = email_service.SendBindEmailCode(cr.Email, code)
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
