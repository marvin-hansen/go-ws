package main

import (
	"errors"
	t "go-ws/sdk/types"
	"log"
)

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
		msg := message
		log.Println(msg)
		println()
	case t.ORDER_EXEC_REPORT_SNAPSHOT:
		msg := message
		log.Println(msg)
		println()
	case t.ORDER_EXEC_REPORT_UPDATE:
		msg := message
		log.Println(msg)
		println()
	case t.BALANCE_SNAPSHOT:
		msg := message
		log.Println(msg)
		println()
	case t.BALANCE_UPDATE:
		msg := message
		log.Println(msg)
		println()
	case t.POSITION_SNAPSHOT:
		msg := message
		log.Println(msg)
		println()
	case t.POSITION_UPDATE:
		msg := message
		log.Println(msg)
		println()
	case t.SYMBOLS_SNAPSHOT:
		msg := message
		log.Println(msg)
		println()
	case t.MESSAGE:
		msg := message
		log.Println(msg)
		println()
	}
}
