package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/server/redis_service/redis_jwt"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.Redis = core.InitRedis()
	//token, err := jwts.GetToken(jwts.Claims{
	//	UserID: 2,
	//	Role:   1,
	//})
	//fmt.Println(token, err)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjEsInVzZXJOYW1lIjoiIiwicm9sZSI6MSwiZXhwIjoxNzc1MTE1NjAzLCJpc3MiOiJTdGFyaGVhcnQifQ.aZmirI4-WbSy-YT2KEYMjYPlt9Jj0oFr4EuAACvPxw8"
	redis_jwt.TokenBlack(token, redis_jwt.UserBlackType)
	blk, ok := redis_jwt.HasTokenBlack(token)
	fmt.Println(blk, ok)
}
