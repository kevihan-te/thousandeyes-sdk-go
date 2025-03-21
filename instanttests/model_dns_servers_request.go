/*
Instant Tests API

The Instant Tests API operations lets you create and run new instant tests. You will need to be a regular user or have the following permissions:   * `API Access`   * `View tests`  The response does not include the immediate test results. Use the Test Results endpoints to get test results after creating and executing an instant test. You can find the URLs for these endpoints in the _links section of the test definition that is returned when you create the instant test. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instanttests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the DnsServersRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DnsServersRequest{}

// DnsServersRequest struct for DnsServersRequest
type DnsServersRequest struct {
	// A list of DNS server FQDN.
	DnsServers []string `json:"dnsServers,omitempty"`
}

// NewDnsServersRequest instantiates a new DnsServersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsServersRequest() *DnsServersRequest {
	this := DnsServersRequest{}
	return &this
}

// NewDnsServersRequestWithDefaults instantiates a new DnsServersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsServersRequestWithDefaults() *DnsServersRequest {
	this := DnsServersRequest{}
	return &this
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise.
func (o *DnsServersRequest) GetDnsServers() []string {
	if o == nil || utils.IsNil(o.DnsServers) {
		var ret []string
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServersRequest) GetDnsServersOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.DnsServers) {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *DnsServersRequest) HasDnsServers() bool {
	if o != nil && !utils.IsNil(o.DnsServers) {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *DnsServersRequest) SetDnsServers(v []string) {
	o.DnsServers = v
}

func (o DnsServersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsServersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.DnsServers) {
		toSerialize["dnsServers"] = o.DnsServers
	}
	return toSerialize, nil
}

type NullableDnsServersRequest struct {
	value *DnsServersRequest
	isSet bool
}

func (v NullableDnsServersRequest) Get() *DnsServersRequest {
	return v.value
}

func (v *NullableDnsServersRequest) Set(val *DnsServersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsServersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsServersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsServersRequest(val *DnsServersRequest) *NullableDnsServersRequest {
	return &NullableDnsServersRequest{value: val, isSet: true}
}

func (v NullableDnsServersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsServersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


