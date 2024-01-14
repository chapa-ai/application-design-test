package dao

import (
	"context"
	"fmt"
)

type (
	OrderDao struct {
		*baseDao[*Order]
	}
)

func NewOrderDao() *OrderDao {
	return &OrderDao{baseDao: &baseDao[*Order]{}}
}

func (o *OrderDao) OrdersByEmail(ctx context.Context, email string) ([]Order, error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	var ordersByEmail []Order

	if orders, ok := o.storage[email]; ok {
		for _, order := range orders {
			ordersByEmail = append(ordersByEmail, *order)
		}
	} else {
		return nil, fmt.Errorf("there are no such email")
	}

	return ordersByEmail, nil
}
