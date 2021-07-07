package types

import (
	"encoding/json"
)

// OrderCancelAllRequest Cancel all orders request object.
type OrderCancelAllRequest struct {
	// Identifier of the exchange from which active orders should be canceled.
	ExchangeId string `json:"exchange_id"`
}

// NewOrderCancelAllRequest instantiates a new OrderCancelAllRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCancelAllRequest(exchangeId string) *OrderCancelAllRequest {
	this := OrderCancelAllRequest{}
	this.ExchangeId = exchangeId
	return &this
}

// NewOrderCancelAllRequestWithDefaults instantiates a new OrderCancelAllRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCancelAllRequestWithDefaults() *OrderCancelAllRequest {
	this := OrderCancelAllRequest{}
	return &this
}

// GetExchangeId returns the ExchangeId field value
func (o *OrderCancelAllRequest) GetExchangeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeId
}

// GetExchangeIdOk returns a tuple with the ExchangeId field value
// and a boolean to check if the value has been set.
func (o *OrderCancelAllRequest) GetExchangeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeId, true
}

// SetExchangeId sets field value
func (o *OrderCancelAllRequest) SetExchangeId(v string) {
	o.ExchangeId = v
}

func (o OrderCancelAllRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["exchange_id"] = o.ExchangeId
	}
	return json.Marshal(toSerialize)
}

type NullableOrderCancelAllRequest struct {
	value *OrderCancelAllRequest
	isSet bool
}

func (v NullableOrderCancelAllRequest) Get() *OrderCancelAllRequest {
	return v.value
}

func (v *NullableOrderCancelAllRequest) Set(val *OrderCancelAllRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelAllRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelAllRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelAllRequest(val *OrderCancelAllRequest) *NullableOrderCancelAllRequest {
	return &NullableOrderCancelAllRequest{value: val, isSet: true}
}

func (v NullableOrderCancelAllRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelAllRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
