package middleware

import (
	"blogx_server/commom/res"
	"blogx_server/utils/email_store"
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type EmailVerifyMiddlewareRequest struct {
	EmailID   string `json:"emailID" binding:"required"`
	EmailCode string `json:"emailCode" binding:"required"`
}

func EmailVerifyMiddleware(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		res.FailWithMsg("get header error", c)
		c.Abort()
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	var cr EmailVerifyMiddlewareRequest
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		logrus.Errorf("email check error %s", err)
		res.FailWithMsg("email check error", c)
		c.Abort()
		return
	}
	info, ok := email_store.Verify(cr.EmailID, cr.EmailCode)
	if !ok {
		res.FailWithMsg("email verify error", c)
		c.Abort()
		return
	}
	c.Set("email", info.Email)
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
}
