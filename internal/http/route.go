package http

import (
	"github.com/chapa-ai/application-gaspar/internal/util"
	"net/http"

	"github.com/chapa-ai/application-gaspar/internal/dao"
	"github.com/go-chi/render"
)

func GetOrders(orderDao dao.OrderRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := GetOrdersRequest{}
		if err := req.Bind(r); err != nil {
			err = render.Render(w, r, ErrInvalidRequest(err))
			if err != nil {
				util.LogErrorf("error rendering response: %s", err.Error())
				return
			}
			return
		}

		res, err := orderDao.OrdersByEmail(r.Context(), req.Email)
		if err != nil {
			util.LogErrorf("error orderDao.OrdersByEmail: %s", err.Error())
			return
		}
		if res == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		err = render.RenderList(w, r, ListOrdersResponse(res))
		if err != nil {
			util.LogErrorf("error rendering response: %s", err.Error())
			return
		}
	}
}

func MakeOrder(orderDao dao.OrderRepo, roomDao dao.RoomRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &MakeOrderRequest{}
		if err := req.Bind(r); err != nil {
			err = render.Render(w, r, ErrInvalidRequest(err))
			if err != nil {
				util.LogErrorf("error rendering response: %s", err.Error())
				return
			}
			return
		}

		if err := req.Validate(r.Context(), orderDao, roomDao); err != nil {
			err = render.Render(w, r, ErrInvalidRequest(err))
			if err != nil {
				util.LogErrorf("error rendering response: %s", err.Error())
				return
			}
			return
		}
		err := orderDao.Save(r.Context(), req.UserEmail, &(req.Order))
		if err != nil {
			err = render.Render(w, r, ErrInvalidRequest(err))
			if err != nil {
				util.LogErrorf("error rendering response: %s", err.Error())
				return
			}
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
