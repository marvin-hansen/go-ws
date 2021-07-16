package main

import (
	"errors"
	SDK "go-ws/sdk"
	t "go-ws/sdk/types"
	"log"
	"time"
)

//  kubectl port-forward svc/oeml-api-composite 8080:80
const url = "ws://127.0.0.1:8080"
const verbose = false
const exchangeID = "BINANCE"

func main() {
	sdk := getSDK()
	TestOrder(sdk)

	//TestCancelAll(sdk)
	//TestRequests(sdk)
	//TestSymbolLookup(sdk)
	//TestConnection(sdk)
}

func TestCancelAll(sdk t.SDK) {
	time.Sleep(1 * time.Second)

	printHeader(" * Construct cancel all request!")
	reqCancelAll := sdk.NewCancelAllOrdersRequest(exchangeID)
	b, _ := reqCancelAll.MarshalJSON()
	println(string(b))

	printHeader("Cancel all orders!")
	err := sdk.CancelAllOrders(&reqCancelAll)
	logError(err)
	time.Sleep(7 * time.Second)

	println()
	println(" * CloseConnection!")
	_ = sdk.CloseConnection()
	println("Goodbye!")

}

func TestRequests(sdk t.SDK) {
	time.Sleep(1 * time.Second)

	baseID := "BTC"
	quoteID := "USDT"
	printHeader(" * Construct order!")
	symbolData, _ := sdk.LookupSymbolData(exchangeID, baseID, quoteID)
	symbolIdExchange := *symbolData.Symbol_id_base_exchange
	symbolIdCoinapi := *symbolData.Symbol_id_coinapi
	clientOrderID := "BINANCE-7d8a-4888"
	amountOrder := 0.1
	price := 30000.00
	orderSide := t.BUY
	orderType := t.LIMIT
	timeInForce := t.GOOD_TILL_CANCEL

	printHeader(" * Construct order request!")

	reqOrder := sdk.NewSingleOrderRequest(exchangeID, symbolIdExchange, symbolIdCoinapi, clientOrderID, amountOrder, price, orderSide, orderType, timeInForce)
	b, _ := reqOrder.MarshalJSON()
	println(string(b))

	printHeader(" * Construct cancel request!")
	reqCancel := sdk.NewCancelSingleOrderRequest(exchangeID, "", clientOrderID)
	b, _ = reqCancel.MarshalJSON()
	println(string(b))

	printHeader(" * Construct cancel all request!")
	reqCancelAll := sdk.NewCancelAllOrdersRequest(exchangeID)
	b, _ = reqCancelAll.MarshalJSON()
	println(string(b))

	println()
	println(" * CloseConnection!")
	_ = sdk.CloseConnection()
	println("Goodbye!")
}

func TestOrder(sdk t.SDK) {
	printHeader("TestOrder!")

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

	printHeader(" * Construct order!")
	symbolData, _ := sdk.LookupSymbolData(exchangeID, baseID, quoteID)
	symbolIdExchange := *symbolData.Symbol_id_base_exchange
	symbolIdCoinapi := *symbolData.Symbol_id_coinapi
	clientOrderID := "BINANCE-7d8a-4888"
	amountOrder := 0.1
	price := 30000.00
	orderSide := t.BUY
	orderType := t.LIMIT
	timeInForce := t.GOOD_TILL_CANCEL

	printHeader(" * Construct order request!")

	reqOrder := sdk.NewSingleOrderRequest(exchangeID, symbolIdExchange, symbolIdCoinapi, clientOrderID, amountOrder, price, orderSide, orderType, timeInForce)
	b, _ := reqOrder.MarshalJSON()
	println(string(b))

	printHeader(" * Place  order!")
	var err error
	err = sdk.PlaceSingleOrder(&reqOrder)
	logError(err)
	time.Sleep(10 * time.Second)

	printHeader(" * Construct cancel request!")
	reqCancel := sdk.NewCancelSingleOrderRequest(exchangeID, "", clientOrderID)
	b, _ = reqCancel.MarshalJSON()
	println(string(b))

	printHeader(" * Cancel order!")
	err = sdk.CancelSingleOrder(&reqCancel)
	logError(err)
	time.Sleep(10 * time.Second)

	printHeader(" * Construct cancel all request!")
	reqCancelAll := sdk.NewCancelAllOrdersRequest(exchangeID)
	b, _ = reqCancelAll.MarshalJSON()
	println(string(b))

	printHeader("Cancel all orders!")
	err = sdk.CancelAllOrders(&reqCancelAll)
	logError(err)
	time.Sleep(10 * time.Second)

	println()
	println(" * CloseConnection!")
	_ = sdk.CloseConnection()

	println("Goodbye!")

}

func TestSymbolLookup(sdk t.SDK) {

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

func TestConnection(sdk t.SDK) {
	printHeader("TestConnection!")

	time.Sleep(3 * time.Second)
	println(" * CloseConnection!")
	_ = sdk.CloseConnection()

	println("Goodbye!")
}

func getSDK() (sdk t.SDK) {
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
		println("OK:")
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
		msg := message.Message.GetMessage()
		if msg != "" {
			log.Println(mtd+"ErrorMessage: ", msg)
			return errors.New(msg)
		}
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
		log.Println("ServerInfo/Heartbeat")
		if verbose {
			msg := message
			log.Println(msg.ServerInfo)
			println()
		}

	case t.ORDER_EXEC_REPORT_SNAPSHOT:
		log.Println("OrderExecutionReportSnapshot")
		msg := message
		log.Println(msg.OrderExecutionReportSnapshot)
		println()
	case t.ORDER_EXEC_REPORT_UPDATE:
		log.Println("OrderExecutionReportUpdate")
		msg := message
		log.Println(msg.OrderExecutionReportUpdate)
		println()
	case t.BALANCE_SNAPSHOT:
		msg := message
		log.Println("BalanceSnapshot")
		log.Println(msg.BalanceSnapshot)
		println()
	case t.BALANCE_UPDATE:
		msg := message
		log.Println("BalanceUpdate")
		log.Println(msg.BalanceUpdate)
		println()
	case t.POSITION_SNAPSHOT:
		log.Println("PositionSnapshot")
		msg := message
		log.Println(msg.PositionSnapshot)
		println()
	case t.POSITION_UPDATE:
		log.Println("PositionUpdate")
		msg := message
		log.Println(msg.PositionUpdate)
		println()
	case t.SYMBOLS_SNAPSHOT:
		log.Println("SymbolSnapshot")
		if verbose {
			msg := message
			log.Println(msg.SymbolSnapshot)
		}
		println()
	case t.MESSAGE:
		log.Println("Message")
		msg := message
		log.Println(msg.Message)
		println()
	}
}
