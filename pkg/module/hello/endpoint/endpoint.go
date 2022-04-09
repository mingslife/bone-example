package endpoint

import (
	"context"

	"github.com/mingslife/bone"
	"github.com/mingslife/bone-example/pkg/module/hello/model"
	"github.com/mingslife/bone-example/pkg/module/hello/service"
)

type HelloEndpoint struct {
	Service *service.HelloService `inject:""`
}

func (e *HelloEndpoint) Hello(ctx context.Context, req any) (rsp any, err error) {
	msg := e.Service.Hello(ctx, req.(*model.HelloReq).Name)
	return &model.HelloRsp{Msg: msg}, nil
}

var _ bone.Endpoint = (*HelloEndpoint)(nil)
