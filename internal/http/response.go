package http

import (
	"net/http"

	"github.com/chapa-ai/application-gaspar/internal/dao"
	"github.com/go-chi/render"
)

type (
	OrderResponse struct {
		dao.Order
	}

	ErrResponse struct {
		Err            error  `json:"-"`
		HTTPStatusCode int    `json:"-"`
		StatusText     string `json:"status"`
		ErrorText      string `json:"error,omitempty"`
	}
)

func (o OrderResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func ListOrdersResponse(orders []dao.Order) []render.Renderer {
	var list []render.Renderer
	for _, order := range orders {
		list = append(list, NewOrderResponse(order))
	}
	return list
}

func NewOrderResponse(order dao.Order) *OrderResponse {
	resp := &OrderResponse{order}

	return resp
}

func (e ErrResponse) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Bad request.",
		ErrorText:      err.Error(),
	}
}
