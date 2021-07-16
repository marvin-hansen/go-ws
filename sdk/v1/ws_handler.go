package v1

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	t "go-ws/sdk/types"
	"log"
)

func (s SDKImpl) StopMessageProcessing() {
	ws.StopReadingByteMessages()
}

func (s SDKImpl) StartMessageProcessing() (err error) {
	errHandler := logError
	handler := s.getWSMessageHandler(errHandler)
	err = ws.StartReadingByteMessages(handler, errHandler)
	if err != nil {
		log.Println("error starting message processing!")
		logError(err)
		return err
	}
	return nil
}

func (s SDKImpl) getWSErrorHandler() (errHandler t.WsErrHandler) {
	errHandler = func(err error) {
		if err != nil {
			log.Println(err)
		}
	}
	return errHandler
}

func (s SDKImpl) getWSMessageHandler(errHandler t.WsErrHandler) (wsHandler t.WsHandler) {
	wsHandler = func(message []byte) {
		err := s.processMessage(message, errHandler)
		errHandler(err)
	}
	return wsHandler
}

func printRawMsg(message []byte) {
	msg := string(message)
	log.Println("raw message: ")
	log.Println(msg)
}

func (s SDKImpl) processMessage(message []byte, errHandler t.WsErrHandler) (err error) {
	var dataMessage = new(t.DataMessage)
	messageType := s.getMessageType(message, errHandler)
	dataMessage.MessageType = &messageType

	switch messageType {

	case t.ORDER_EXEC_REPORT_SNAPSHOT:
		// https://docs.coinapi.io/oeml.html#order_exec_report_snapshot-in
		msg := new(t.OrderExecutionReport)
		_ = json.Unmarshal(message, msg)
		dataMessage.OrderExecutionReportSnapshot = msg
		err = executionSnapshotInvoke(dataMessage)
		errHandler(err)
		return checkError(err)

	case t.ORDER_EXEC_REPORT_UPDATE:
		// https://docs.coinapi.io/oeml.html#order_exec_report_update-in
		msg := new(t.OrderExecutionReport)
		_ = json.Unmarshal(message, msg)
		dataMessage.OrderExecutionReportUpdate = msg
		err = executionUpdateInvoke(dataMessage)
		errHandler(err)
		return checkError(err)

	case t.BALANCE_SNAPSHOT:
		// https://docs.coinapi.io/oeml.html#balance_snapshot-in
		msg := new(t.Balance)
		_ = json.Unmarshal(message, msg)
		dataMessage.BalanceSnapshot = msg
		err = balanceSnapshotInvoke(dataMessage)
		errHandler(err)
		return checkError(err)

	case t.BALANCE_UPDATE:
		// https://docs.coinapi.io/oeml.html#balance_update-in
		msg := new(t.Balance)
		_ = json.Unmarshal(message, msg)
		dataMessage.BalanceUpdate = msg
		err = balanceUpdateInvoke(dataMessage)
		errHandler(err)
		return checkError(err)

	case t.POSITION_SNAPSHOT:
		// https://docs.coinapi.io/oeml.html#position_snapshot-in
		msg := new(t.Position)
		_ = json.Unmarshal(message, msg)
		dataMessage.PositionSnapshot = msg
		err = positionSnapshotInvoke(dataMessage)
		errHandler(err)
		return checkError(err)

	case t.POSITION_UPDATE:
		// https://docs.coinapi.io/oeml.html#position_update-in
		msg := new(t.Position)
		_ = json.Unmarshal(message, msg)
		dataMessage.PositionUpdate = msg
		err = positionUpdateInvoke(dataMessage)
		errHandler(err)
		return checkError(err)

	case t.SYMBOLS_SNAPSHOT:
		// https://docs.coinapi.io/oeml.html#symbols_snapshot-in
		msg := new(t.Symbols)
		_ = json.Unmarshal(message, msg)
		dataMessage.SymbolSnapshot = msg

		exchangeID := *dataMessage.SymbolSnapshot.ExchangeId
		s.AddNewSymbolMap(exchangeID)
		for _, symbolData := range *dataMessage.SymbolSnapshot.Data {
			s.addSymbolData(exchangeID, symbolData)
		}

		err = symbolSnapshotInvoke(dataMessage)
		errHandler(err)
		return checkError(err)

	case t.SERVER_INFO:
		// https://docs.coinapi.io/oeml.html#symbols_snapshot-in
		msg := new(t.ServerInfo)
		_ = json.Unmarshal(message, msg)
		dataMessage.ServerInfo = msg
		err = serverInfoInvoke(dataMessage)
		errHandler(err)
		return checkError(err)

	case t.MESSAGE:
		msg := new(t.Message)
		_ = json.Unmarshal(message, msg)
		dataMessage.Message = msg
		err = errorMessageInvoke(dataMessage)
		errHandler(err)
		return checkError(err)
	}

	return nil
}

func (s SDKImpl) getMessageType(message []byte, errHandler t.WsErrHandler) (messageType t.MessageType) {
	j, err := newJSON(message)
	if err != nil {
		errHandler(err)
		return
	}
	messageType = t.MessageType(j.Get("type").MustString())
	return messageType
}

func newJSON(data []byte) (j *simplejson.Json, err error) {
	j, err = simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}
