package types

import (
	"encoding/json"
)

type OrderCancelAllRequest struct {
	// Message type to identity the request
	MessageType MessageType `json:"type"`
	// Identifier of the exchange from which active orders should be canceled.
	ExchangeId string `json:"exchange_id"`
}

func (o OrderCancelAllRequest) MarshalJSON() (b []byte, err error) {

	toSerialize := map[string]string{}
	toSerialize["type"] = string(o.MessageType)
	toSerialize["exchange_id"] = o.ExchangeId

	b, err = json.Marshal(toSerialize)

	return b, err
}
