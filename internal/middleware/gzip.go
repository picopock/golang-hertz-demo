package middleware

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/gzip"
)

func registerGzip(h *server.Hertz) {
	h.Use(gzip.Gzip(gzip.DefaultCompression))
}
