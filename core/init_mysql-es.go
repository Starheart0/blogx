package core

import (
	"blogx_server/global"
	river "blogx_server/service/river_service"

	"github.com/sirupsen/logrus"
)

func InitMysqlES() {
	if !global.Config.River.Enable {
		logrus.Infof("without conf es, close mysql sync operation")
		return
	}
	r, err := river.NewRiver()
	if err != nil {
		logrus.Fatal(err)
	}
	go r.Run()
}
