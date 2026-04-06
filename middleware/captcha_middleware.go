package middleware

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CaptchaMiddlewareRequest struct {
	CaptchaID   string `json:"captchaID" binding:"required"`
	CaptchaCode string `json:"captchaCode" binding:"required"`
}

func CaptchaMiddleware(c *gin.Context) {
	if !global.Config.Site.Login.Captcha {
		c.Next()
		return
	}
	body, err := c.GetRawData()
	if err != nil {
		res.FailWithMsg("get header error", c)
		c.Abort()
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	var cr CaptchaMiddlewareRequest
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		logrus.Errorf("image check error %s", err)
		res.FailWithMsg("image check error", c)
		c.Abort()
		return
	}
	if !global.CaptchaStore.Verify(cr.CaptchaID, cr.CaptchaCode, true) {
		res.FailWithMsg("verify code error", c)
		c.Abort()
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
}
