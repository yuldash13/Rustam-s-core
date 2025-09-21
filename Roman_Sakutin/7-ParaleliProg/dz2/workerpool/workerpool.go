package workerpool

import (
	"Roman_Sakutin/Roman_Sakutin/7-ParaleliProg/dz2/model"
	"Roman_Sakutin/Roman_Sakutin/7-ParaleliProg/dz2/pipeline"
	"context"
	"fmt"
	"sync"
)

type OrderWorkerPool interface {
	StartWorkerPool(ctx context.Context, orders <-chan model.OrderInitialized, additionalActions model.OrderActions, workersCount int) <-chan model.OrderProcessFinished
}

type OrderWorkerPoolImplementation struct{}

func NewOrderWorkerPoolImplementation() *OrderWorkerPoolImplementation {
	return &OrderWorkerPoolImplementation{}
}

func (o *OrderWorkerPoolImplementation) StartWorkerPool(
	ctx context.Context,
	orders <-chan model.OrderInitialized,
	additionalActions model.OrderActions,
	workersCount int,
) <-chan model.OrderProcessFinished {

	var ch = make(chan model.OrderProcessFinished)
	wg := sync.WaitGroup{}
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				pipelineBoss := pipeline.NewOrderPipelineImplementation()
				pipelineBoss.Start(ctx, additionalActions, orders, ch)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
		fmt.Println("workerPool done")
	}()
	return ch
}
