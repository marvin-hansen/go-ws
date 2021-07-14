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

func main() {
	sdk := getSDK()

	TestConnection(sdk)
}

func TestOrder(sdk t.SDK) {
	printHeader("TestOrder!")

	time.Sleep(1 * time.Second)

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

func printHeader(msg string) {
	println()
	println("=====================")
	println("Start: " + msg)
	println("=====================")
	println()
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
		log.Println("ServerInfo")

		msg := message
		log.Println(msg.ServerInfo)
		println()
	case t.ORDER_EXEC_REPORT_SNAPSHOT:
		msg := message
		log.Println("OrderExecutionReportSnapshot")
		log.Println(msg.OrderExecutionReportSnapshot)
		println()
	case t.ORDER_EXEC_REPORT_UPDATE:
		msg := message
		log.Println("OrderExecutionReportUpdate")
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
		msg := message
		log.Println("PositionSnapshot")
		log.Println(msg.PositionSnapshot)
		println()
	case t.POSITION_UPDATE:
		msg := message
		log.Println("PositionUpdate")
		log.Println(msg.PositionUpdate)
		println()
	case t.SYMBOLS_SNAPSHOT:
		log.Println("SymbolSnapshot")
		msg := message
		log.Println(msg.SymbolSnapshot)
		println()
	case t.MESSAGE:
		log.Println("Message")
		msg := message
		log.Println(msg.Message)
		println()
	}
}
