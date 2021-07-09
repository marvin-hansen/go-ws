package main

import (
	SDK "go-ws/sdk"
)

const url = "ws://"

func main() {
	TestConnection()
}

func TestConnection() {
	printHeader("TestConnection!")

	println(" * NewSDK!")
	sdk := SDK.NewSDK(url)

	println(" * SetSysInvoke!")
	sysInvoke := GetSysInvoke()
	sdk.SetSystemInvoke(sysInvoke)

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
