package types

import "fmt"

type SymbolData struct {
	Symbol_id_coinapi       *string  `json:"symbol_id_coinapi,omitempty"`
	Symbol_id_base_exchange *string  `json:"symbol_id_exchange,omitempty"`
	Asset_id_base_exchange  *string  `json:"asset_id_base_exchange,omitempty"`
	Asset_id_quote_exchange *string  `json:"asset_id_quote_exchange,omitempty"`
	Asset_id_base_coinapi   *string  `json:"asset_id_base_coinapi,omitempty"`
	Asset_id_quote_coinapi  *string  `json:"asset_id_quote_coinapi,omitempty"`
	Price_precision         *float64 `json:"price_precision,omitempty"`
	Size_precision          *float64 `json:"size_precision,omitempty"`
}

func (s SymbolData) String() string {
	return fmt.Sprintf("<SymbolData> Symbol_id_coinapi: %v, Symbol_id_base_exchange: %v, Asset_id_base_exchange: %v, Asset_id_quote_exchange: %v, Asset_id_base_coinapi: %v, Asset_id_quote_coinapi:  %v, Price_precision  %v, Size_precision: %v",
		*s.Symbol_id_coinapi,
		*s.Symbol_id_base_exchange,
		*s.Asset_id_base_exchange,
		*s.Asset_id_quote_exchange,
		*s.Asset_id_base_coinapi,
		*s.Asset_id_quote_coinapi,
		*s.Price_precision,
		*s.Size_precision,
	)
}
