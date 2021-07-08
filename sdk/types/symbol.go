package types

type Symbols struct {
	// Exchange identifier used to identify the routing destination.
	ExchangeId *string       `json:"exchange_id,omitempty"`
	Data       *[]SymbolData `json:"data,omitempty"`
}

func NewSymbols() (this *Symbols) {
	this = new(Symbols)
	return this
}
