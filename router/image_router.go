package router

import (
	"blogx_server/api"
	"blogx_server/commom"
	"blogx_server/middleware"
	"blogx_server/models"

	"github.com/gin-gonic/gin"
)

func ImageRouter(r *gin.RouterGroup) {
	app := api.App.ImageApi
	r.POST("images", middleware.AuthMiddleware, app.ImageUploadView)
	r.GET("images", middleware.AuthMiddleware, middleware.BindJsonMiddleware[commom.PageInfo], app.ImageListView)
	r.DELETE("images", middleware.AuthMiddleware, middleware.BindJsonMiddleware[models.RemoveRequest], app.ImageRemoveView)
}
