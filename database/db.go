package database

import "errors"

type Insertable interface {
	SetId(id uint64)
	GetId() uint64
}

type DBApplication[T Insertable] struct {
	data map[uint64]T
}

var InMemoryDbNotFoundError = errors.New("the resource wasn't found")

func (d *DBApplication[T]) GetAll() map[uint64]T {
	return d.data
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
