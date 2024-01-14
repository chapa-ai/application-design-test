package dao

import (
	"context"
	"time"
)

type (
	Base[T Model] interface {
		Save(_ context.Context, key string, src T) error
	}

	OrderRepo interface {
		Base[*Order]

		OrdersByEmail(ctx context.Context, email string) ([]Order, error)
	}

	RoomRepo interface {
		Base[*Room]

		GetAvailableRooms(_ context.Context, categoryRoom string, from, to time.Time) ([]Room, error)

		//GetRoomByType(ctx context.Context, room string) (*Room, error)
	}
)
