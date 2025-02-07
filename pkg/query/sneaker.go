package query

import "context"

type Sneaker struct{}

type Sneakers []*Sneaker

type SneakerQuery interface {
	GetAll(ctx context.Context, id int) (Sneakers, error)
	GetByID(id int) (*Sneaker, error)
}
