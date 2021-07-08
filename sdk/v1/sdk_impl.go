package v1

import (
	"go-ws/sdk/types"
	"go-ws/sdk/web_socket"
)

type SDKImpl struct {
	url *string
	ws  *web_socket.WebSocket
}

var (
	running         bool
	errorInvoke     types.InvokeFunction
	heartBeatInvoke types.InvokeFunction
	reconnectInvoke types.InvokeFunction
)

func NewOemlSdkV1(url string) (sdk *SDKImpl) {
	sdk = new(SDKImpl)
	sdk.init(url)
	return sdk
}

func (s SDKImpl) init(url string) {
	ws := web_socket.NewWebSocket(url)
	s.url = &url
	s.ws = ws
}
