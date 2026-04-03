package middleware

import (
	"blogx_server/commom/res"
	"blogx_server/models/enum"
	"blogx_server/server/redis_service/redis_jwt"
	"blogx_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	blk, ok := redis_jwt.HasTokenBlackByGin(c)
	if ok {
		res.FailWithMsg(blk.Msg(), c)
		c.Abort()
		return
	}
	c.Set("claims", claims)
	return
}

func AdminMiddleware(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	if claims.Role != enum.AdminRole {
		res.FailWithMsg("root error", c)
		c.Abort()
		return
	}
	blk, ok := redis_jwt.HasTokenBlackByGin(c)
	if ok {
		res.FailWithMsg(blk.Msg(), c)
		c.Abort()
		return
	}
	c.Set("claims", claims)
	return
}
