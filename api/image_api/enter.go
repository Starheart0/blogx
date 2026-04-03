package image_api

import (
	"blogx_server/commom"
	"blogx_server/commom/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/service/log_server"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ImageApi struct {
}

type ImageListResponse struct {
	models.ImageModel
	WebPath string `json:"webPath"`
}

func (ImageApi) ImageListView(c *gin.Context) {
	cr := middleware.BindJson[commom.PageInfo](c)
	_list, count, _ := commom.ListQuery(models.ImageModel{}, commom.Option{
		PageInfo: cr,
		Likes:    []string{"filename"},
	})
	var list = make([]ImageListResponse, 0)
	for _, model := range _list {
		list = append(list, ImageListResponse{
			ImageModel: model,
			WebPath:    model.WebPath(),
		})
	}
	res.OkWithList(list, count, c)
}

func (ImageApi) ImageRemoveView(c *gin.Context) {
	cr := middleware.BindJson[models.RemoveRequest](c)
	log := log_server.GetLog(c)
	log.ShowRequest()
	log.ShowResponse()

	var List []models.ImageModel
	global.DB.Find(&List, "id in ?", cr.IDlist)

	var successCount, errCount int64
	if len(List) > 0 {
		successCount = global.DB.Delete(&List).RowsAffected
	}
	errCount = int64(len(List)) - successCount
	msg := fmt.Sprintf("image delete successful, successfully delete %d logs, cant delete %d logs", successCount, errCount)
	res.OkWithMsg(msg, c)
}
