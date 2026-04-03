package user_api

import (
	"blogx_server/commom/res"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/utils/jwts"
	"blogx_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

type PwdLoginRequest struct {
	Val      string `json:"val" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (UserApi) PwdLoginApi(c *gin.Context) {
	var cr PwdLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	if !global.Config.Site.Login.UsernamePwdLogin {
		res.FailWithMsg("site cant user password login", c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, "(username = ? or email = ?) and password <> ''",
		cr.Val, cr.Val).Error
	if err != nil {
		res.FailWithMsg("username error", c)
		return
	}
	if !pwd.CompareHashAndPassword(user.Password, cr.Password) {
		res.FailWithMsg("username/password error", c)
		return
	}
	token, _ := jwts.GetToken(jwts.Claims{
		UserID:   user.ID,
		UserName: user.Username,
		Role:     user.Role,
	})
	res.OkWithData(token, c)
}
