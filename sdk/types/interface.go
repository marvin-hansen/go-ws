package types

type SDK interface {
	OpenConnection() (err error)
	CloseConnection() (err error)
	ResetConnection() (err error)
	//
	SetErrorInvoke(function InvokeFunction)
	SetHeartBeatInvoke(function InvokeFunction)
	SetReconnectInvoke(function InvokeFunction)
	//
}
