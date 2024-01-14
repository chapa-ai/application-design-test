package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/chapa-ai/application-gaspar/internal/dao"
	"github.com/chapa-ai/application-gaspar/internal/util"
)

type (
	MakeOrderRequest struct {
		dao.Order
	}

	GetOrdersRequest struct {
		Email string
	}
)

func (o *MakeOrderRequest) FromAsTime() (time.Time, error) {
	return time.Parse(time.DateOnly, o.From)
}

func (o *MakeOrderRequest) ToAsTime() (time.Time, error) {
	return time.Parse(time.DateOnly, o.To)
}

func (o *MakeOrderRequest) Bind(r *http.Request) error {
	o.Order.UserEmail = r.URL.Query().Get("email")
	if !util.ValidateEmail(o.Order.UserEmail) {
		return fmt.Errorf("email is invalid: %s", o.Order.UserEmail)
	}

	o.Order.CategoryRoom = r.URL.Query().Get("room")
	if o.Order.CategoryRoom == "" {
		return fmt.Errorf("room is required: %s", o.Order.CategoryRoom)
	}

	o.Order.From = r.URL.Query().Get("from")
	if o.Order.From == "" {
		return fmt.Errorf("from is required: %s", o.Order.From)
	}

	o.Order.To = r.URL.Query().Get("to")
	if o.Order.To == "" {
		return fmt.Errorf("to is required: %s", o.Order.To)
	}

	return nil
}

func (o *MakeOrderRequest) Validate(ctx context.Context, orderDao dao.OrderRepo, roomDao dao.RoomRepo) error {
	fromTime, err := o.FromAsTime()
	if err != nil {
		return fmt.Errorf("failed fromTime parsing")
	}
	if time.Now().After(fromTime) {
		return fmt.Errorf("from can not be past: %s", fromTime.Format(time.DateOnly))
	}

	toTime, err := o.ToAsTime()
	if err != nil {
		return fmt.Errorf("failed toTime parsing: %s", err.Error())
	}

	if fromTime.Equal(toTime) {
		return fmt.Errorf("dates %s - %s matches",
			fromTime.Format(time.DateOnly), toTime.Format(time.DateOnly))
	}

	if toTime.Before(fromTime) {
		return fmt.Errorf("invalid date range: %s - %s",
			fromTime.Format(time.DateOnly), toTime.Format(time.DateOnly))
	}

	_, err = roomDao.GetAvailableRooms(ctx, o.CategoryRoom, fromTime, toTime)
	if err != nil {
		return fmt.Errorf("failed GetAvailableRooms: %s", err.Error())
	}

	return nil
}

func (o *GetOrdersRequest) Bind(r *http.Request) error {
	o.Email = r.URL.Query().Get("email")
	if !util.ValidateEmail(o.Email) {
		return fmt.Errorf("email is invalid: %s", o.Email)
	}

	return nil
}
