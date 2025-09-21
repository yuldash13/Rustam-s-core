package generator

import (
	"Roman_Sakutin/Roman_Sakutin/7-ParaleliProg/dz2/model"
	"context"
	"fmt"
)

type OrderGenerator interface {
	GenerateOrdersStream(ctx context.Context, orders []model.OrderInitialized) <-chan model.OrderInitialized
}

type OrderGeneratorImplementation struct{}

func NewOrderGeneratorImplementation() *OrderGeneratorImplementation {
	return &OrderGeneratorImplementation{}
}

func (o *OrderGeneratorImplementation) GenerateOrdersStream(
	ctx context.Context,
	orders []model.OrderInitialized,
) <-chan model.OrderInitialized {

	var ch = make(chan model.OrderInitialized)
	go func() {
		defer close(ch)
		for _, r := range orders {
			select {
			case <-ctx.Done():
				return
			case ch <- r:
			}
		}
		fmt.Println("generator done")
	}()
	return ch
}
