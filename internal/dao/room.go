package dao

import (
	"context"
	"fmt"
	"github.com/chapa-ai/application-gaspar/internal/util"
	"time"
)

type (
	roomDao struct {
		*baseDao[*Room]
	}
)

func NewRoomDao() RoomRepo {
	rooms := &roomDao{baseDao: &baseDao[*Room]{}}
	// data seed
	_ = rooms.Save(context.TODO(), "econom", &Room{Type: "econom", Count: 1})
	_ = rooms.Save(context.TODO(), "standart", &Room{Type: "standart", Count: 1})
	_ = rooms.Save(context.TODO(), "lux", &Room{Type: "lux", Count: 1})

	return rooms
}

func (o *roomDao) GetAvailableRooms(_ context.Context, categoryRoom string, from, to time.Time) ([]Room, error) {

	o.mu.Lock()
	defer o.mu.Unlock()

	if len(o.storage[categoryRoom]) == 0 {
		return nil, fmt.Errorf("there are no such rooms")
	}

	var ret []Room

	for _, room := range o.storage[categoryRoom] {
		orderFrom, _ := time.Parse(time.DateOnly, room.From)
		orderTo, _ := time.Parse(time.DateOnly, room.To)
		if util.TimeOverlap(orderFrom, orderTo, from, to) {
			ret = append(ret, *room)
		}
	}

	return ret, nil
}
