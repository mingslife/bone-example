package transport

import (
	"context"
	"net/http"

	"github.com/mingslife/bone-example/pkg/module/hello/model"
)

type HelloDecoder struct{}

func (*HelloDecoder) Hello(_ context.Context, r *http.Request) (any, error) {
	req := &model.HelloReq{}
	req.Name = r.URL.Query().Get("name")
	return req, nil
}
