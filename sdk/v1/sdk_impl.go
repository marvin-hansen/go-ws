package v1

import (
	t "go-ws/sdk/types"
	"go-ws/sdk/web_socket"
	"log"
)

type SDKImpl struct {
	url *string
	ws  *web_socket.WebSocket
}

// AssetKey is a composite key for a hashmap to access symbol description
// https://stackoverflow.com/questions/52348514/how-to-make-composite-key-for-a-hash-map-in-golang
type AssetKey struct {
	Base, Quote string
}

var (
	errorMessageInvoke   t.InvokeFunction
	serverInfoInvoke     t.InvokeFunction
	symbolSnapshotInvoke t.InvokeFunction

	executionUpdateInvoke   t.InvokeFunction
	executionSnapshotInvoke t.InvokeFunction
	balanceUpdateInvoke     t.InvokeFunction
	balanceSnapshotInvoke   t.InvokeFunction
	positionUpdateInvoke    t.InvokeFunction
	positionSnapshotInvoke  t.InvokeFunction
	//
	symbolMap map[string]map[AssetKey]t.SymbolData
)

func NewOemlSdkV1(url string) (sdk *SDKImpl) {
	sdk = new(SDKImpl)
	sdk.init(url)
	return sdk
}

func (s SDKImpl) init(url string) {
	// init AssetMap
	symbolMap = make(map[string]map[AssetKey]t.SymbolData)

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
