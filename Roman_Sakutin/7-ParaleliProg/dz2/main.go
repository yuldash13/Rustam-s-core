package main

import (
	"Roman_Sakutin/Roman_Sakutin/7-ParaleliProg/dz2/generator"
	"Roman_Sakutin/Roman_Sakutin/7-ParaleliProg/dz2/model"
	"Roman_Sakutin/Roman_Sakutin/7-ParaleliProg/dz2/workerpool"
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()

	generator1 := generator.NewOrderGeneratorImplementation()

	orders := []model.OrderInitialized{
		*model.NewOrderInitialized(1, 1, nil),
		//*model.NewOrderInitialized(2, 2, nil),
		//*model.NewOrderInitialized(3, 3, nil),
		//*model.NewOrderInitialized(4, 4, nil),
		//*model.NewOrderInitialized(5, 5, nil),
		//*model.NewOrderInitialized(6, 6, nil),
		//*model.NewOrderInitialized(7, 7, nil),
		//*model.NewOrderInitialized(8, 8, nil),
		//*model.NewOrderInitialized(9, 9, nil),
		//*model.NewOrderInitialized(10, 10, nil),
	}

	job := generator1.GenerateOrdersStream(ctx, orders)

	workerPool := workerpool.NewOrderWorkerPoolImplementation()

	additionalActions, countChecker := getDefaultAdditionalActions()
	workersCount := 1

	result := workerPool.StartWorkerPool(ctx, job, additionalActions, workersCount)

	wg := sync.WaitGroup{}
	doneOrders := make([]model.Order, 0)
	for r := range result {
		g := r
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				or := model.NewOrder(
					g.OrderFinishedExternalInteraction.OrderProcessStarted.OrderInitialized.OrderID,
					g.OrderFinishedExternalInteraction.OrderProcessStarted.OrderInitialized.ProductID,
					g.OrderFinishedExternalInteraction.StorageID,
					g.OrderFinishedExternalInteraction.PickupPointID,
					true,
					g.OrderStates,
				)
				doneOrders = append(doneOrders, *or)
				countChecker.initToStartedCounter++
				countChecker.StartedToFinishedExternalInteractionCounter++
				countChecker.FinishedExternalInteractionToProcessFinishedCounter++
			}
		}()
	}
	wg.Wait()
	for i := 0; i < len(doneOrders); i++ {
		fmt.Println(doneOrders[i])
	}
}

type CountChecker struct {
	initToStartedCounter                                int
	StartedToFinishedExternalInteractionCounter         int
	FinishedExternalInteractionToProcessFinishedCounter int
}

func getDefaultAdditionalActions() (model.OrderActions, *CountChecker) {
	countChecker := &CountChecker{}

	initToStarted := func() {
		countChecker.initToStartedCounter++
	}

	startedToFinishedExternalInteraction := func() {
		countChecker.StartedToFinishedExternalInteractionCounter++
	}

	finishedExternalInteractionToProcessFinished := func() {
		countChecker.FinishedExternalInteractionToProcessFinishedCounter++
	}

	orderActions := model.OrderActions{
		InitToStarted:                                initToStarted,
		StartedToFinishedExternalInteraction:         startedToFinishedExternalInteraction,
		FinishedExternalInteractionToProcessFinished: finishedExternalInteractionToProcessFinished,
	}

	return orderActions, countChecker
}
