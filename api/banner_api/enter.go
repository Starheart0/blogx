package banner_api

import (
	"blogx_server/commom"
	"blogx_server/commom/res"
	"blogx_server/global"
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
	var cr BannerCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	err = global.DB.Create(&models.BannerModel{
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
	commom.PageInfo
	Show bool `form:"show"`
}

func (BannerApi) BannerListVier(c *gin.Context) {
	var cr BannerListRequest
	c.ShouldBindQuery(&cr)

	list, count, _ := commom.ListQuery(models.BannerModel{}, commom.Option{
		PageInfo: cr.PageInfo,
	})

	res.OkWithList(list, count, c)
}

func (BannerApi) BannerRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var list []models.BannerModel
	global.DB.Find(&list, "id in ?", cr.IDlist)
	if len(list) > 0 {
		global.DB.Delete(&list)
	}
	res.OkWithMsg(fmt.Sprintf("need delete %d banner, successfully delete %d banner", len(cr.IDlist), len(list)), c)
}
func (BannerApi) BannerUpdateView(c *gin.Context) {
	var id models.IDRequest
	err := c.ShouldBindUri(&id)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var cr BannerCreateRequest
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var model models.BannerModel
	err = global.DB.Take(&model, id.ID).Error
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
