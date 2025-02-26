/*
Dashboards API

Manage ThousandEyes Dashboards.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboards

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the ApiTestTableData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiTestTableData{}

// ApiTestTableData Data shown in a test table widget.
type ApiTestTableData struct {
	// Identifier of the test.
	TestId *string `json:"testId,omitempty"`
	// Name of the test.
	TestName *string `json:"testName,omitempty"`
	// Configured target of the test.
	Target *string `json:"target,omitempty"`
	// Type of the test.
	TestType *string `json:"testType,omitempty"`
	// Number of active alerts of the test.
	AlertCount *int64 `json:"alertCount,omitempty"`
	// Set to `true` if test is shared, `false` otherwise.
	IsShared *bool `json:"isShared,omitempty"`
	// List of time series points for test metrics in the last 12 hours.
	Graphlets []ApiTestTableGraphletsData `json:"graphlets,omitempty"`
}

// NewApiTestTableData instantiates a new ApiTestTableData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTestTableData() *ApiTestTableData {
	this := ApiTestTableData{}
	return &this
}

// NewApiTestTableDataWithDefaults instantiates a new ApiTestTableData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTestTableDataWithDefaults() *ApiTestTableData {
	this := ApiTestTableData{}
	return &this
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *ApiTestTableData) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableData) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *ApiTestTableData) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *ApiTestTableData) SetTestId(v string) {
	o.TestId = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *ApiTestTableData) GetTestName() string {
	if o == nil || utils.IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableData) GetTestNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *ApiTestTableData) HasTestName() bool {
	if o != nil && !utils.IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *ApiTestTableData) SetTestName(v string) {
	o.TestName = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ApiTestTableData) GetTarget() string {
	if o == nil || utils.IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableData) GetTargetOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ApiTestTableData) HasTarget() bool {
	if o != nil && !utils.IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *ApiTestTableData) SetTarget(v string) {
	o.Target = &v
}

// GetTestType returns the TestType field value if set, zero value otherwise.
func (o *ApiTestTableData) GetTestType() string {
	if o == nil || utils.IsNil(o.TestType) {
		var ret string
		return ret
	}
	return *o.TestType
}

// GetTestTypeOk returns a tuple with the TestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableData) GetTestTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestType) {
		return nil, false
	}
	return o.TestType, true
}

// HasTestType returns a boolean if a field has been set.
func (o *ApiTestTableData) HasTestType() bool {
	if o != nil && !utils.IsNil(o.TestType) {
		return true
	}

	return false
}

// SetTestType gets a reference to the given string and assigns it to the TestType field.
func (o *ApiTestTableData) SetTestType(v string) {
	o.TestType = &v
}

// GetAlertCount returns the AlertCount field value if set, zero value otherwise.
func (o *ApiTestTableData) GetAlertCount() int64 {
	if o == nil || utils.IsNil(o.AlertCount) {
		var ret int64
		return ret
	}
	return *o.AlertCount
}

// GetAlertCountOk returns a tuple with the AlertCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableData) GetAlertCountOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.AlertCount) {
		return nil, false
	}
	return o.AlertCount, true
}

// HasAlertCount returns a boolean if a field has been set.
func (o *ApiTestTableData) HasAlertCount() bool {
	if o != nil && !utils.IsNil(o.AlertCount) {
		return true
	}

	return false
}

// SetAlertCount gets a reference to the given int64 and assigns it to the AlertCount field.
func (o *ApiTestTableData) SetAlertCount(v int64) {
	o.AlertCount = &v
}

// GetIsShared returns the IsShared field value if set, zero value otherwise.
func (o *ApiTestTableData) GetIsShared() bool {
	if o == nil || utils.IsNil(o.IsShared) {
		var ret bool
		return ret
	}
	return *o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableData) GetIsSharedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsShared) {
		return nil, false
	}
	return o.IsShared, true
}

// HasIsShared returns a boolean if a field has been set.
func (o *ApiTestTableData) HasIsShared() bool {
	if o != nil && !utils.IsNil(o.IsShared) {
		return true
	}

	return false
}

// SetIsShared gets a reference to the given bool and assigns it to the IsShared field.
func (o *ApiTestTableData) SetIsShared(v bool) {
	o.IsShared = &v
}

// GetGraphlets returns the Graphlets field value if set, zero value otherwise.
func (o *ApiTestTableData) GetGraphlets() []ApiTestTableGraphletsData {
	if o == nil || utils.IsNil(o.Graphlets) {
		var ret []ApiTestTableGraphletsData
		return ret
	}
	return o.Graphlets
}

// GetGraphletsOk returns a tuple with the Graphlets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableData) GetGraphletsOk() ([]ApiTestTableGraphletsData, bool) {
	if o == nil || utils.IsNil(o.Graphlets) {
		return nil, false
	}
	return o.Graphlets, true
}

// HasGraphlets returns a boolean if a field has been set.
func (o *ApiTestTableData) HasGraphlets() bool {
	if o != nil && !utils.IsNil(o.Graphlets) {
		return true
	}

	return false
}

// SetGraphlets gets a reference to the given []ApiTestTableGraphletsData and assigns it to the Graphlets field.
func (o *ApiTestTableData) SetGraphlets(v []ApiTestTableGraphletsData) {
	o.Graphlets = v
}

func (o ApiTestTableData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTestTableData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.TestId) {
		toSerialize["testId"] = o.TestId
	}
	if !utils.IsNil(o.TestName) {
		toSerialize["testName"] = o.TestName
	}
	if !utils.IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !utils.IsNil(o.TestType) {
		toSerialize["testType"] = o.TestType
	}
	if !utils.IsNil(o.AlertCount) {
		toSerialize["alertCount"] = o.AlertCount
	}
	if !utils.IsNil(o.IsShared) {
		toSerialize["isShared"] = o.IsShared
	}
	if !utils.IsNil(o.Graphlets) {
		toSerialize["graphlets"] = o.Graphlets
	}
	return toSerialize, nil
}

type NullableApiTestTableData struct {
	value *ApiTestTableData
	isSet bool
}

func (v NullableApiTestTableData) Get() *ApiTestTableData {
	return v.value
}

func (v *NullableApiTestTableData) Set(val *ApiTestTableData) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTestTableData) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTestTableData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTestTableData(val *ApiTestTableData) *NullableApiTestTableData {
	return &NullableApiTestTableData{value: val, isSet: true}
}

func (v NullableApiTestTableData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTestTableData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


