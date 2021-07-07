package types

import (
	"encoding/json"
)

// Position struct for Position
type Position struct {
	// Exchange identifier used to identify the routing destination.
	ExchangeId *string         `json:"exchange_id,omitempty"`
	Data       *[]PositionData `json:"data,omitempty"`
}

// NewPosition instantiates a new Position object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPosition() *Position {
	this := Position{}
	return &this
}

// NewPositionWithDefaults instantiates a new Position object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionWithDefaults() *Position {
	this := Position{}
	return &this
}

// GetExchangeId returns the ExchangeId field value if set, zero value otherwise.
func (o *Position) GetExchangeId() string {
	if o == nil || o.ExchangeId == nil {
		var ret string
		return ret
	}
	return *o.ExchangeId
}

// GetExchangeIdOk returns a tuple with the ExchangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Position) GetExchangeIdOk() (*string, bool) {
	if o == nil || o.ExchangeId == nil {
		return nil, false
	}
	return o.ExchangeId, true
}

// HasExchangeId returns a boolean if a field has been set.
func (o *Position) HasExchangeId() bool {
	if o != nil && o.ExchangeId != nil {
		return true
	}

	return false
}

// SetExchangeId gets a reference to the given string and assigns it to the ExchangeId field.
func (o *Position) SetExchangeId(v string) {
	o.ExchangeId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Position) GetData() []PositionData {
	if o == nil || o.Data == nil {
		var ret []PositionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Position) GetDataOk() (*[]PositionData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Position) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []PositionData and assigns it to the Data field.
func (o *Position) SetData(v []PositionData) {
	o.Data = &v
}

func (o Position) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExchangeId != nil {
		toSerialize["exchange_id"] = o.ExchangeId
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePosition struct {
	value *Position
	isSet bool
}

func (v NullablePosition) Get() *Position {
	return v.value
}

func (v *NullablePosition) Set(val *Position) {
	v.value = val
	v.isSet = true
}

func (v NullablePosition) IsSet() bool {
	return v.isSet
}

func (v *NullablePosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePosition(val *Position) *NullablePosition {
	return &NullablePosition{value: val, isSet: true}
}

func (v NullablePosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
