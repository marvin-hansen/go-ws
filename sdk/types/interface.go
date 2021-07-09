package types

type SDK interface {
	OpenConnection() (err error)
	CloseConnection() (err error)
	ResetConnection() (err error)
	//
	SetSystemInvoke(function SystemInvoke)
	SetExecUpdateInvoke(function InvokeFunction)
	SetExecSnapshotInvoke(function InvokeFunction)
	SetBalanceUpdateInvoke(function InvokeFunction)
	SetBalanceSnapshotInvoke(function InvokeFunction)
	SetPositionUpdateInvoke(function InvokeFunction)
	SetPositionSnapshotInvoke(function InvokeFunction)
}
