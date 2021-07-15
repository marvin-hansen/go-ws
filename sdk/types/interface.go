package types

type SDK interface {
	OpenConnection() (err error)
	CloseConnection() (err error)
	ResetConnection() (err error)
	//
	LookupSymbolData(exchangeID, baseSymbolCoinApi, quoteSymbolCoinApi string) (ok bool, symbolData SymbolData)
	//
	SetSystemInvoke(function SystemInvoke)
	SetSnapshotInvoke(function SnapshotInvoke)
	SetUpdateInvoke(function UpdateInvoke)
	//
	SetExecUpdateInvoke(function InvokeFunction)
	SetExecSnapshotInvoke(function InvokeFunction)
	SetBalanceUpdateInvoke(function InvokeFunction)
	SetBalanceSnapshotInvoke(function InvokeFunction)
	SetPositionUpdateInvoke(function InvokeFunction)
	SetPositionSnapshotInvoke(function InvokeFunction)
}
