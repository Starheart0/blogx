package core

import (
	"blogx_server/global"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func EsConnect() *elastic.Client {
	es := global.Config.ES
	if es.Addr == "" {
		return nil
	}
	client, err := elastic.NewClient(
		elastic.SetURL(es.Url()),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(es.Username, es.Password),
	)
	if err != nil {
		logrus.Panicf("es connection error %s", err)
		return nil
	}
	logrus.Infof("es connection Ac!")
	return client
}
