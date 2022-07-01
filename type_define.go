package shopeego

type OrderStatus string

const (
	// UNPAID:Order is created, buyer has not paid yet.
	OrderStatusUnpaid OrderStatus = "UNPAID"
	// READY_TO_SHIP:Seller can arrange shipment.
	OrderStatusReadyToShip OrderStatus = "READY_TO_SHIP"
	// PROCESSED:Seller has arranged shipment online and got tracking number from 3PL.
	OrderStatusProcessed OrderStatus = "PROCESSED"
	// RETRY_SHIP:3PL pickup parcel fail. Need to re arrange shipment.
	OrderStatusRetryShip OrderStatus = "RETRY_SHIP"
	// SHIPPED:The parcel has been drop to 3PL or picked up by 3PL.
	OrderStatusShipped OrderStatus = "SHIPPED"
	// TO_CONFIRM_RECEIVE:The order has been received by buyer.
	OrderStatusToConfirmReceive OrderStatus = "TO_CONFIRM_RECEIVE"
	// IN_CANCEL:The order's cancelation is under processing.
	OrderStatusInCancel OrderStatus = "IN_CANCEL"
	// CANCELLED:The order has been canceled.
	OrderStatusCancelled OrderStatus = "CANCELLED"
	// TO_RETURN:The buyer requested to return the order and order's return is processing.
	OrderStatusToReturn OrderStatus = "TO_RETURN"
	// COMPLETED:The order has been completed.
	OrderStatusCompleted OrderStatus = "COMPLETED"
)

type LogisticsStatus string

const (
	// Initial status, order not ready for fulfilment
	LogisticsStatusLogisticsNotStarted LogisticsStatus = "LOGISTICS_NOT_STARTED"
	// order arranged shipment
	LogisticsStatusLogisticsRequestCreated LogisticsStatus = "LOGISTICS_REQUEST_CREATED"
	// order handed over to 3PL
	LogisticsStatusPickupDone LogisticsStatus = "LOGISTICS_PICKUP_DONE"
	// order pending 3PL retry pickup
	LogisticsStatusPickupRetry LogisticsStatus = "LOGISTICS_PICKUP_RETRY"
	// order cancelled by 3PL due to failed pickup or picked up but not able to proceed with delivery
	LogisticsStatusPickupFailed LogisticsStatus = "LOGISTICS_PICKUP_FAILED"
	// order successfully delivered
	LogisticsStatusDeliveryDone LogisticsStatus = "LOGISTICS_DELIVERY_DONE"
	// order cancelled due to 3PL delivery failed
	LogisticsStatusDeliveryFailed LogisticsStatus = "LOGISTICS_DELIVERY_FAILED"
	// order cancelled when order at LOGISTICS_REQUEST_CREATED
	LogisticsStatusRequestCanceled LogisticsStatus = "LOGISTICS_REQUEST_CANCELED"
	// Integrated logistics COD: Order rejected for COD.
	LogisticsStatusLogisticsCODRejected LogisticsStatus = "LOGISTICS_COD_REJECTED"
	// order ready for fulfilment from payment perspective: non-COD: order paid; COD: order passed COD screening
	LogisticsStatusLogisticsReady LogisticsStatus = "LOGISTICS_READY"
	// order cancelled when order at LOGISTICS_READY
	LogisticsStatusLogisticsInvalid LogisticsStatus = "LOGISTICS_INVALID"
	// order cancelled due to 3PL lost the order
	LogisticsStatusLogisticsLost LogisticsStatus = "LOGISTICS_LOST"
	// order logistics pending arrangement
	LogisticsStatusLogisticsPenddingArrange LogisticsStatus = "LOGISTICS_PENDING_ARRANGE"
)
