package router

import (
	"blogx_server/api"
	"blogx_server/middleware"

	"github.com/gin-gonic/gin"
)

func BannerRouter(r *gin.RouterGroup) {
	app := api.App.BannerApi
	r.GET("banner", app.BannerListVier)
	r.PUT("banner/:id", middleware.AdminMiddleware, app.BannerUpdateView)
	r.POST("banner", middleware.AdminMiddleware, app.BannerCreateVier)
	r.DELETE("banner", middleware.AdminMiddleware, app.BannerRemoveView)
}
