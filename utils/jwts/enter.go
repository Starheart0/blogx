package jwts

import (
	"blogx_server/global"
	"blogx_server/models/enum"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID   uint          `json:"userID"`
	UserName string        `json:"userName"`
	Role     enum.RoleType `json:"role"`
}

type MyClaims struct {
	Claims
	jwt.StandardClaims
}

// get token
func GetToken(claims Claims) (string, error) {
	cla := MyClaims{
		Claims: claims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(global.Config.Jwt.Expire) * time.Hour * 100).Unix(), // 过期时间
			Issuer:    global.Config.Jwt.Issuer,                                                         // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	return token.SignedString([]byte(global.Config.Jwt.Secret)) // 进行签名生成对应的token
}

// parse token
func ParseToken(tokenString string) (*MyClaims, error) {
	if tokenString == "" {
		return nil, errors.New("need login")
	}
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.Secret), nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return nil, errors.New("token timeout")
		}
		if strings.Contains(err.Error(), "invalid") {
			return nil, errors.New("token invalid")
		}
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func ParseTokenByGin(c *gin.Context) (*MyClaims, error) {
	token := c.GetHeader("token")
	if token == "" {
		token = c.Query("token")
	}
	return ParseToken(token)
}

func GetCliams(c *gin.Context) (claims *MyClaims) {
	_claims, ok := c.Get("claims")
	if !ok {
		return
	}
	claims, ok = _claims.(*MyClaims)
	if !ok {
		return
	}
	return
}
