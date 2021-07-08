package main

import (
	SDK "go-ws/sdk"
)

func main() {
	//TestWebSocket()

	TestConnection()
}

func TestConnection() {
	printHeader("TestConnection!")

	println(" * NewSDK!")
	sdk := SDK.NewSDK()

	println(" * SetErrorInvoke!")
	errH := GetErrorInvoke()
	sdk.SetErrorInvoke(errH)

	println(" * CloseConnection!")
	_ = sdk.CloseConnection()

	println("Goodbye!")

}

func printHeader(msg string) {
	println()
	println("=====================")
	println("Start: " + msg)
	println("=====================")
	println()
}

func TestWebSocket() {
	println("test in go-stream project!")
}
