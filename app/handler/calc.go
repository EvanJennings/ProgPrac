package handler

import "github.com/kevinanthony/gorps/http"

type Calc interface {
}

type calc struct {
	rh http.RequestHandler
}

func NewCalc(rh http.RequestHandler) Calc {
	if rh == nil {
		panic("request handler is required")
	}

	return calc{
		rh: rh,
	}
}
