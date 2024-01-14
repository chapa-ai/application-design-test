package dao

import (
	"context"
	"sync"
)

type (
	Model interface {
		*Order | *Room
	}

	baseDao[T Model] struct {
		mu sync.Mutex
		//storage []T
		storage map[string][]T
	}
)

func (d *baseDao[T]) Save(_ context.Context, key string, src T) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.storage == nil {
		d.storage = make(map[string][]T)
	}

	if _, ok := d.storage[key]; !ok {
		d.storage[key] = make([]T, 0)
	}

	d.storage[key] = append(d.storage[key], src)

	return nil
}
