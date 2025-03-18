package middleware

import (
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func registerRecovery(h *server.Hertz) {
	h.Use(recovery.Recovery())
}
