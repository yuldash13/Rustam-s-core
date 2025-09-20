package model

func NewOrderFinishedExternalInteraction(orderProcessStarted OrderProcessStarted, storageID int, pickupPointID int, orderStates []OrderState, error error) *OrderFinishedExternalInteraction {
	return &OrderFinishedExternalInteraction{
		OrderProcessStarted: orderProcessStarted,
		StorageID:           storageID,
		PickupPointID:       pickupPointID,
		OrderStates:         orderStates,
		Error:               error,
	}
}
