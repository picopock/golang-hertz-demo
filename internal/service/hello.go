package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	hello "hertz_demo/gen/model/hello"
)

type HelloService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHelloService(Context context.Context, RequestContext *app.RequestContext) *HelloService {
	return &HelloService{RequestContext: RequestContext, Context: Context}
}

func (h *HelloService) Run(req *hello.HelloReq) (resp *hello.HelloResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
