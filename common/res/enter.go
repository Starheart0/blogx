package res

import (
	"blogx_server/utils/validate"

	"github.com/gin-gonic/gin"
)

type Code int

const (
	SuccessCode     Code = 0
	FaliValidCode   Code = 1001
	FaliServiceCode Code = 1002
)

func (c Code) String() string {
	switch c {
	case SuccessCode:
		return "AC"
	case FaliValidCode:
		return "Valid Fail"
	case FaliServiceCode:
		return "Service Fail"
	}
	return ""
}

var empty = map[string]any{}

func (r Response) Json(c *gin.Context) {
	c.JSON(200, r)
}

type Response struct {
	Code Code   `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Ok(data any, msg string, c *gin.Context) {
	Response{SuccessCode, data, msg}.Json(c)
}

func OkWithData(data any, c *gin.Context) {
	Response{SuccessCode, data, "AC"}.Json(c)
}

func OkWithList(list any, count int, c *gin.Context) {
	Response{SuccessCode, map[string]any{
		"list":  list,
		"count": count,
	}, "AC"}.Json(c)
}

func OkWithMsg(msg string, c *gin.Context) {
	Response{SuccessCode, empty, msg}.Json(c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Response{FaliValidCode, empty, msg}.Json(c)
}

func FailWithData(data any, msg string, c *gin.Context) {
	Response{FaliServiceCode, data, msg}.Json(c)
}

func FailWithCode(code Code, c *gin.Context) {
	Response{code, empty, code.String()}.Json(c)
}

func FailWithError(err error, c *gin.Context) {
	data, msg := validate.ValidateError(err)
	FailWithData(data, msg, c)
}
