package user_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

func (UserApi) BindEmailView(c *gin.Context) {
	if !global.Config.Site.Login.EmailLogin {
		res.FailWithMsg("site hasn't email function", c)
		return
	}
	e, _ := c.Get("email")
	email, ok := e.(string)
	if !ok {
		res.FailWithMsg("email verify error", c)
		return
	}
	user, err := jwts.GetClaims(c).GetUser()
	if err != nil {
		res.FailWithMsg("user not exist", c)
		return
	}
	err = global.DB.Model(&user).Update("email", email).Error
	if err != nil {
		res.FailWithMsg("email update error", c)
		return
	}
	res.OkWithMsg("email bind successful", c)
}
