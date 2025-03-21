/*
Endpoint Test Results API

Retrieve results for scheduled and dynamic tests on endpoint agents.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointtestresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the DynamicTestLinks type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DynamicTestLinks{}

// DynamicTestLinks A list of links that can be accessed to get more information.
type DynamicTestLinks struct {
	Self *DynamicTestSelfLink `json:"self,omitempty"`
	// Reference to the test results.
	TestResults []Link `json:"testResults,omitempty"`
}

// NewDynamicTestLinks instantiates a new DynamicTestLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicTestLinks() *DynamicTestLinks {
	this := DynamicTestLinks{}
	return &this
}

// NewDynamicTestLinksWithDefaults instantiates a new DynamicTestLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicTestLinksWithDefaults() *DynamicTestLinks {
	this := DynamicTestLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *DynamicTestLinks) GetSelf() DynamicTestSelfLink {
	if o == nil || utils.IsNil(o.Self) {
		var ret DynamicTestSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTestLinks) GetSelfOk() (*DynamicTestSelfLink, bool) {
	if o == nil || utils.IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *DynamicTestLinks) HasSelf() bool {
	if o != nil && !utils.IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given DynamicTestSelfLink and assigns it to the Self field.
func (o *DynamicTestLinks) SetSelf(v DynamicTestSelfLink) {
	o.Self = &v
}

// GetTestResults returns the TestResults field value if set, zero value otherwise.
func (o *DynamicTestLinks) GetTestResults() []Link {
	if o == nil || utils.IsNil(o.TestResults) {
		var ret []Link
		return ret
	}
	return o.TestResults
}

// GetTestResultsOk returns a tuple with the TestResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTestLinks) GetTestResultsOk() ([]Link, bool) {
	if o == nil || utils.IsNil(o.TestResults) {
		return nil, false
	}
	return o.TestResults, true
}

// HasTestResults returns a boolean if a field has been set.
func (o *DynamicTestLinks) HasTestResults() bool {
	if o != nil && !utils.IsNil(o.TestResults) {
		return true
	}

	return false
}

// SetTestResults gets a reference to the given []Link and assigns it to the TestResults field.
func (o *DynamicTestLinks) SetTestResults(v []Link) {
	o.TestResults = v
}

func (o DynamicTestLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DynamicTestLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !utils.IsNil(o.TestResults) {
		toSerialize["testResults"] = o.TestResults
	}
	return toSerialize, nil
}

type NullableDynamicTestLinks struct {
	value *DynamicTestLinks
	isSet bool
}

func (v NullableDynamicTestLinks) Get() *DynamicTestLinks {
	return v.value
}

func (v *NullableDynamicTestLinks) Set(val *DynamicTestLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicTestLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicTestLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicTestLinks(val *DynamicTestLinks) *NullableDynamicTestLinks {
	return &NullableDynamicTestLinks{value: val, isSet: true}
}

func (v NullableDynamicTestLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicTestLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


