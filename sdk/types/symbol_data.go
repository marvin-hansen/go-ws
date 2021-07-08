package types

type SymbolData struct {
	Symbol_id_coinapi       *string  `json:"symbol_id_coinapi,omitempty"`
	Symbol_id_base_exchange *string  `json:"symbol_id_exchange,omitempty"`
	Asset_id_base_exchange  *string  `json:"asset_id_base_exchange,omitempty"`
	Asset_id_quote_exchange *string  `json:"asset_id_quote_exchange,omitempty"`
	Asset_id_base_coinapi   *string  `json:"asset_id_base_coinapi,omitempty"`
	Asset_id_quote_coinapi  *string  `json:"asset_id_quote_coinapi,omitempty"`
	Price_precision         *float32 `json:"price_precision,omitempty"`
	Size_precision          *float32 `json:"size_precision,omitempty"`
}

func NewSymbolData() (this *SymbolData) {
	this = new(SymbolData)
	return this
}
