package hello

import (
	"github.com/mingslife/bone"
	"github.com/mingslife/bone-example/pkg/module/hello/transport"
)

type Module struct {
	Http *transport.HelloHttp `inject:""`
}

func (*Module) Name() string {
	return "module.hello"
}

func (*Module) Init() error {
	return nil
}

func (m *Module) Register() error {
	return m.Http.Register()
}

func (*Module) Unregister() error {
	return nil
}

var _ bone.Module = (*Module)(nil)
