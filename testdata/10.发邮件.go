package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/service/email_service"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()

	//err := email_service.SendRegisterCode("2505909854@qq.com", "5431")
	err := email_service.SendResetPwdCode("2505909854@qq.com", "5431")
	fmt.Println(err)
}
