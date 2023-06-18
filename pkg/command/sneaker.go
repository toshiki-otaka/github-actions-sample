package command

import "context"

type Sneaker struct{}

type SneakerCommand interface {
	Create(ctx context.Context, sneaker *Sneaker) error
	Update(sneaker *Sneaker) error
}
