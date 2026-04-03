package middleware

import (
	"blogx_server/commom/res"

	"github.com/gin-gonic/gin"
)

func BindJsonMiddleware[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	c.Set("json", cr)
	return
}

func BindJson[T any](c *gin.Context) T {
	return c.MustGet("json").(T)
}

func BindQueryMiddleware[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	c.Set("query", cr)
	return
}

func BindQuery[T any](c *gin.Context) T {
	return c.MustGet("query").(T)
}

func BindUriMiddleware[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	c.Set("uri", cr)
	return
}

func BindUri[T any](c *gin.Context) T {
	return c.MustGet("uri").(T)
}
