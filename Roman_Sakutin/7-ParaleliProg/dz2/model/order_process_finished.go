package model

func NewOrderProcessFinished(orderFinishedExternalInteraction OrderFinishedExternalInteraction, orderStates []OrderState, error error) *OrderProcessFinished {
	return &OrderProcessFinished{
		OrderFinishedExternalInteraction: orderFinishedExternalInteraction,
		OrderStates:                      orderStates,
		Error:                            error,
	}
}
