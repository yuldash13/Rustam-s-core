package model

func NewOrderProcessStarted(orderInitialized OrderInitialized, orderStates []OrderState, error error) *OrderProcessStarted {
	return &OrderProcessStarted{
		OrderInitialized: orderInitialized,
		OrderStates:      orderStates,
		Error:            error,
	}
}
