package router

import (
	_ "hertz_demo/docs"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

func registerSwagger(h *server.Hertz) {
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, swagger.URL("/swagger/doc.json")))
}
