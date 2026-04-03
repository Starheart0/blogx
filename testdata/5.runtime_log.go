package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/service/log_server"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()

	log := log_server.NewRuntimeLog("abc", log_server.RuntimeDateHour)
	log.SetItem("1", 11)
	log.Save()
	log.SetItem("2", 12)
	log.Save()
}
