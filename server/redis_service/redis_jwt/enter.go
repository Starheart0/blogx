package redis_jwt

import (
	"blogx_server/global"
	"blogx_server/utils/jwts"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type BlackType int8

const (
	UserBlackType   BlackType = 1
	AdminBlackType  BlackType = 2
	DeviceBlackType BlackType = 3 // preemption
)

func (b BlackType) String() string {
	return fmt.Sprintf("%d", b)
}

func (b BlackType) Msg() string {
	switch b {
	case UserBlackType:
		return "cancelled"
	case AdminBlackType:
		return "banned"
	case DeviceBlackType:
		return "device logout"
	}
	return "cancelled"
}

func ParseBlackType(value string) BlackType {
	switch value {
	case "1":
		return UserBlackType
	case "2":
		return AdminBlackType
	case "3":
		return DeviceBlackType
	}
	return UserBlackType
}

func TokenBlack(token string, value BlackType) {
	key := fmt.Sprintf("token_black_%s", token)
	claims, err := jwts.ParseToken(token)
	if err != nil || claims == nil {
		logrus.Errorf("token analyze error %s", err)
		return
	}
	second := claims.ExpiresAt - time.Now().Unix()
	_, err = global.Redis.Set(key, value.String(), time.Duration(second)*time.Second).Result()
	if err != nil {
		logrus.Errorf("redis add black error %s", err)
		return
	}
}

func HasTokenBlack(token string) (blk BlackType, ok bool) {
	key := fmt.Sprintf("token_black_%s", token)
	value, err := global.Redis.Get(key).Result()
	if err != nil {
		return
	}
	blk = ParseBlackType(value)
	ok = true
	return
}

func HasTokenBlackByGin(c *gin.Context) (blk BlackType, ok bool) {
	token := c.GetHeader("token")
	if token == "" {
		token = c.Query("token")
	}
	return HasTokenBlack(token)
}
