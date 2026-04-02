package router

import (
	"blogx_server/global"
	"blogx_server/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(global.Config.System.GinMode)
	r := gin.Default()

	r.Static("/uploads", "uploads")

	nr := r.Group("/api")
	nr.Use(middleware.LogMiddleware)
	SiteRouter(nr)
	LogRouter(nr)
	ImageRouter(nr)
	BannerRouter(nr)
	addr := global.Config.System.Addr()
	r.Run(addr)
}
