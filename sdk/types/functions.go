package types

// SystemInvoke just bundles system invoke functions
type SystemInvoke struct {
	ErrorMessageInvoke   InvokeFunction
	ServerInfoInvoke     InvokeFunction
	SymbolSnapshotInvoke InvokeFunction
}

// InvokeFunction is a unified function type for all event handlers.
// https://yourbasic.org/golang/function-pointer-type-declaration/
type InvokeFunction func(message *DataMessage) (err error)

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// WsErrHandler handles raw websocket errors
type WsErrHandler func(err error)
