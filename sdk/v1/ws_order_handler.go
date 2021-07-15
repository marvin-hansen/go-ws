package v1

import (
	t "go-ws/sdk/types"
	"log"
)

func (s SDKImpl) PlaceSingleOrder(req *t.OrderNewSingleRequest) (err error) {
	b, err := req.MarshalJSON()
	logError(err)
	err = s.sendMessage(b)
	return checkError(err)
}

func (s SDKImpl) CancelSingleOrder(req *t.OrderCancelSingleRequest) (err error) {
	b, err := req.MarshalJSON()
	logError(err)
	err = s.sendMessage(b)
	return checkError(err)
}

func (s SDKImpl) CancelAllOrders(req *t.OrderCancelAllRequest) (err error) {
	b, err := req.MarshalJSON()
	logError(err)
	err = s.sendMessage(b)
	return checkError(err)
}

func (s SDKImpl) sendMessage(b []byte) (err error) {
	err = s.ws.WriteByteMessage(b)
	if err != nil {
		log.Println("can't send message!")
		logError(err)
		return err
	}
	return nil
}
