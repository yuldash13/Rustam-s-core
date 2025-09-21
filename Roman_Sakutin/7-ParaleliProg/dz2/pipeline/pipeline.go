package pipeline

import (
	"Roman_Sakutin/Roman_Sakutin/7-ParaleliProg/dz2/model"
	"context"
	"fmt"
	"sync"
)

type OrderPipeline interface {
	Start(ctx context.Context, actions model.OrderActions, orders <-chan model.OrderInitialized, processed chan<- model.OrderProcessFinished)
}

type OrderPipelineImplementation struct{}

func NewOrderPipelineImplementation() *OrderPipelineImplementation {
	return &OrderPipelineImplementation{}
}

func (o *OrderPipelineImplementation) Start(
	ctx context.Context,
	actions model.OrderActions,
	orders <-chan model.OrderInitialized,
	processed chan<- model.OrderProcessFinished,
) {

	processStarted := o.pipeline1(ctx, actions, orders)
	interaction := o.pipeline2(ctx, actions, processStarted)
	processedFinished := o.pipeline3(ctx, actions, interaction)
	for r := range processedFinished {
		processed <- r
	}
}

func (o *OrderPipelineImplementation) pipeline1(
	ctx context.Context,
	actions model.OrderActions,
	orders <-chan model.OrderInitialized,
) <-chan model.OrderProcessStarted {

	wg := sync.WaitGroup{}
	out := make(chan model.OrderProcessStarted)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("pipeline1 done")
		for order := range orders {
			if order.Error != nil {
				orderProcessStarted := model.NewOrderProcessStarted(order, order.OrderStates, order.Error)
				out <- *orderProcessStarted
				continue
			}
			select {
			case <-ctx.Done():
				return
			default:
				actions.InitToStarted()
				order.OrderStates = append(order.OrderStates, model.ProcessStarted)
				orderProcessStarted := model.NewOrderProcessStarted(order, order.OrderStates, order.Error)
				out <- *orderProcessStarted
			}
		}
	}()
	go func() {
		wg.Wait()
		defer close(out)
	}()
	return out
}

func (o *OrderPipelineImplementation) pipeline2(
	ctx context.Context,
	actions model.OrderActions,
	orders <-chan model.OrderProcessStarted,
) <-chan model.OrderFinishedExternalInteraction {

	wg := sync.WaitGroup{}
	out := make(chan model.OrderFinishedExternalInteraction)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("pipeline2 done")
		fan := o.FanOut(ctx, orders)
		for orderFinishedExternalInteraction := range fan {
			if orderFinishedExternalInteraction.Error != nil {
				out <- orderFinishedExternalInteraction
				continue
			}
			select {
			case <-ctx.Done():
				return
			default:
				actions.StartedToFinishedExternalInteraction()
				out <- orderFinishedExternalInteraction
			}
		}
	}()
	go func() {
		wg.Wait()
		defer close(out)
	}()
	return out
}

func (o *OrderPipelineImplementation) FanOut(
	ctx context.Context,
	orders <-chan model.OrderProcessStarted,
) <-chan model.OrderFinishedExternalInteraction {

	k := 10
	var output = make(chan model.OrderFinishedExternalInteraction)
	wg := sync.WaitGroup{}
	for i := 0; i < k; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer fmt.Println("Fan out done")
			for order := range orders {
				var storageID, pickupPointID int
				if order.Error != nil {
					orderFinishedExternalInteraction := model.NewOrderFinishedExternalInteraction(order, storageID, pickupPointID, order.OrderStates, order.Error)
					output <- *orderFinishedExternalInteraction
					continue
				}
				select {
				case <-ctx.Done():
					return
				default:
					storageID = order.OrderInitialized.ProductID%2 + 1
					pickupPointID = order.OrderInitialized.ProductID%3 + 1
					order.OrderStates = append(order.OrderStates, model.FinishedExternalInteraction)
					orderFinishedExternalInteraction := model.NewOrderFinishedExternalInteraction(order, storageID, pickupPointID, order.OrderStates, order.Error)
					output <- *orderFinishedExternalInteraction
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

//func (o *OrderPipelineImplementation) FanIn(
//	ctx context.Context,
//	input <-chan model.OrderFinishedExternalInteraction,
//) []<-chan model.OrderFinishedExternalInteraction {
//
//	var output = make([]<-chan model.OrderFinishedExternalInteraction, 0)
//	wg := sync.WaitGroup{}
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		defer fmt.Println("Fan in done")
//		for ch := range input {
//			select {
//			case <-ctx.Done():
//				return
//			default:
//				output = append(output, ch)
//			}
//		}
//	}()
//	wg.Wait()
//	return output
//}

func (o *OrderPipelineImplementation) pipeline3(
	ctx context.Context,
	actions model.OrderActions,
	orders <-chan model.OrderFinishedExternalInteraction,
) chan model.OrderProcessFinished {

	wg := sync.WaitGroup{}
	out := make(chan model.OrderProcessFinished)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("pipeline3 done")
		for order := range orders {
			if order.Error != nil {
				orderProcessFinished := model.NewOrderProcessFinished(order, order.OrderStates, order.Error)
				out <- *orderProcessFinished
				continue
			}
			select {
			case <-ctx.Done():
				return
			default:
				actions.FinishedExternalInteractionToProcessFinished()
				order.OrderStates = append(order.OrderStates, model.ProcessFinished)
				orderProcessFinished := model.NewOrderProcessFinished(order, order.OrderStates, order.Error)
				out <- *orderProcessFinished
			}
		}
	}()
	go func() {
		wg.Wait()
		defer close(out)
	}()
	return out
}
