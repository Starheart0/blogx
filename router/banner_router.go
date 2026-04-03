package router

import (
	"blogx_server/api"
	"blogx_server/api/banner_api"
	"blogx_server/middleware"
	"blogx_server/models"

	"github.com/gin-gonic/gin"
)

func BannerRouter(r *gin.RouterGroup) {
	app := api.App.BannerApi
	r.GET("banner", middleware.BindQueryMiddleware[banner_api.BannerListRequest], app.BannerListView)
	r.PUT("banner/:id", middleware.AdminMiddleware, middleware.BindUriMiddleware[models.IDRequest], middleware.BindJsonMiddleware[banner_api.BannerCreateRequest], app.BannerUpdateView)
	r.POST("banner", middleware.AdminMiddleware, app.BannerCreateVier)
	r.DELETE("banner", middleware.AdminMiddleware, app.BannerRemoveView)
}
