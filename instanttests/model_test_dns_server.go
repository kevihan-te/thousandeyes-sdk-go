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

// checks if the TestDnsServer type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TestDnsServer{}

// TestDnsServer struct for TestDnsServer
type TestDnsServer struct {
	// Unique identifier of the DNS server.
	ServerId *string `json:"serverId,omitempty"`
	// Fully qualified domain name (FQDN) of DNS server.
	ServerName *string `json:"serverName,omitempty"`
}

// NewTestDnsServer instantiates a new TestDnsServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestDnsServer() *TestDnsServer {
	this := TestDnsServer{}
	return &this
}

// NewTestDnsServerWithDefaults instantiates a new TestDnsServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestDnsServerWithDefaults() *TestDnsServer {
	this := TestDnsServer{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *TestDnsServer) GetServerId() string {
	if o == nil || utils.IsNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestDnsServer) GetServerIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *TestDnsServer) HasServerId() bool {
	if o != nil && !utils.IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *TestDnsServer) SetServerId(v string) {
	o.ServerId = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *TestDnsServer) GetServerName() string {
	if o == nil || utils.IsNil(o.ServerName) {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestDnsServer) GetServerNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ServerName) {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *TestDnsServer) HasServerName() bool {
	if o != nil && !utils.IsNil(o.ServerName) {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *TestDnsServer) SetServerName(v string) {
	o.ServerName = &v
}

func (o TestDnsServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestDnsServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if !utils.IsNil(o.ServerName) {
		toSerialize["serverName"] = o.ServerName
	}
	return toSerialize, nil
}

type NullableTestDnsServer struct {
	value *TestDnsServer
	isSet bool
}

func (v NullableTestDnsServer) Get() *TestDnsServer {
	return v.value
}

func (v *NullableTestDnsServer) Set(val *TestDnsServer) {
	v.value = val
	v.isSet = true
}

func (v NullableTestDnsServer) IsSet() bool {
	return v.isSet
}

func (v *NullableTestDnsServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestDnsServer(val *TestDnsServer) *NullableTestDnsServer {
	return &NullableTestDnsServer{value: val, isSet: true}
}

func (v NullableTestDnsServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestDnsServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


