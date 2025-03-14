/*
Alerts API

You can manage the following alert functionalities on the ThousandEyes platform using the Alerts API:  * **Alerts**: Retrieve alert details. Alerts are assigned to tests through alert rules.  * **Alert Rules**: Conditions that you configure in order to highlight or be notified of events of interest in your ThousandEyes tests. When an alert rule’s conditions are met, the associated alert is triggered and the alert becomes active. It remains active until the alert is cleared. Alert rules are reusable across multiple tests..  * **Alert Suppression Windows**: Suppress alerts for tests during periods such as planned maintenance. Windows can be one-time events or recurring events to handle periodic occurrences such as monthly downtime for maintenance.  For more information about the alerts, see [Alerts](https://docs.thousandeyes.com/product-documentation/alerts). 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alerts

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the TestLinks type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TestLinks{}

// TestLinks A list of links that can be accessed to get more information
type TestLinks struct {
	Self *TestSelfLink `json:"self,omitempty"`
	// Reference to the test results.
	TestResults []Link `json:"testResults,omitempty"`
}

// NewTestLinks instantiates a new TestLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestLinks() *TestLinks {
	this := TestLinks{}
	return &this
}

// NewTestLinksWithDefaults instantiates a new TestLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestLinksWithDefaults() *TestLinks {
	this := TestLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *TestLinks) GetSelf() TestSelfLink {
	if o == nil || utils.IsNil(o.Self) {
		var ret TestSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestLinks) GetSelfOk() (*TestSelfLink, bool) {
	if o == nil || utils.IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *TestLinks) HasSelf() bool {
	if o != nil && !utils.IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given TestSelfLink and assigns it to the Self field.
func (o *TestLinks) SetSelf(v TestSelfLink) {
	o.Self = &v
}

// GetTestResults returns the TestResults field value if set, zero value otherwise.
func (o *TestLinks) GetTestResults() []Link {
	if o == nil || utils.IsNil(o.TestResults) {
		var ret []Link
		return ret
	}
	return o.TestResults
}

// GetTestResultsOk returns a tuple with the TestResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestLinks) GetTestResultsOk() ([]Link, bool) {
	if o == nil || utils.IsNil(o.TestResults) {
		return nil, false
	}
	return o.TestResults, true
}

// HasTestResults returns a boolean if a field has been set.
func (o *TestLinks) HasTestResults() bool {
	if o != nil && !utils.IsNil(o.TestResults) {
		return true
	}

	return false
}

// SetTestResults gets a reference to the given []Link and assigns it to the TestResults field.
func (o *TestLinks) SetTestResults(v []Link) {
	o.TestResults = v
}

func (o TestLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !utils.IsNil(o.TestResults) {
		toSerialize["testResults"] = o.TestResults
	}
	return toSerialize, nil
}

type NullableTestLinks struct {
	value *TestLinks
	isSet bool
}

func (v NullableTestLinks) Get() *TestLinks {
	return v.value
}

func (v *NullableTestLinks) Set(val *TestLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableTestLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableTestLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestLinks(val *TestLinks) *NullableTestLinks {
	return &NullableTestLinks{value: val, isSet: true}
}

func (v NullableTestLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


