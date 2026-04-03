package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/utils/pwd"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()

	p, _ := pwd.GenerateFromPassword("123456")
	fmt.Println(p)
	fmt.Println(pwd.CompareHashAndPassword(p, "123456"))

	var userModel = models.UserModel{
		Username: "Starheart",
		Password: p,
	}
	err := global.DB.Create(&userModel).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
}
