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

// checks if the ConnectionString type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ConnectionString{}

// ConnectionString struct for ConnectionString
type ConnectionString struct {
	// The connection string is used for some integrations and other client types. 
	ConnectionString *string `json:"connectionString,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewConnectionString instantiates a new ConnectionString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionString() *ConnectionString {
	this := ConnectionString{}
	return &this
}

// NewConnectionStringWithDefaults instantiates a new ConnectionString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionStringWithDefaults() *ConnectionString {
	this := ConnectionString{}
	return &this
}

// GetConnectionString returns the ConnectionString field value if set, zero value otherwise.
func (o *ConnectionString) GetConnectionString() string {
	if o == nil || utils.IsNil(o.ConnectionString) {
		var ret string
		return ret
	}
	return *o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionString) GetConnectionStringOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ConnectionString) {
		return nil, false
	}
	return o.ConnectionString, true
}

// HasConnectionString returns a boolean if a field has been set.
func (o *ConnectionString) HasConnectionString() bool {
	if o != nil && !utils.IsNil(o.ConnectionString) {
		return true
	}

	return false
}

// SetConnectionString gets a reference to the given string and assigns it to the ConnectionString field.
func (o *ConnectionString) SetConnectionString(v string) {
	o.ConnectionString = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ConnectionString) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionString) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ConnectionString) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *ConnectionString) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o ConnectionString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ConnectionString) {
		toSerialize["connectionString"] = o.ConnectionString
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableConnectionString struct {
	value *ConnectionString
	isSet bool
}

func (v NullableConnectionString) Get() *ConnectionString {
	return v.value
}

func (v *NullableConnectionString) Set(val *ConnectionString) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionString) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionString(val *ConnectionString) *NullableConnectionString {
	return &NullableConnectionString{value: val, isSet: true}
}

func (v NullableConnectionString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


