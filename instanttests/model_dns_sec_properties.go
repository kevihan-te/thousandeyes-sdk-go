/*
Instant Tests API

The Instant Tests API operations lets you create and run new instant tests. You will need to be a regular user or have the following permissions:   * `API Access`   * `View tests`  The response does not include the immediate test results. Use the Test Results endpoints to get test results after creating and executing an instant test. You can find the URLs for these endpoints in the _links section of the test definition that is returned when you create the instant test. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instanttests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"fmt"
)

// checks if the DnsSecProperties type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DnsSecProperties{}

// DnsSecProperties struct for DnsSecProperties
type DnsSecProperties struct {
	// The target record for the test, with the record type suffixed. If no record type is specified, the test defaults to an ANY record.
	Domain string `json:"domain"`
	DnsQueryClass *DnsQueryClass `json:"dnsQueryClass,omitempty"`
	// Indicates whether agents should randomize the start time in each test round.
	RandomizedStartTime *bool `json:"randomizedStartTime,omitempty"`
	Type *string `json:"type,omitempty"`
}

type _DnsSecProperties DnsSecProperties

// NewDnsSecProperties instantiates a new DnsSecProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsSecProperties(domain string) *DnsSecProperties {
	this := DnsSecProperties{}
	this.Domain = domain
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	return &this
}

// NewDnsSecPropertiesWithDefaults instantiates a new DnsSecProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsSecPropertiesWithDefaults() *DnsSecProperties {
	this := DnsSecProperties{}
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	return &this
}

// GetDomain returns the Domain field value
func (o *DnsSecProperties) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DnsSecProperties) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *DnsSecProperties) SetDomain(v string) {
	o.Domain = v
}

// GetDnsQueryClass returns the DnsQueryClass field value if set, zero value otherwise.
func (o *DnsSecProperties) GetDnsQueryClass() DnsQueryClass {
	if o == nil || utils.IsNil(o.DnsQueryClass) {
		var ret DnsQueryClass
		return ret
	}
	return *o.DnsQueryClass
}

// GetDnsQueryClassOk returns a tuple with the DnsQueryClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsSecProperties) GetDnsQueryClassOk() (*DnsQueryClass, bool) {
	if o == nil || utils.IsNil(o.DnsQueryClass) {
		return nil, false
	}
	return o.DnsQueryClass, true
}

// HasDnsQueryClass returns a boolean if a field has been set.
func (o *DnsSecProperties) HasDnsQueryClass() bool {
	if o != nil && !utils.IsNil(o.DnsQueryClass) {
		return true
	}

	return false
}

// SetDnsQueryClass gets a reference to the given DnsQueryClass and assigns it to the DnsQueryClass field.
func (o *DnsSecProperties) SetDnsQueryClass(v DnsQueryClass) {
	o.DnsQueryClass = &v
}

// GetRandomizedStartTime returns the RandomizedStartTime field value if set, zero value otherwise.
func (o *DnsSecProperties) GetRandomizedStartTime() bool {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		var ret bool
		return ret
	}
	return *o.RandomizedStartTime
}

// GetRandomizedStartTimeOk returns a tuple with the RandomizedStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsSecProperties) GetRandomizedStartTimeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		return nil, false
	}
	return o.RandomizedStartTime, true
}

// HasRandomizedStartTime returns a boolean if a field has been set.
func (o *DnsSecProperties) HasRandomizedStartTime() bool {
	if o != nil && !utils.IsNil(o.RandomizedStartTime) {
		return true
	}

	return false
}

// SetRandomizedStartTime gets a reference to the given bool and assigns it to the RandomizedStartTime field.
func (o *DnsSecProperties) SetRandomizedStartTime(v bool) {
	o.RandomizedStartTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DnsSecProperties) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsSecProperties) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DnsSecProperties) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DnsSecProperties) SetType(v string) {
	o.Type = &v
}

func (o DnsSecProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsSecProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	if !utils.IsNil(o.DnsQueryClass) {
		toSerialize["dnsQueryClass"] = o.DnsQueryClass
	}
	if !utils.IsNil(o.RandomizedStartTime) {
		toSerialize["randomizedStartTime"] = o.RandomizedStartTime
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *DnsSecProperties) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDnsSecProperties := _DnsSecProperties{}

    err = json.Unmarshal(data, &varDnsSecProperties)

	if err != nil {
		return err
	}

	*o = DnsSecProperties(varDnsSecProperties)

	return err
}

type NullableDnsSecProperties struct {
	value *DnsSecProperties
	isSet bool
}

func (v NullableDnsSecProperties) Get() *DnsSecProperties {
	return v.value
}

func (v *NullableDnsSecProperties) Set(val *DnsSecProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsSecProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsSecProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsSecProperties(val *DnsSecProperties) *NullableDnsSecProperties {
	return &NullableDnsSecProperties{value: val, isSet: true}
}

func (v NullableDnsSecProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsSecProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


