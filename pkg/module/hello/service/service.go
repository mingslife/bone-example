package service

import (
	"context"
	"fmt"

	"github.com/mingslife/bone"
)

type HelloService struct {
}

func (s *HelloService) Hello(ctx context.Context, name string) string {
	if name == "" {
		return "Hello!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

var _ bone.Service = (*HelloService)(nil)
