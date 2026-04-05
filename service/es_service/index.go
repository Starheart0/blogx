package es_service

import (
	"blogx_server/global"
	"context"

	"github.com/sirupsen/logrus"
)

func CreateIndexV2(index, mapping string) {
	if ExistsIndex(index) {
		DeleteIndex(index)
	}
	CreateIndex(index, mapping)
}

func CreateIndex(index, mapping string) {
	_, err := global.ESClient.
		CreateIndex(index).
		BodyString(mapping).Do(context.Background())
	if err != nil {
		logrus.Errorf("%s index create error %s", index, err)
		return
	}
	logrus.Infof("%s index create successfully", index)
}

// ExistsIndex 判断索引是否存在
func ExistsIndex(index string) bool {
	exists, _ := global.ESClient.IndexExists(index).Do(context.Background())
	return exists
}

func DeleteIndex(index string) {
	_, err := global.ESClient.
		DeleteIndex(index).Do(context.Background())
	if err != nil {
		logrus.Errorf("%s index delete error %s", index, err)
		return
	}
	logrus.Infof("%s index delete successfully", index)
}
