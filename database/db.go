package database

import (
	"errors"
)

func NewApplication[T Insertable]() *DBApplication[T] {
	return &DBApplication[T]{
		data: make(map[uint64]T),
	}
}

type Insertable interface {
	SetId(id uint64)
	GetId() uint64
}

type DBApplication[T Insertable] struct {
	data map[uint64]T
}

var InMemoryDbNotFoundError = errors.New("the resource wasn't found")

func (d *DBApplication[T]) GetAll() []T {
	result := make([]T, 0, len(d.data))

	for _, data := range d.data {
		result = append(result, data)
	}
	return result
}

func (d *DBApplication[T]) GetById(id uint64) (T, error) {
	data, ok := d.data[id]

	if !ok {
		var zero T
		return zero, InMemoryDbNotFoundError
	}
	return data, nil
}

func (d *DBApplication[T]) Insert(data T) T {
	id := uint64(len(d.data) + 1)
	data.SetId(id)
	d.data[id] = data
	return data
}

func (d *DBApplication[T]) Update(id uint64, data T) (T, error) {
	oldData, err := d.GetById(id)
	if err != nil {
		return oldData, err
	}
	data.SetId(oldData.GetId())
	d.data[id] = data
	return d.data[id], nil
}

func (d *DBApplication[T]) Delete(id uint64) error {
	if _, ok := d.data[id]; !ok {
		return InMemoryDbNotFoundError
	}

	delete(d.data, id)
	return nil
}
