package v1

import (
	t "go-ws/sdk/types"
	"log"
)

func (s SDKImpl) SendSingleOrder(req t.OrderNewSingleRequest) (err error) {
	b, err := req.MarshalJSON()
	logError(err)
	err = s.sendMessage(b)
	return checkError(err)
}

func (s SDKImpl) SendCancelSingleOrder(req t.OrderCancelSingleRequest) (err error) {
	b, err := req.MarshalJSON()
	logError(err)
	err = s.sendMessage(b)
	return checkError(err)
}

func (s SDKImpl) SendCancelAllOrders(req t.OrderCancelAllRequest) (err error) {
	b, err := req.MarshalJSON()
	logError(err)
	err = s.sendMessage(b)
	return checkError(err)
}

func (s SDKImpl) sendMessage(b []byte) (err error) {
	err = s.ws.WriteByteMessage(b)
	if err != nil {
		log.Println("can't send Hello message!")
		logError(err)
		return err
	}
	return nil
}

func (s SDKImpl) StopMessageProcessing() {
	s.ws.StopReadingByteMessages()
}

func (s SDKImpl) StartMessageProcessing() (err error) {
	errHandler := logError
	handler := s.getWSMessageHandler(errHandler)
	err = s.ws.StartReadingByteMessages(handler, errHandler)
	if err != nil {
		log.Println("error starting message processing!")
		logError(err)
		running = false
		return err
	}
	running = true
	return nil
}
