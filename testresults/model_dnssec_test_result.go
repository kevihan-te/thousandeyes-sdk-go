/*
Test Results API

Get test result metrics for Cloud and Enterprise Agent tests.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
)

// checks if the DnssecTestResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DnssecTestResult{}

// DnssecTestResult struct for DnssecTestResult
type DnssecTestResult struct {
	// Data point date UTC (ISO date-time format).
	Date *time.Time `json:"date,omitempty"`
	// Epoch time (seconds) indicating the start time of the round
	RoundId *int32 `json:"roundId,omitempty"`
	Links *TestResultAppLinks `json:"_links,omitempty"`
	// Epoch time (seconds) indicating the start time of the round
	StartTime *int32 `json:"startTime,omitempty"`
	// Epoch time (seconds) indicating the end time of the round
	EndTime *int32 `json:"endTime,omitempty"`
	Agent *TestResultAgent `json:"agent,omitempty"`
	// Indicates if keychain is valid (if false see errorDetails field)
	IsValid *bool `json:"isValid,omitempty"`
	// Error details, if an error were encountered
	ErrorDetails *string `json:"errorDetails,omitempty"`
}

// NewDnssecTestResult instantiates a new DnssecTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnssecTestResult() *DnssecTestResult {
	this := DnssecTestResult{}
	return &this
}

// NewDnssecTestResultWithDefaults instantiates a new DnssecTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnssecTestResultWithDefaults() *DnssecTestResult {
	this := DnssecTestResult{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *DnssecTestResult) GetDate() time.Time {
	if o == nil || utils.IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnssecTestResult) GetDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *DnssecTestResult) HasDate() bool {
	if o != nil && !utils.IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *DnssecTestResult) SetDate(v time.Time) {
	o.Date = &v
}

// GetRoundId returns the RoundId field value if set, zero value otherwise.
func (o *DnssecTestResult) GetRoundId() int32 {
	if o == nil || utils.IsNil(o.RoundId) {
		var ret int32
		return ret
	}
	return *o.RoundId
}

// GetRoundIdOk returns a tuple with the RoundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnssecTestResult) GetRoundIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RoundId) {
		return nil, false
	}
	return o.RoundId, true
}

// HasRoundId returns a boolean if a field has been set.
func (o *DnssecTestResult) HasRoundId() bool {
	if o != nil && !utils.IsNil(o.RoundId) {
		return true
	}

	return false
}

// SetRoundId gets a reference to the given int32 and assigns it to the RoundId field.
func (o *DnssecTestResult) SetRoundId(v int32) {
	o.RoundId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DnssecTestResult) GetLinks() TestResultAppLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestResultAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnssecTestResult) GetLinksOk() (*TestResultAppLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DnssecTestResult) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestResultAppLinks and assigns it to the Links field.
func (o *DnssecTestResult) SetLinks(v TestResultAppLinks) {
	o.Links = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *DnssecTestResult) GetStartTime() int32 {
	if o == nil || utils.IsNil(o.StartTime) {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnssecTestResult) GetStartTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *DnssecTestResult) HasStartTime() bool {
	if o != nil && !utils.IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *DnssecTestResult) SetStartTime(v int32) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *DnssecTestResult) GetEndTime() int32 {
	if o == nil || utils.IsNil(o.EndTime) {
		var ret int32
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnssecTestResult) GetEndTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *DnssecTestResult) HasEndTime() bool {
	if o != nil && !utils.IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int32 and assigns it to the EndTime field.
func (o *DnssecTestResult) SetEndTime(v int32) {
	o.EndTime = &v
}

// GetAgent returns the Agent field value if set, zero value otherwise.
func (o *DnssecTestResult) GetAgent() TestResultAgent {
	if o == nil || utils.IsNil(o.Agent) {
		var ret TestResultAgent
		return ret
	}
	return *o.Agent
}

// GetAgentOk returns a tuple with the Agent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnssecTestResult) GetAgentOk() (*TestResultAgent, bool) {
	if o == nil || utils.IsNil(o.Agent) {
		return nil, false
	}
	return o.Agent, true
}

// HasAgent returns a boolean if a field has been set.
func (o *DnssecTestResult) HasAgent() bool {
	if o != nil && !utils.IsNil(o.Agent) {
		return true
	}

	return false
}

// SetAgent gets a reference to the given TestResultAgent and assigns it to the Agent field.
func (o *DnssecTestResult) SetAgent(v TestResultAgent) {
	o.Agent = &v
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *DnssecTestResult) GetIsValid() bool {
	if o == nil || utils.IsNil(o.IsValid) {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnssecTestResult) GetIsValidOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsValid) {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *DnssecTestResult) HasIsValid() bool {
	if o != nil && !utils.IsNil(o.IsValid) {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *DnssecTestResult) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *DnssecTestResult) GetErrorDetails() string {
	if o == nil || utils.IsNil(o.ErrorDetails) {
		var ret string
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnssecTestResult) GetErrorDetailsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *DnssecTestResult) HasErrorDetails() bool {
	if o != nil && !utils.IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given string and assigns it to the ErrorDetails field.
func (o *DnssecTestResult) SetErrorDetails(v string) {
	o.ErrorDetails = &v
}

func (o DnssecTestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnssecTestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !utils.IsNil(o.RoundId) {
		toSerialize["roundId"] = o.RoundId
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !utils.IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !utils.IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !utils.IsNil(o.Agent) {
		toSerialize["agent"] = o.Agent
	}
	if !utils.IsNil(o.IsValid) {
		toSerialize["isValid"] = o.IsValid
	}
	if !utils.IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	return toSerialize, nil
}

type NullableDnssecTestResult struct {
	value *DnssecTestResult
	isSet bool
}

func (v NullableDnssecTestResult) Get() *DnssecTestResult {
	return v.value
}

func (v *NullableDnssecTestResult) Set(val *DnssecTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDnssecTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDnssecTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnssecTestResult(val *DnssecTestResult) *NullableDnssecTestResult {
	return &NullableDnssecTestResult{value: val, isSet: true}
}

func (v NullableDnssecTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnssecTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


