package main

import (
	"github.com/mingslife/bone"
	"github.com/mingslife/bone-example/pkg/conf"
	"github.com/mingslife/bone-example/pkg/module/hello"
)

var cfg *conf.Config

func main() {
	app := bone.NewApplication(&bone.ApplicationOptions{
		Host: cfg.Host,
		Port: cfg.Port,
	})
	app.Use(
		new(hello.Module),
	)

	app.Run()
}

func init() {
	cfg = conf.ParseConfig()
}
