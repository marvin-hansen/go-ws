package v1

import "go-ws/sdk/types"

func (s SDKImpl) NewSingleOrderRequest(exchangeId, symbolIdExchange, symbolIdCoinapi string, clientOrderId string, amountOrder float64, price float64, side types.OrdSide, orderType types.OrdType, timeInForce types.TimeInForce) (req types.OrderNewSingleRequest) {
	req = types.OrderNewSingleRequest{
		MessageType:      types.ORDER_NEW_SINGLE_REQUEST,
		ExchangeId:       exchangeId,
		SymbolIdExchange: &symbolIdExchange,
		SymbolIdCoinapi:  &symbolIdCoinapi,
		ClientOrderId:    clientOrderId,
		AmountOrder:      amountOrder,
		Price:            price,
		Side:             side,
		OrderType:        orderType,
		TimeInForce:      timeInForce,
	}
	return req
}

// NewCancelSingleOrderRequest constructs a new cancel request.
// One of the properties (`exchange_order_id`, `client_order_id`) is required to identify the order.
// exchangeId: 	Identifier of the exchange from which active orders should be canceled.
func (s SDKImpl) NewCancelSingleOrderRequest(exchangeId, exchangeOrderId, clientOrderId string) (req types.OrderCancelSingleRequest) {

	r := new(types.OrderCancelSingleRequest)
	r.MessageType = types.ORDER_CANCEL_SINGLE_REQUEST
	r.ExchangeId = exchangeId

	if exchangeOrderId == "" {
		r.ClientOrderId = &clientOrderId
	}

	if clientOrderId == "" {
		r.ExchangeOrderId = &exchangeOrderId
	}

	return *r
}

// NewCancelAllOrdersRequest cancels all open orders at the exchange
// exchangeId: 	Identifier of the exchange from which active orders should be canceled.
func (s SDKImpl) NewCancelAllOrdersRequest(exchangeId string) (req types.OrderCancelAllRequest) {
	req = types.OrderCancelAllRequest{
		MessageType: types.ORDER_CANCEL_ALL_REQUEST,
		ExchangeId:  exchangeId,
	}
	return req
}
