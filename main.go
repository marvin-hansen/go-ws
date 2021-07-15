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

func main() {
	sdk := getSDK()
	TestOrder(sdk)
	//TestConnection(sdk)
}

func TestOrder(sdk t.SDK) {
	printHeader("TestOrder!")
	time.Sleep(1 * time.Second)

	printHeader("Lookup symbol!")

	exchangeID := "BINANCE"
	baseID := "ATOM"
	quoteID := "BUSD"
	printSymbol(sdk, exchangeID, baseID, quoteID)

	exchangeID = "BINANCE"
	baseID = "BTC"
	quoteID = "USDT"
	printSymbol(sdk, exchangeID, baseID, quoteID)

	printHeader("Construct order!")

	var err error
	clientOrderID := "BINANCE-7d8a-4888-a733-6007093f8332"
	//reqOrder := t.NewOrderNewSingleRequest(exchangeID, clientOrderID )
	//err := sdk.PlaceSingleOrder(reqOrder)
	//logError(err)

	printHeader("Place  order!")

	printHeader("Cancel order!")
	reqCancel := sdk.NewCancelSingleOrderRequest(exchangeID, "", clientOrderID)
	err = sdk.CancelSingleOrder(reqCancel)
	logError(err)

	printHeader("Cancel all orders!")
	reqCacelAll := sdk.NewCancelAllOrdersRequest(exchangeID)
	err = sdk.CancelAllOrders(reqCacelAll)
	logError(err)

	println(" * CloseConnection!")
	_ = sdk.CloseConnection()

	println("Goodbye!")

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
	println("Start: " + msg)
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
