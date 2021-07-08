package v1

import (
	"go-ws/sdk/types"
	"go-ws/sdk/web_socket"
	"log"
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

func (s SDKImpl) OpenConnection() (err error) {
	url := *s.url
	err = s.ws.Connect(url, nil)
	return err
}

func (s SDKImpl) CloseConnection() (err error) {
	running = false // Stop processing messages
	err = s.ws.Close()
	if err != nil {
		log.Println("Can't close connection")
		log.Println(err)
	}
	return err
}

func (s SDKImpl) ResetConnection() (err error) {
	err = s.CloseConnection()
	logError(err)

	err = s.OpenConnection()
	logError(err)
	return err
}
