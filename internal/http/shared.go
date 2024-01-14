package http

import (
	"github.com/chapa-ai/application-gaspar/internal/config"
	"net/http"

	"github.com/chapa-ai/application-gaspar/internal/dao"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GetServer() *http.Server {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	ordersDao := dao.NewOrderDao()
	roomsDao := dao.NewRoomDao()

	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.Route("/v1", func(r chi.Router) {
		r.Route("/orders", func(r chi.Router) {
			r.Post("/", MakeOrder(ordersDao, roomsDao))
			r.Get("/list", GetOrders(ordersDao))
		})
	})

	return &http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}
}
