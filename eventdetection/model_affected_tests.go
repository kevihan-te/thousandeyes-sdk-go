/*
Event Detection API

 Event detection occurs when ThousandEyes identifies that error signals related to a component (proxy, network node, AS, server etc) have deviated from the baselines established by events. * To determine this, ThousandEyes takes the test results from all accounts groups within an organization, and analyzes that data. * Noisy test results (those that have too many errors in a short window) are removed until they stabilize, and the rest of the results are tagged with the components associated with that test result (for example, proxy, network, or server). * Next, any increase in failures from the test results and each component helps in determining the problem domain and which component may be at fault. * When this failure rate increases beyond a pre-defined threshold (set by the algorithm), an event is triggered and an email notification is sent to the user (if they've enabled email alerts).  With the Events API, you can perform the following tasks on the ThousandEyes platform: * **Retrieve Events**: Obtain a list of events and detailed information for each event. For more information about events, see [Event Detection](https://docs.thousandeyes.com/product-documentation/event-detection). 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eventdetection

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the AffectedTests type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffectedTests{}

// AffectedTests struct for AffectedTests
type AffectedTests struct {
	// The total number affected.
	Total *int32 `json:"total,omitempty"`
	// Indicates if in the affected account group.
	InAccountGroup *int32 `json:"inAccountGroup,omitempty"`
	// List of affected tests.
	Tests []EventApiAffectedTest `json:"tests,omitempty"`
}

// NewAffectedTests instantiates a new AffectedTests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffectedTests() *AffectedTests {
	this := AffectedTests{}
	return &this
}

// NewAffectedTestsWithDefaults instantiates a new AffectedTests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffectedTestsWithDefaults() *AffectedTests {
	this := AffectedTests{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *AffectedTests) GetTotal() int32 {
	if o == nil || utils.IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedTests) GetTotalOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *AffectedTests) HasTotal() bool {
	if o != nil && !utils.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *AffectedTests) SetTotal(v int32) {
	o.Total = &v
}

// GetInAccountGroup returns the InAccountGroup field value if set, zero value otherwise.
func (o *AffectedTests) GetInAccountGroup() int32 {
	if o == nil || utils.IsNil(o.InAccountGroup) {
		var ret int32
		return ret
	}
	return *o.InAccountGroup
}

// GetInAccountGroupOk returns a tuple with the InAccountGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedTests) GetInAccountGroupOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.InAccountGroup) {
		return nil, false
	}
	return o.InAccountGroup, true
}

// HasInAccountGroup returns a boolean if a field has been set.
func (o *AffectedTests) HasInAccountGroup() bool {
	if o != nil && !utils.IsNil(o.InAccountGroup) {
		return true
	}

	return false
}

// SetInAccountGroup gets a reference to the given int32 and assigns it to the InAccountGroup field.
func (o *AffectedTests) SetInAccountGroup(v int32) {
	o.InAccountGroup = &v
}

// GetTests returns the Tests field value if set, zero value otherwise.
func (o *AffectedTests) GetTests() []EventApiAffectedTest {
	if o == nil || utils.IsNil(o.Tests) {
		var ret []EventApiAffectedTest
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedTests) GetTestsOk() ([]EventApiAffectedTest, bool) {
	if o == nil || utils.IsNil(o.Tests) {
		return nil, false
	}
	return o.Tests, true
}

// HasTests returns a boolean if a field has been set.
func (o *AffectedTests) HasTests() bool {
	if o != nil && !utils.IsNil(o.Tests) {
		return true
	}

	return false
}

// SetTests gets a reference to the given []EventApiAffectedTest and assigns it to the Tests field.
func (o *AffectedTests) SetTests(v []EventApiAffectedTest) {
	o.Tests = v
}

func (o AffectedTests) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AffectedTests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !utils.IsNil(o.InAccountGroup) {
		toSerialize["inAccountGroup"] = o.InAccountGroup
	}
	if !utils.IsNil(o.Tests) {
		toSerialize["tests"] = o.Tests
	}
	return toSerialize, nil
}

type NullableAffectedTests struct {
	value *AffectedTests
	isSet bool
}

func (v NullableAffectedTests) Get() *AffectedTests {
	return v.value
}

func (v *NullableAffectedTests) Set(val *AffectedTests) {
	v.value = val
	v.isSet = true
}

func (v NullableAffectedTests) IsSet() bool {
	return v.isSet
}

func (v *NullableAffectedTests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffectedTests(val *AffectedTests) *NullableAffectedTests {
	return &NullableAffectedTests{value: val, isSet: true}
}

func (v NullableAffectedTests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffectedTests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


