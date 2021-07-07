package v1

import (
	"github.com/gorilla/websocket"
	"go-ws/sdk/types"
)

type SDKImpl struct {
}

var (
	con     *websocket.Conn
	stopC   chan struct{}
	running bool
	//
	// sys handlers
	errorInvoke     types.InvokeFunction
	heartBeatInvoke types.InvokeFunction
	reconnectInvoke types.InvokeFunction
)

func NewOemlSdkV1() (sdk *SDKImpl) {
	sdk = new(SDKImpl)
	sdk.init()
	return sdk
}

func (s SDKImpl) init() {

}
