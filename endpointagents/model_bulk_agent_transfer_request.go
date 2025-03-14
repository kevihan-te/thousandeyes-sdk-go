/*
Endpoint Agents API

Manage ThousandEyes Endpoint Agents using this API.   For more information about Endpoint Agents, see [Endpoint Agents](https://docs.thousandeyes.com/product-documentation/global-vantage-points/endpoint-agents).

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointagents

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the BulkAgentTransferRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BulkAgentTransferRequest{}

// BulkAgentTransferRequest struct for BulkAgentTransferRequest
type BulkAgentTransferRequest struct {
	Transfers []AgentTransfer `json:"transfers,omitempty"`
}

// NewBulkAgentTransferRequest instantiates a new BulkAgentTransferRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkAgentTransferRequest() *BulkAgentTransferRequest {
	this := BulkAgentTransferRequest{}
	return &this
}

// NewBulkAgentTransferRequestWithDefaults instantiates a new BulkAgentTransferRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkAgentTransferRequestWithDefaults() *BulkAgentTransferRequest {
	this := BulkAgentTransferRequest{}
	return &this
}

// GetTransfers returns the Transfers field value if set, zero value otherwise.
func (o *BulkAgentTransferRequest) GetTransfers() []AgentTransfer {
	if o == nil || utils.IsNil(o.Transfers) {
		var ret []AgentTransfer
		return ret
	}
	return o.Transfers
}

// GetTransfersOk returns a tuple with the Transfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAgentTransferRequest) GetTransfersOk() ([]AgentTransfer, bool) {
	if o == nil || utils.IsNil(o.Transfers) {
		return nil, false
	}
	return o.Transfers, true
}

// HasTransfers returns a boolean if a field has been set.
func (o *BulkAgentTransferRequest) HasTransfers() bool {
	if o != nil && !utils.IsNil(o.Transfers) {
		return true
	}

	return false
}

// SetTransfers gets a reference to the given []AgentTransfer and assigns it to the Transfers field.
func (o *BulkAgentTransferRequest) SetTransfers(v []AgentTransfer) {
	o.Transfers = v
}

func (o BulkAgentTransferRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkAgentTransferRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Transfers) {
		toSerialize["transfers"] = o.Transfers
	}
	return toSerialize, nil
}

type NullableBulkAgentTransferRequest struct {
	value *BulkAgentTransferRequest
	isSet bool
}

func (v NullableBulkAgentTransferRequest) Get() *BulkAgentTransferRequest {
	return v.value
}

func (v *NullableBulkAgentTransferRequest) Set(val *BulkAgentTransferRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkAgentTransferRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkAgentTransferRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkAgentTransferRequest(val *BulkAgentTransferRequest) *NullableBulkAgentTransferRequest {
	return &NullableBulkAgentTransferRequest{value: val, isSet: true}
}

func (v NullableBulkAgentTransferRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkAgentTransferRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


