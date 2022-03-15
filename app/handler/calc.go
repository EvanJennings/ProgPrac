package handler

import (
	"context"
	"net/http"

	rh "github.com/kevinanthony/gorps/http"
)

type Calc interface {
	Add(ctx context.Context, r *http.Request) (interface{}, error)
	Subtract(ctx context.Context, r *http.Request) (interface{}, error)
	Multiply(ctx context.Context, r *http.Request) (interface{}, error)
	Divide(ctx context.Context, r *http.Request) (interface{}, error)
}

type calc struct {
	rh rh.RequestHandler
}

func (c calc) Add(ctx context.Context, r *http.Request) (interface{}, error) {
	var req addReq
	err := c.rh.MarshalAndVerify(r, &req)
	if err != nil {
		return nil, err
	}

	return addResponse{Sum: req.Body.First + req.Body.Second + req.Body.Third}, nil
}

type addResponse struct {
	Sum int `json:"sum"`
}
type addReq struct {
	Body struct {
		First  int `json:"first"`
		Second int `json:"second"`
		Third  int `json:"third"`
	} `body:"2"`
}

func (c calc) Subtract(ctx context.Context, r *http.Request) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c calc) Multiply(ctx context.Context, r *http.Request) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c calc) Divide(ctx context.Context, r *http.Request) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func NewCalc(rh rh.RequestHandler) Calc {
	if rh == nil {
		panic("request handler is required")
	}

	return calc{
		rh: rh,
	}
}
