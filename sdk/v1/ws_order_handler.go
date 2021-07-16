package v1

import (
	t "go-ws/sdk/types"
	"log"
)

const verbose = false

func (s SDKImpl) PlaceSingleOrder(req *t.OrderNewSingleRequest) (err error) {

	b, err := req.MarshalJSON()
	logError(err)

	if verbose {
		println("PlaceSingleOrder")
		println(string(b))
	}

	err = ws.WriteJSONMessage(b)
	if err != nil {
		log.Println("can't send message!")
		logError(err)
		return err
	}
	return checkError(err)
}

func (s SDKImpl) CancelSingleOrder(req *t.OrderCancelSingleRequest) (err error) {
	b, err := req.MarshalJSON()

	if verbose {
		println("CancelSingleOrder")
		println(string(b))
	}

	logError(err)
	err = s.sendMessage(b)
	return checkError(err)
}

func (s SDKImpl) CancelAllOrders(req *t.OrderCancelAllRequest) (err error) {
	b, err := req.MarshalJSON()

	if verbose {
		println("CancelAllOrders")
		println(string(b))
	}

	logError(err)
	err = s.sendMessage(b)
	return checkError(err)
}

func (s SDKImpl) sendMessage(b []byte) (err error) {
	err = ws.WriteByteMessage(b)
	if err != nil {
		log.Println("can't send message!")
		logError(err)
		return err
	}
	return nil
}
