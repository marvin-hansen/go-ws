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
	errorMessageInvoke   types.InvokeFunction
	serverInfoInvoke     types.InvokeFunction
	symbolSnapshotInvoke types.InvokeFunction

	executionUpdateInvoke   types.InvokeFunction
	executionSnapshotInvoke types.InvokeFunction
	balanceUpdateInvoke     types.InvokeFunction
	balanceSnapshotInvoke   types.InvokeFunction
	positionUpdateInvoke    types.InvokeFunction
	positionSnapshotInvoke  types.InvokeFunction
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
	err := s.StartMessageProcessing()
	if err != nil {
		logError(err)
		panic(err)
	}
}

func (s SDKImpl) OpenConnection() (err error) {
	url := *s.url
	err = s.ws.Connect(url, nil)
	return err
}

func (s SDKImpl) CloseConnection() (err error) {

	s.StopMessageProcessing()
	err = s.ws.Close()
	if err != nil {
		log.Println("Can't close connection")
		log.Println(err)
	}
	return err
}

func (s SDKImpl) ResetConnection() (err error) {

	s.StopMessageProcessing()

	err = s.CloseConnection()
	logError(err)

	err = s.OpenConnection()
	logError(err)

	err = s.StartMessageProcessing()
	logError(err)

	return err
}
