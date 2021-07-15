package v1

import "go-ws/sdk/types"

func (s SDKImpl) NewSingleOrderRequest(symbolData types.SymbolData, exchangeId string, clientOrderId string, amountOrder float64, price float64, side types.OrdSide, orderType types.OrdType, timeInForce types.TimeInForce) (req *types.OrderNewSingleRequest) {
	req = &types.OrderNewSingleRequest{
		ExchangeId:       exchangeId,
		SymbolIdExchange: symbolData.Symbol_id_base_exchange,
		SymbolIdCoinapi:  symbolData.Symbol_id_coinapi,
		ClientOrderId:    clientOrderId,
		AmountOrder:      amountOrder,
		Price:            price,
		Side:             side,
		OrderType:        orderType,
		TimeInForce:      timeInForce,
	}
	return req
}

func (s SDKImpl) NewCancelSingleOrderRequest(exchangeId, exchangeOrderId, clientOrderId string) (req *types.OrderCancelSingleRequest) {
	req = &types.OrderCancelSingleRequest{
		ExchangeId:      exchangeId,
		ExchangeOrderId: &exchangeOrderId,
		ClientOrderId:   &clientOrderId,
	}
	return req
}

func (s SDKImpl) NewCancelAllOrdersRequest(exchangeId string) (req *types.OrderCancelAllRequest) {
	req = &types.OrderCancelAllRequest{
		ExchangeId: exchangeId,
	}
	return req
}
