package main

import (
	SDK "go-ws/sdk"
	t "go-ws/sdk/types"
	"log"
	"time"
)

// kubectl port-forward svc/oeml-api-composite 8080:80
// websocat -s 8080
const (
	url           = "ws://127.0.0.1:8080"
	wait          = 5
	websocat      = false
	verbose       = false
	exchangeID    = "BINANCE"
	clientOrderID = "BINANCE-7d8a-4888"
)

func main() {
	sdk := getSDK(url)

	TestSymbolLookup(sdk)
	TestPlaceSingleOrder(sdk)
	TestCancelSingleOrder(sdk)
	TestCancelAll(sdk)
	CloseSocket(sdk)
}

func getSDK(url string) (sdk t.SDK) {
	println(" * NewSDK!")
	sdk = SDK.NewSDK(url)

	println(" * SetSysInvoke!")
	sysInvoke := GetSysInvoke()
	sdk.SetSystemInvoke(sysInvoke)

	println(" * SetSysSnapshotInvoke!")
	snapInvoke := GetSnapshotInvoke()
	sdk.SetSnapshotInvoke(snapInvoke)

	println(" * SetSysSnapshotInvoke!")
	updInvoke := GetUpdateInvoke()
	sdk.SetUpdateInvoke(updInvoke)

	return sdk
}

func TestSymbolLookup(sdk t.SDK) {

	if websocat {
		return // cannot test symbol lookup against mock socket
	}
	time.Sleep(1 * time.Second)

	printHeader(" * Lookup symbol!")
	// https://www.binance.com/en/trade/ATOM_BUSD?theme=dark&type=spot
	baseID := "ATOM"
	quoteID := "BUSD"
	printSymbol(sdk, exchangeID, baseID, quoteID)
	time.Sleep(1 * time.Second)

	// https://www.binance.com/en/trade/BTC_USDT?theme=dark&type=spot
	baseID = "BTC"
	quoteID = "USDT"
	printSymbol(sdk, exchangeID, baseID, quoteID)

	time.Sleep(1 * time.Second)
}

func TestPlaceSingleOrder(sdk t.SDK) {

	printHeader(" * Lookup order symbol!")

	var symbolIdCoinApi string
	if websocat {
		// cannot lookup symbol on mock web socket
		symbolIdCoinApi = "BINANCE_SPOT_BTC_USDT"
	} else {
		baseID := "BTC"
		quoteID := "USDT"
		symbolData, _ := sdk.LookupSymbolData(exchangeID, baseID, quoteID)
		symbolIdCoinApi = *symbolData.Symbol_id_coinapi
	}

	printHeader(" Symbol: " + symbolIdCoinApi)

	printHeader(" * Construct order request!")
	amountOrder := 0.045
	price := 0.0783
	orderSide := t.BUY
	orderType := t.LIMIT
	timeInForce := t.GOOD_TILL_CANCEL
	reqOrder := sdk.NewSingleOrderRequest(exchangeID, symbolIdCoinApi, clientOrderID, amountOrder, price, orderSide, orderType, timeInForce)
	b, _ := reqOrder.MarshalJSON()
	println(string(b))

	printHeader("Place single order!")
	err := sdk.PlaceSingleOrder(reqOrder)
	logError(err)
	time.Sleep(wait * time.Second)
}

func TestCancelSingleOrder(sdk t.SDK) {
	printHeader(" * Construct cancel request!")
	reqCancel := sdk.NewCancelSingleOrderRequest(exchangeID, clientOrderID)
	b, _ := reqCancel.MarshalJSON()
	println(string(b))

	printHeader(" * Cancel order!")
	err := sdk.CancelSingleOrder(reqCancel)
	logError(err)
	time.Sleep(wait * time.Second)
}

func TestCancelAll(sdk t.SDK) {
	time.Sleep(1 * time.Second)

	printHeader(" * Construct cancel all request!")
	reqCancelAll := sdk.NewCancelAllOrdersRequest(exchangeID)
	b, _ := reqCancelAll.MarshalJSON()
	println(string(b))

	printHeader("Cancel all orders!")
	err := sdk.CancelAllOrders(reqCancelAll)
	logError(err)
	time.Sleep(wait * time.Second)
}

func CloseSocket(sdk t.SDK) {
	println(" * Close websocket!")
	if websocat {
		return // Don't  close connection of a local mock socket
	}
	println()
	println(" * CloseConnection!")
	_ = sdk.CloseConnection()
	println("Goodbye!")
}

func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func printHeader(msg string) {
	println()
	println("=====================")
	println(msg)
	println("=====================")
	println()
}

func printSymbol(sdk t.SDK, exchangeID, baseID, quoteID string) {
	println("")
	println("Lookup symbol for: " + baseID + "/" + quoteID)
	symbol, ok := sdk.LookupSymbolData(exchangeID, baseID, quoteID)
	if ok {
		println(symbol.String())
	}
	println("")

}

func GetSnapshotInvoke() (snapInv t.SnapshotInvoke) {
	snapInv = t.SnapshotInvoke{
		ExecutionSnapshotInvoke: GetInvokeFunction(t.ORDER_EXEC_REPORT_SNAPSHOT),
		BalanceSnapshotInvoke:   GetInvokeFunction(t.BALANCE_SNAPSHOT),
		PositionSnapshotInvoke:  GetInvokeFunction(t.POSITION_SNAPSHOT),
	}
	return snapInv
}

func GetUpdateInvoke() (updInv t.UpdateInvoke) {
	updInv = t.UpdateInvoke{
		ExecutionUpdateInvoke: GetInvokeFunction(t.ORDER_EXEC_REPORT_UPDATE),
		BalanceUpdateInvoke:   GetInvokeFunction(t.BALANCE_UPDATE),
		PositionUpdateInvoke:  GetInvokeFunction(t.POSITION_UPDATE),
	}
	return updInv
}

func GetSysInvoke() (sysInv t.SystemInvoke) {
	sysInv = t.SystemInvoke{
		ErrorMessageInvoke:   GetErrorInvoke(),
		ServerInfoInvoke:     GetInvokeFunction(t.SERVER_INFO),
		SymbolSnapshotInvoke: GetInvokeFunction(t.SYMBOLS_SNAPSHOT),
	}
	return sysInv
}

func GetErrorInvoke() t.InvokeFunction {
	// You need to be prepared to receive an error message from us when you send something wrong;
	// all errors are permanent and you should expect that the underlying
	// WebSocket connection will be closed by us after sending an error message.
	// Good practice is to store all error messages somewhere for further manual review.
	// https://docs.coinapi.io/#error-handling
	return func(message *t.DataMessage) (err error) {
		mtd := "ErrorHandler: "
		println(mtd)
		msg := message.Message //.GetMessage()
		log.Println("Message")
		log.Println("Type: ", *msg.Type)
		log.Println("Severity: ", *msg.Severity)
		log.Println("Message: ", *msg.Message)

		return nil
	}
}

func GetInvokeFunction(msgType t.MessageType) t.InvokeFunction {
	return func(message *t.DataMessage) (err error) {
		printMessage(msgType, message)
		return nil
	}
}

func printMessage(msgType t.MessageType, message *t.DataMessage) {
	switch msgType {
	case t.SERVER_INFO:
		log.Println("ServerInfo/Heartbeat: " + *message.ServerInfo.ExchangeId)
		if verbose {
			msg := message
			log.Println(msg.ServerInfo)
			println()
		}

	case t.ORDER_EXEC_REPORT_SNAPSHOT:
		log.Println("OrderExecutionReportSnapshot: " + message.OrderExecutionReportSnapshot.ExchangeId)
		if verbose {
			msg := message
			log.Println(msg.OrderExecutionReportSnapshot)
			println()
		}

	case t.ORDER_EXEC_REPORT_UPDATE:
		log.Println("OrderExecutionReportUpdate")
		msg := message
		log.Println(msg.OrderExecutionReportUpdate)
		println()
	case t.BALANCE_SNAPSHOT:
		msg := message
		if verbose {
			log.Println("BalanceSnapshot" + *message.BalanceSnapshot.ExchangeId)
			log.Println(msg.BalanceSnapshot)
			println()
		}
	case t.BALANCE_UPDATE:
		msg := message
		log.Println("BalanceUpdate")
		log.Println(msg.BalanceUpdate.String())
		println()

	case t.POSITION_SNAPSHOT:
		log.Println("PositionSnapshot: " + *message.PositionSnapshot.ExchangeId)
		if verbose {
			msg := message
			log.Println(msg.PositionSnapshot)
			println()
		}
	case t.POSITION_UPDATE:
		log.Println("PositionUpdate")
		msg := message
		log.Println(msg.PositionUpdate)
		println()
	case t.SYMBOLS_SNAPSHOT:
		log.Println("SymbolSnapshot: " + *message.SymbolSnapshot.ExchangeId)
		if verbose {
			msg := message
			log.Println(msg.SymbolSnapshot)
			println()
		}
	case t.MESSAGE:
		log.Println("Message")
		msg := message.Message
		log.Println("Type: ", msg.Type)
		log.Println("Severity: ", msg.Severity)
		log.Println("Message: ", msg.Message)
		println()
	}
}
