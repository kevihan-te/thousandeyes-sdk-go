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

// checks if the WebTransactionTestResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebTransactionTestResult{}

// WebTransactionTestResult struct for WebTransactionTestResult
type WebTransactionTestResult struct {
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
	// Number of components which did not successfully load
	ComponentErrors *int32 `json:"componentErrors,omitempty"`
	// Elapsed execution time of the web transaction script in milliseconds
	TransactionTime *int32 `json:"transactionTime,omitempty"`
	// Type of error encountered; corresponds to phase of connection
	ErrorType *string `json:"errorType,omitempty"`
	// Error details, if an error were encountered
	ErrorDetails *string `json:"errorDetails,omitempty"`
}

// NewWebTransactionTestResult instantiates a new WebTransactionTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebTransactionTestResult() *WebTransactionTestResult {
	this := WebTransactionTestResult{}
	return &this
}

// NewWebTransactionTestResultWithDefaults instantiates a new WebTransactionTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebTransactionTestResultWithDefaults() *WebTransactionTestResult {
	this := WebTransactionTestResult{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *WebTransactionTestResult) GetDate() time.Time {
	if o == nil || utils.IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebTransactionTestResult) GetDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *WebTransactionTestResult) HasDate() bool {
	if o != nil && !utils.IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *WebTransactionTestResult) SetDate(v time.Time) {
	o.Date = &v
}

// GetRoundId returns the RoundId field value if set, zero value otherwise.
func (o *WebTransactionTestResult) GetRoundId() int32 {
	if o == nil || utils.IsNil(o.RoundId) {
		var ret int32
		return ret
	}
	return *o.RoundId
}

// GetRoundIdOk returns a tuple with the RoundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebTransactionTestResult) GetRoundIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RoundId) {
		return nil, false
	}
	return o.RoundId, true
}

// HasRoundId returns a boolean if a field has been set.
func (o *WebTransactionTestResult) HasRoundId() bool {
	if o != nil && !utils.IsNil(o.RoundId) {
		return true
	}

	return false
}

// SetRoundId gets a reference to the given int32 and assigns it to the RoundId field.
func (o *WebTransactionTestResult) SetRoundId(v int32) {
	o.RoundId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WebTransactionTestResult) GetLinks() TestResultAppLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestResultAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebTransactionTestResult) GetLinksOk() (*TestResultAppLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WebTransactionTestResult) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestResultAppLinks and assigns it to the Links field.
func (o *WebTransactionTestResult) SetLinks(v TestResultAppLinks) {
	o.Links = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *WebTransactionTestResult) GetStartTime() int32 {
	if o == nil || utils.IsNil(o.StartTime) {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebTransactionTestResult) GetStartTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *WebTransactionTestResult) HasStartTime() bool {
	if o != nil && !utils.IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *WebTransactionTestResult) SetStartTime(v int32) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *WebTransactionTestResult) GetEndTime() int32 {
	if o == nil || utils.IsNil(o.EndTime) {
		var ret int32
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebTransactionTestResult) GetEndTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *WebTransactionTestResult) HasEndTime() bool {
	if o != nil && !utils.IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int32 and assigns it to the EndTime field.
func (o *WebTransactionTestResult) SetEndTime(v int32) {
	o.EndTime = &v
}

// GetAgent returns the Agent field value if set, zero value otherwise.
func (o *WebTransactionTestResult) GetAgent() TestResultAgent {
	if o == nil || utils.IsNil(o.Agent) {
		var ret TestResultAgent
		return ret
	}
	return *o.Agent
}

// GetAgentOk returns a tuple with the Agent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebTransactionTestResult) GetAgentOk() (*TestResultAgent, bool) {
	if o == nil || utils.IsNil(o.Agent) {
		return nil, false
	}
	return o.Agent, true
}

// HasAgent returns a boolean if a field has been set.
func (o *WebTransactionTestResult) HasAgent() bool {
	if o != nil && !utils.IsNil(o.Agent) {
		return true
	}

	return false
}

// SetAgent gets a reference to the given TestResultAgent and assigns it to the Agent field.
func (o *WebTransactionTestResult) SetAgent(v TestResultAgent) {
	o.Agent = &v
}

// GetComponentErrors returns the ComponentErrors field value if set, zero value otherwise.
func (o *WebTransactionTestResult) GetComponentErrors() int32 {
	if o == nil || utils.IsNil(o.ComponentErrors) {
		var ret int32
		return ret
	}
	return *o.ComponentErrors
}

// GetComponentErrorsOk returns a tuple with the ComponentErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebTransactionTestResult) GetComponentErrorsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ComponentErrors) {
		return nil, false
	}
	return o.ComponentErrors, true
}

// HasComponentErrors returns a boolean if a field has been set.
func (o *WebTransactionTestResult) HasComponentErrors() bool {
	if o != nil && !utils.IsNil(o.ComponentErrors) {
		return true
	}

	return false
}

// SetComponentErrors gets a reference to the given int32 and assigns it to the ComponentErrors field.
func (o *WebTransactionTestResult) SetComponentErrors(v int32) {
	o.ComponentErrors = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *WebTransactionTestResult) GetTransactionTime() int32 {
	if o == nil || utils.IsNil(o.TransactionTime) {
		var ret int32
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebTransactionTestResult) GetTransactionTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *WebTransactionTestResult) HasTransactionTime() bool {
	if o != nil && !utils.IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int32 and assigns it to the TransactionTime field.
func (o *WebTransactionTestResult) SetTransactionTime(v int32) {
	o.TransactionTime = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *WebTransactionTestResult) GetErrorType() string {
	if o == nil || utils.IsNil(o.ErrorType) {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebTransactionTestResult) GetErrorTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorType) {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *WebTransactionTestResult) HasErrorType() bool {
	if o != nil && !utils.IsNil(o.ErrorType) {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *WebTransactionTestResult) SetErrorType(v string) {
	o.ErrorType = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *WebTransactionTestResult) GetErrorDetails() string {
	if o == nil || utils.IsNil(o.ErrorDetails) {
		var ret string
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebTransactionTestResult) GetErrorDetailsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *WebTransactionTestResult) HasErrorDetails() bool {
	if o != nil && !utils.IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given string and assigns it to the ErrorDetails field.
func (o *WebTransactionTestResult) SetErrorDetails(v string) {
	o.ErrorDetails = &v
}

func (o WebTransactionTestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebTransactionTestResult) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.ComponentErrors) {
		toSerialize["componentErrors"] = o.ComponentErrors
	}
	if !utils.IsNil(o.TransactionTime) {
		toSerialize["transactionTime"] = o.TransactionTime
	}
	if !utils.IsNil(o.ErrorType) {
		toSerialize["errorType"] = o.ErrorType
	}
	if !utils.IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	return toSerialize, nil
}

type NullableWebTransactionTestResult struct {
	value *WebTransactionTestResult
	isSet bool
}

func (v NullableWebTransactionTestResult) Get() *WebTransactionTestResult {
	return v.value
}

func (v *NullableWebTransactionTestResult) Set(val *WebTransactionTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableWebTransactionTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableWebTransactionTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebTransactionTestResult(val *WebTransactionTestResult) *NullableWebTransactionTestResult {
	return &NullableWebTransactionTestResult{value: val, isSet: true}
}

func (v NullableWebTransactionTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebTransactionTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


