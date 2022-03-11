package main

import (
	"github.com/evanjennings/progprac/app/config"
	"github.com/evanjennings/progprac/app/handler"
	"github.com/evanjennings/progprac/app/server"
	"github.com/kevinanthony/gorps/http"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	_ = cfg

	rh := http.NewRequestHandler(http.NewRequestHandlerHelper())

	calc := handler.NewCalc(rh)

	server.NewServer(rh, calc).Run()
}
