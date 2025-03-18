package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	hello "hertz_demo/gen/model/hello"
)

type WorldService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewWorldService(Context context.Context, RequestContext *app.RequestContext) *WorldService {
	return &WorldService{RequestContext: RequestContext, Context: Context}
}

func (h *WorldService) Run(req *hello.WorldReq) (resp *hello.WorldResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
