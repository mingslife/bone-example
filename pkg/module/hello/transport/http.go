package transport

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/mingslife/bone"
	"github.com/mingslife/bone-example/pkg/module/hello/endpoint"
)

type HelloHttp struct {
	Router   *bone.Router            `inject:"application.router"`
	Endpoint *endpoint.HelloEndpoint `inject:""`
	Decoder  *HelloDecoder           `inject:""`
}

func (h *HelloHttp) Register() error {
	h.Router.Methods(http.MethodGet).Path("/v1/hello").Handler(bone.NewServer(h.Endpoint.Hello, h.Decoder.Hello, httptransport.EncodeJSONResponse))
	return nil
}

var _ bone.Transport = (*HelloHttp)(nil)
