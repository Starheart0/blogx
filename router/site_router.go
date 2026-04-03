package router

import (
	"blogx_server/api"
	"blogx_server/api/site_api"
	"blogx_server/middleware"

	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.SiteApi
	r.GET("site/qq_url", app.SiteInfoQQView)
	r.GET("site/:name", middleware.BindJsonMiddleware[site_api.SiteInfoRequest], app.SiteInfoView)
	r.PUT("site/:name", middleware.AdminMiddleware, middleware.BindUriMiddleware[site_api.SiteInfoRequest], app.SiteUpdateView)
}
