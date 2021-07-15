package types

import (
	"encoding/json"
)

// OrderExecutionReportAllOf The order execution report message.
type OrderExecutionReportAllOf struct {
	// The unique identifier of the order assigned by the client converted to the exchange order tag format for the purpose of tracking it.
	ClientOrderIdFormatExchange string `json:"client_order_id_format_exchange"`
	// Unique identifier of the order assigned by the exchange or executing system.
	ExchangeOrderId *string `json:"exchange_order_id,omitempty"`
	// Quantity open for further execution. `amount_open` = `amount_order` - `amount_filled`
	AmountOpen float64 `json:"amount_open"`
	// Total quantity filled.
	AmountFilled float64 `json:"amount_filled"`
	// Calculated average price of all fills on this order.
	AvgPx  *float64  `json:"avg_px,omitempty"`
	Status OrdStatus `json:"status"`
	// Timestamped history of order status changes.
	StatusHistory *[][]string `json:"status_history,omitempty"`
	// Error message.
	ErrorMessage *string `json:"error_message,omitempty"`
	// Relay fill information on working orders.
	Fills *[]Fills `json:"fills,omitempty"`
}

// NewOrderExecutionReportAllOf instantiates a new OrderExecutionReportAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderExecutionReportAllOf(clientOrderIdFormatExchange string, amountOpen float64, amountFilled float64, status OrdStatus) *OrderExecutionReportAllOf {
	this := OrderExecutionReportAllOf{}
	this.ClientOrderIdFormatExchange = clientOrderIdFormatExchange
	this.AmountOpen = amountOpen
	this.AmountFilled = amountFilled
	this.Status = status
	return &this
}

// NewOrderExecutionReportAllOfWithDefaults instantiates a new OrderExecutionReportAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderExecutionReportAllOfWithDefaults() *OrderExecutionReportAllOf {
	this := OrderExecutionReportAllOf{}
	return &this
}

// GetClientOrderIdFormatExchange returns the ClientOrderIdFormatExchange field value
func (o *OrderExecutionReportAllOf) GetClientOrderIdFormatExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientOrderIdFormatExchange
}

// GetClientOrderIdFormatExchangeOk returns a tuple with the ClientOrderIdFormatExchange field value
// and a boolean to check if the value has been set.
func (o *OrderExecutionReportAllOf) GetClientOrderIdFormatExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientOrderIdFormatExchange, true
}

// SetClientOrderIdFormatExchange sets field value
func (o *OrderExecutionReportAllOf) SetClientOrderIdFormatExchange(v string) {
	o.ClientOrderIdFormatExchange = v
}

// GetExchangeOrderId returns the ExchangeOrderId field value if set, zero value otherwise.
func (o *OrderExecutionReportAllOf) GetExchangeOrderId() string {
	if o == nil || o.ExchangeOrderId == nil {
		var ret string
		return ret
	}
	return *o.ExchangeOrderId
}

// GetExchangeOrderIdOk returns a tuple with the ExchangeOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderExecutionReportAllOf) GetExchangeOrderIdOk() (*string, bool) {
	if o == nil || o.ExchangeOrderId == nil {
		return nil, false
	}
	return o.ExchangeOrderId, true
}

// HasExchangeOrderId returns a boolean if a field has been set.
func (o *OrderExecutionReportAllOf) HasExchangeOrderId() bool {
	if o != nil && o.ExchangeOrderId != nil {
		return true
	}

	return false
}

// SetExchangeOrderId gets a reference to the given string and assigns it to the ExchangeOrderId field.
func (o *OrderExecutionReportAllOf) SetExchangeOrderId(v string) {
	o.ExchangeOrderId = &v
}

// GetAmountOpen returns the AmountOpen field value
func (o *OrderExecutionReportAllOf) GetAmountOpen() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AmountOpen
}

// GetAmountOpenOk returns a tuple with the AmountOpen field value
// and a boolean to check if the value has been set.
func (o *OrderExecutionReportAllOf) GetAmountOpenOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountOpen, true
}

// SetAmountOpen sets field value
func (o *OrderExecutionReportAllOf) SetAmountOpen(v float64) {
	o.AmountOpen = v
}

// GetAmountFilled returns the AmountFilled field value
func (o *OrderExecutionReportAllOf) GetAmountFilled() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AmountFilled
}

// GetAmountFilledOk returns a tuple with the AmountFilled field value
// and a boolean to check if the value has been set.
func (o *OrderExecutionReportAllOf) GetAmountFilledOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountFilled, true
}

// SetAmountFilled sets field value
func (o *OrderExecutionReportAllOf) SetAmountFilled(v float64) {
	o.AmountFilled = v
}

// GetAvgPx returns the AvgPx field value if set, zero value otherwise.
func (o *OrderExecutionReportAllOf) GetAvgPx() float64 {
	if o == nil || o.AvgPx == nil {
		var ret float64
		return ret
	}
	return *o.AvgPx
}

// GetAvgPxOk returns a tuple with the AvgPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderExecutionReportAllOf) GetAvgPxOk() (*float64, bool) {
	if o == nil || o.AvgPx == nil {
		return nil, false
	}
	return o.AvgPx, true
}

// HasAvgPx returns a boolean if a field has been set.
func (o *OrderExecutionReportAllOf) HasAvgPx() bool {
	if o != nil && o.AvgPx != nil {
		return true
	}

	return false
}

// SetAvgPx gets a reference to the given float64 and assigns it to the AvgPx field.
func (o *OrderExecutionReportAllOf) SetAvgPx(v float64) {
	o.AvgPx = &v
}

// GetStatus returns the Status field value
func (o *OrderExecutionReportAllOf) GetStatus() OrdStatus {
	if o == nil {
		var ret OrdStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrderExecutionReportAllOf) GetStatusOk() (*OrdStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrderExecutionReportAllOf) SetStatus(v OrdStatus) {
	o.Status = v
}

// GetStatusHistory returns the StatusHistory field value if set, zero value otherwise.
func (o *OrderExecutionReportAllOf) GetStatusHistory() [][]string {
	if o == nil || o.StatusHistory == nil {
		var ret [][]string
		return ret
	}
	return *o.StatusHistory
}

// GetStatusHistoryOk returns a tuple with the StatusHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderExecutionReportAllOf) GetStatusHistoryOk() (*[][]string, bool) {
	if o == nil || o.StatusHistory == nil {
		return nil, false
	}
	return o.StatusHistory, true
}

// HasStatusHistory returns a boolean if a field has been set.
func (o *OrderExecutionReportAllOf) HasStatusHistory() bool {
	if o != nil && o.StatusHistory != nil {
		return true
	}

	return false
}

// SetStatusHistory gets a reference to the given [][]string and assigns it to the StatusHistory field.
func (o *OrderExecutionReportAllOf) SetStatusHistory(v [][]string) {
	o.StatusHistory = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *OrderExecutionReportAllOf) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderExecutionReportAllOf) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *OrderExecutionReportAllOf) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *OrderExecutionReportAllOf) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetFills returns the Fills field value if set, zero value otherwise.
func (o *OrderExecutionReportAllOf) GetFills() []Fills {
	if o == nil || o.Fills == nil {
		var ret []Fills
		return ret
	}
	return *o.Fills
}

// GetFillsOk returns a tuple with the Fills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderExecutionReportAllOf) GetFillsOk() (*[]Fills, bool) {
	if o == nil || o.Fills == nil {
		return nil, false
	}
	return o.Fills, true
}

// HasFills returns a boolean if a field has been set.
func (o *OrderExecutionReportAllOf) HasFills() bool {
	if o != nil && o.Fills != nil {
		return true
	}

	return false
}

// SetFills gets a reference to the given []Fills and assigns it to the Fills field.
func (o *OrderExecutionReportAllOf) SetFills(v []Fills) {
	o.Fills = &v
}

func (o OrderExecutionReportAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_order_id_format_exchange"] = o.ClientOrderIdFormatExchange
	}
	if o.ExchangeOrderId != nil {
		toSerialize["exchange_order_id"] = o.ExchangeOrderId
	}
	if true {
		toSerialize["amount_open"] = o.AmountOpen
	}
	if true {
		toSerialize["amount_filled"] = o.AmountFilled
	}
	if o.AvgPx != nil {
		toSerialize["avg_px"] = o.AvgPx
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.StatusHistory != nil {
		toSerialize["status_history"] = o.StatusHistory
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.Fills != nil {
		toSerialize["fills"] = o.Fills
	}
	return json.Marshal(toSerialize)
}

type NullableOrderExecutionReportAllOf struct {
	value *OrderExecutionReportAllOf
	isSet bool
}

func (v NullableOrderExecutionReportAllOf) Get() *OrderExecutionReportAllOf {
	return v.value
}

func (v *NullableOrderExecutionReportAllOf) Set(val *OrderExecutionReportAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderExecutionReportAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderExecutionReportAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderExecutionReportAllOf(val *OrderExecutionReportAllOf) *NullableOrderExecutionReportAllOf {
	return &NullableOrderExecutionReportAllOf{value: val, isSet: true}
}

func (v NullableOrderExecutionReportAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderExecutionReportAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
