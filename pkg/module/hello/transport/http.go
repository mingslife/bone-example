package transport

import (
	"github.com/mingslife/bone"
	"github.com/mingslife/bone-example/pkg/module/hello/endpoint"
)

type HelloHttp struct {
	Router   *bone.Router            `inject:"application.router"`
	Endpoint *endpoint.HelloEndpoint `inject:""`
	Decoder  *HelloDecoder           `inject:""`
}

func (h *HelloHttp) Register() error {
	h.Router.RegisterEndpoint("GET", "/v1/hello", h.Endpoint.Hello, h.Decoder.Hello)
	return nil
}

var _ bone.Transport = (*HelloHttp)(nil)
