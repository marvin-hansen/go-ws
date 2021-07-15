package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// BalanceData struct for BalanceData
type BalanceData struct {
	// Exchange currency code.
	AssetIdExchange *string `json:"asset_id_exchange,omitempty"`
	// CoinAPI currency code.
	AssetIdCoinapi *string `json:"asset_id_coinapi,omitempty"`
	// Value of the current total currency balance on the exchange.
	Balance *float64 `json:"balance,omitempty"`
	// Value of the current available currency balance on the exchange that can be used as collateral.
	Available *float64 `json:"available,omitempty"`
	// Value of the current locked currency balance by the exchange.
	Locked *float64 `json:"locked,omitempty"`
	// Source of the last modification.
	LastUpdatedBy *string `json:"last_updated_by,omitempty"`
	// Current exchange rate to the USD for the single unit of the currency.
	RateUsd *float64 `json:"rate_usd,omitempty"`
	// Value of the current total traded.
	Traded *float64 `json:"traded,omitempty"`
}

func (o BalanceData) String() string {
	return fmt.Sprintf("<BalanceData> AssetIdExchange: %v, AssetIdCoinapi: %v, Balance: %v, Available: %v, Locked: %v, LastUpdatedBy: %v, RateUsd: %v,  Traded: %v",
		o.GetAssetIdExchange(),
		o.GetAssetIdCoinapi(),
		convertFloatToString(o.GetBalance()),
		convertFloatToString(o.GetAvailable()),
		convertFloatToString(o.GetLocked()),
		o.GetLastUpdatedBy(),
		convertFloatToString(o.GetRateUsd()),
		convertFloatToString(o.GetTraded()),
	)
}

func convertFloatToString(f float64) (s string) {
	s = strconv.FormatFloat(float64(f), 'f', 8, 64)
	//s = fmt.Sprintf("%.8f", f)
	return s
}

// NewBalanceData instantiates a new BalanceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceData() *BalanceData {
	this := BalanceData{}
	return &this
}

// NewBalanceDataWithDefaults instantiates a new BalanceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceDataWithDefaults() *BalanceData {
	this := BalanceData{}
	return &this
}

// GetAssetIdExchange returns the AssetIdExchange field value if set, zero value otherwise.
func (o *BalanceData) GetAssetIdExchange() string {
	if o == nil || o.AssetIdExchange == nil {
		var ret string = ""
		return ret
	}
	return *o.AssetIdExchange
}

// GetAssetIdExchangeOk returns a tuple with the AssetIdExchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceData) GetAssetIdExchangeOk() (*string, bool) {
	if o == nil || o.AssetIdExchange == nil {
		return nil, false
	}
	return o.AssetIdExchange, true
}

// HasAssetIdExchange returns a boolean if a field has been set.
func (o *BalanceData) HasAssetIdExchange() bool {
	if o != nil && o.AssetIdExchange != nil {
		return true
	}

	return false
}

// SetAssetIdExchange gets a reference to the given string and assigns it to the AssetIdExchange field.
func (o *BalanceData) SetAssetIdExchange(v string) {
	o.AssetIdExchange = &v
}

// GetAssetIdCoinapi returns the AssetIdCoinapi field value if set, zero value otherwise.
func (o *BalanceData) GetAssetIdCoinapi() string {
	if o == nil || o.AssetIdCoinapi == nil {
		var ret string = ""
		return ret
	}
	return *o.AssetIdCoinapi
}

// GetAssetIdCoinapiOk returns a tuple with the AssetIdCoinapi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceData) GetAssetIdCoinapiOk() (*string, bool) {
	if o == nil || o.AssetIdCoinapi == nil {
		return nil, false
	}
	return o.AssetIdCoinapi, true
}

// HasAssetIdCoinapi returns a boolean if a field has been set.
func (o *BalanceData) HasAssetIdCoinapi() bool {
	if o != nil && o.AssetIdCoinapi != nil {
		return true
	}

	return false
}

// SetAssetIdCoinapi gets a reference to the given string and assigns it to the AssetIdCoinapi field.
func (o *BalanceData) SetAssetIdCoinapi(v string) {
	o.AssetIdCoinapi = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *BalanceData) GetBalance() float64 {
	if o == nil || o.Balance == nil {
		var ret float64 = 0.0
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceData) GetBalanceOk() (*float64, bool) {
	if o == nil || o.Balance == nil {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *BalanceData) HasBalance() bool {
	if o != nil && o.Balance != nil {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float64 and assigns it to the Balance field.
func (o *BalanceData) SetBalance(v float64) {
	o.Balance = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *BalanceData) GetAvailable() float64 {
	if o == nil || o.Available == nil {
		var ret float64 = 0.0
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceData) GetAvailableOk() (*float64, bool) {
	if o == nil || o.Available == nil {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *BalanceData) HasAvailable() bool {
	if o != nil && o.Available != nil {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given float64 and assigns it to the Available field.
func (o *BalanceData) SetAvailable(v float64) {
	o.Available = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *BalanceData) GetLocked() float64 {
	if o == nil || o.Locked == nil {
		var ret float64 = 0.0
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceData) GetLockedOk() (*float64, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *BalanceData) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given float64 and assigns it to the Locked field.
func (o *BalanceData) SetLocked(v float64) {
	o.Locked = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *BalanceData) GetLastUpdatedBy() string {
	if o == nil || o.LastUpdatedBy == nil {
		var ret string = ""
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceData) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *BalanceData) HasLastUpdatedBy() bool {
	if o != nil && o.LastUpdatedBy != nil {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *BalanceData) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetRateUsd returns the RateUsd field value if set, zero value otherwise.
func (o *BalanceData) GetRateUsd() float64 {
	if o == nil || o.RateUsd == nil {
		var ret float64 = 0.0
		return ret
	}
	return *o.RateUsd
}

// GetRateUsdOk returns a tuple with the RateUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceData) GetRateUsdOk() (*float64, bool) {
	if o == nil || o.RateUsd == nil {
		return nil, false
	}
	return o.RateUsd, true
}

// HasRateUsd returns a boolean if a field has been set.
func (o *BalanceData) HasRateUsd() bool {
	if o != nil && o.RateUsd != nil {
		return true
	}

	return false
}

// SetRateUsd gets a reference to the given float64 and assigns it to the RateUsd field.
func (o *BalanceData) SetRateUsd(v float64) {
	o.RateUsd = &v
}

// GetTraded returns the Traded field value if set, zero value otherwise.
func (o *BalanceData) GetTraded() float64 {
	if o == nil || o.Traded == nil {
		var ret float64 = 0.0
		return ret
	}
	return *o.Traded
}

// GetTradedOk returns a tuple with the Traded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceData) GetTradedOk() (*float64, bool) {
	if o == nil || o.Traded == nil {
		return nil, false
	}
	return o.Traded, true
}

// HasTraded returns a boolean if a field has been set.
func (o *BalanceData) HasTraded() bool {
	if o != nil && o.Traded != nil {
		return true
	}

	return false
}

// SetTraded gets a reference to the given float64 and assigns it to the Traded field.
func (o *BalanceData) SetTraded(v float64) {
	o.Traded = &v
}

func (o BalanceData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetIdExchange != nil {
		toSerialize["asset_id_exchange"] = o.AssetIdExchange
	}
	if o.AssetIdCoinapi != nil {
		toSerialize["asset_id_coinapi"] = o.AssetIdCoinapi
	}
	if o.Balance != nil {
		toSerialize["balance"] = o.Balance
	}
	if o.Available != nil {
		toSerialize["available"] = o.Available
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.LastUpdatedBy != nil {
		toSerialize["last_updated_by"] = o.LastUpdatedBy
	}
	if o.RateUsd != nil {
		toSerialize["rate_usd"] = o.RateUsd
	}
	if o.Traded != nil {
		toSerialize["traded"] = o.Traded
	}
	return json.Marshal(toSerialize)
}

type NullableBalanceData struct {
	value *BalanceData
	isSet bool
}

func (v NullableBalanceData) Get() *BalanceData {
	return v.value
}

func (v *NullableBalanceData) Set(val *BalanceData) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceData) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceData(val *BalanceData) *NullableBalanceData {
	return &NullableBalanceData{value: val, isSet: true}
}

func (v NullableBalanceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
