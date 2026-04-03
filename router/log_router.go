package router

import (
	"blogx_server/api"
	"blogx_server/api/log_api"
	"blogx_server/middleware"
	"blogx_server/models"

	"github.com/gin-gonic/gin"
)

func LogRouter(rr *gin.RouterGroup) {
	app := api.App.LogApi
	r := rr.Group("").Use(middleware.AdminMiddleware)
	r.GET("logs", middleware.BindQueryMiddleware[log_api.LogListRequest], app.LogListView)
	r.GET("logs/:id", middleware.BindUriMiddleware[models.IDRequest], app.LogReadView)
	r.DELETE("logs", middleware.BindJsonMiddleware[models.RemoveRequest], app.LogRemoveView)
}
