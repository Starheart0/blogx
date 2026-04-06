package banner_api

import (
	"blogx_server/common"
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type BannerApi struct {
}
type BannerCreateRequest struct {
	Cover string `json:"cover" binding:"required"`
	Href  string `json:"href"`
	Show  bool   `json:"show"`
}

func (BannerApi) BannerCreateVier(c *gin.Context) {
	cr := middleware.BindJson[BannerCreateRequest](c)
	err := global.DB.Create(&models.BannerModel{
		Cover: cr.Cover,
		Href:  cr.Href,
		Show:  true,
	}).Error

	if err != nil {
		res.FailWithError(err, c)
		return
	}
	res.OkWithMsg("banner create AC", c)
}

type BannerListRequest struct {
	common.PageInfo
	Show bool `form:"show"`
}

func (BannerApi) BannerListView(c *gin.Context) {
	cr := middleware.BindQuery[BannerListRequest](c)

	list, count, _ := common.ListQuery(models.BannerModel{}, common.Options{
		PageInfo: cr.PageInfo,
	})

	res.OkWithList(list, count, c)
}

func (BannerApi) BannerRemoveView(c *gin.Context) {
	cr := middleware.BindJson[models.RemoveRequest](c)
	var list []models.BannerModel
	global.DB.Find(&list, "id in ?", cr.IDList)
	if len(list) > 0 {
		global.DB.Delete(&list)
	}
	res.OkWithMsg(fmt.Sprintf("need delete %d banner, successfully delete %d banner", len(cr.IDList), len(list)), c)
}
func (BannerApi) BannerUpdateView(c *gin.Context) {
	id := middleware.BindUri[models.IDRequest](c)
	cr := middleware.BindJson[BannerCreateRequest](c)

	var model models.BannerModel
	err := global.DB.Take(&model, id.ID).Error
	if err != nil {
		res.FailWithMsg("not exist such banner", c)
		return
	}
	err = global.DB.Model(&model).Updates(map[string]any{
		"cover": cr.Cover,
		"href":  cr.Href,
		"show":  cr.Show,
	}).Error
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	res.OkWithMsg("banner update AC", c)
}
