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
	"time"
)

// checks if the RealUserEndpointTestPageResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RealUserEndpointTestPageResult{}

// RealUserEndpointTestPageResult struct for RealUserEndpointTestPageResult
type RealUserEndpointTestPageResult struct {
	// Web page ID. The page ID is unique only within an endpoint real user test.
	PageId *string `json:"pageId,omitempty"`
	// Web page title.
	PageTitle *string `json:"pageTitle,omitempty"`
	// Web page url
	PageUrl *string `json:"pageUrl,omitempty"`
	// UTC date when page load started (ISO date-time format).
	LoadDate *time.Time `json:"loadDate,omitempty"`
	// HTTP response code.
	ResponseCode *int32 `json:"responseCode,omitempty"`
	PageTimings *RealUserEndpointTestPageTimings `json:"pageTimings,omitempty"`
	// Unique ID of endpoint agent, from `/endpoint/agents` endpoint.
	AgentId *string `json:"agentId,omitempty"`
	// Endpoint real user test ID. Each endpoint real user test occurrence has a unique ID.
	Id *string `json:"id,omitempty"`
	// Epoch time (seconds) indicating the start time of the round.
	RoundId *int32 `json:"roundId,omitempty"`
	// HTTP server response in milliseconds.
	ResponseTime *int32 `json:"responseTime,omitempty"`
	SystemMetrics *SystemMetrics `json:"systemMetrics,omitempty"`
}

// NewRealUserEndpointTestPageResult instantiates a new RealUserEndpointTestPageResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealUserEndpointTestPageResult() *RealUserEndpointTestPageResult {
	this := RealUserEndpointTestPageResult{}
	return &this
}

// NewRealUserEndpointTestPageResultWithDefaults instantiates a new RealUserEndpointTestPageResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealUserEndpointTestPageResultWithDefaults() *RealUserEndpointTestPageResult {
	this := RealUserEndpointTestPageResult{}
	return &this
}

// GetPageId returns the PageId field value if set, zero value otherwise.
func (o *RealUserEndpointTestPageResult) GetPageId() string {
	if o == nil || utils.IsNil(o.PageId) {
		var ret string
		return ret
	}
	return *o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestPageResult) GetPageIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PageId) {
		return nil, false
	}
	return o.PageId, true
}

// HasPageId returns a boolean if a field has been set.
func (o *RealUserEndpointTestPageResult) HasPageId() bool {
	if o != nil && !utils.IsNil(o.PageId) {
		return true
	}

	return false
}

// SetPageId gets a reference to the given string and assigns it to the PageId field.
func (o *RealUserEndpointTestPageResult) SetPageId(v string) {
	o.PageId = &v
}

// GetPageTitle returns the PageTitle field value if set, zero value otherwise.
func (o *RealUserEndpointTestPageResult) GetPageTitle() string {
	if o == nil || utils.IsNil(o.PageTitle) {
		var ret string
		return ret
	}
	return *o.PageTitle
}

// GetPageTitleOk returns a tuple with the PageTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestPageResult) GetPageTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PageTitle) {
		return nil, false
	}
	return o.PageTitle, true
}

// HasPageTitle returns a boolean if a field has been set.
func (o *RealUserEndpointTestPageResult) HasPageTitle() bool {
	if o != nil && !utils.IsNil(o.PageTitle) {
		return true
	}

	return false
}

// SetPageTitle gets a reference to the given string and assigns it to the PageTitle field.
func (o *RealUserEndpointTestPageResult) SetPageTitle(v string) {
	o.PageTitle = &v
}

// GetPageUrl returns the PageUrl field value if set, zero value otherwise.
func (o *RealUserEndpointTestPageResult) GetPageUrl() string {
	if o == nil || utils.IsNil(o.PageUrl) {
		var ret string
		return ret
	}
	return *o.PageUrl
}

// GetPageUrlOk returns a tuple with the PageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestPageResult) GetPageUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PageUrl) {
		return nil, false
	}
	return o.PageUrl, true
}

// HasPageUrl returns a boolean if a field has been set.
func (o *RealUserEndpointTestPageResult) HasPageUrl() bool {
	if o != nil && !utils.IsNil(o.PageUrl) {
		return true
	}

	return false
}

// SetPageUrl gets a reference to the given string and assigns it to the PageUrl field.
func (o *RealUserEndpointTestPageResult) SetPageUrl(v string) {
	o.PageUrl = &v
}

// GetLoadDate returns the LoadDate field value if set, zero value otherwise.
func (o *RealUserEndpointTestPageResult) GetLoadDate() time.Time {
	if o == nil || utils.IsNil(o.LoadDate) {
		var ret time.Time
		return ret
	}
	return *o.LoadDate
}

// GetLoadDateOk returns a tuple with the LoadDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestPageResult) GetLoadDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.LoadDate) {
		return nil, false
	}
	return o.LoadDate, true
}

// HasLoadDate returns a boolean if a field has been set.
func (o *RealUserEndpointTestPageResult) HasLoadDate() bool {
	if o != nil && !utils.IsNil(o.LoadDate) {
		return true
	}

	return false
}

// SetLoadDate gets a reference to the given time.Time and assigns it to the LoadDate field.
func (o *RealUserEndpointTestPageResult) SetLoadDate(v time.Time) {
	o.LoadDate = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *RealUserEndpointTestPageResult) GetResponseCode() int32 {
	if o == nil || utils.IsNil(o.ResponseCode) {
		var ret int32
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestPageResult) GetResponseCodeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ResponseCode) {
		return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *RealUserEndpointTestPageResult) HasResponseCode() bool {
	if o != nil && !utils.IsNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given int32 and assigns it to the ResponseCode field.
func (o *RealUserEndpointTestPageResult) SetResponseCode(v int32) {
	o.ResponseCode = &v
}

// GetPageTimings returns the PageTimings field value if set, zero value otherwise.
func (o *RealUserEndpointTestPageResult) GetPageTimings() RealUserEndpointTestPageTimings {
	if o == nil || utils.IsNil(o.PageTimings) {
		var ret RealUserEndpointTestPageTimings
		return ret
	}
	return *o.PageTimings
}

// GetPageTimingsOk returns a tuple with the PageTimings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestPageResult) GetPageTimingsOk() (*RealUserEndpointTestPageTimings, bool) {
	if o == nil || utils.IsNil(o.PageTimings) {
		return nil, false
	}
	return o.PageTimings, true
}

// HasPageTimings returns a boolean if a field has been set.
func (o *RealUserEndpointTestPageResult) HasPageTimings() bool {
	if o != nil && !utils.IsNil(o.PageTimings) {
		return true
	}

	return false
}

// SetPageTimings gets a reference to the given RealUserEndpointTestPageTimings and assigns it to the PageTimings field.
func (o *RealUserEndpointTestPageResult) SetPageTimings(v RealUserEndpointTestPageTimings) {
	o.PageTimings = &v
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *RealUserEndpointTestPageResult) GetAgentId() string {
	if o == nil || utils.IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestPageResult) GetAgentIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *RealUserEndpointTestPageResult) HasAgentId() bool {
	if o != nil && !utils.IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *RealUserEndpointTestPageResult) SetAgentId(v string) {
	o.AgentId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RealUserEndpointTestPageResult) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestPageResult) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RealUserEndpointTestPageResult) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RealUserEndpointTestPageResult) SetId(v string) {
	o.Id = &v
}

// GetRoundId returns the RoundId field value if set, zero value otherwise.
func (o *RealUserEndpointTestPageResult) GetRoundId() int32 {
	if o == nil || utils.IsNil(o.RoundId) {
		var ret int32
		return ret
	}
	return *o.RoundId
}

// GetRoundIdOk returns a tuple with the RoundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestPageResult) GetRoundIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RoundId) {
		return nil, false
	}
	return o.RoundId, true
}

// HasRoundId returns a boolean if a field has been set.
func (o *RealUserEndpointTestPageResult) HasRoundId() bool {
	if o != nil && !utils.IsNil(o.RoundId) {
		return true
	}

	return false
}

// SetRoundId gets a reference to the given int32 and assigns it to the RoundId field.
func (o *RealUserEndpointTestPageResult) SetRoundId(v int32) {
	o.RoundId = &v
}

// GetResponseTime returns the ResponseTime field value if set, zero value otherwise.
func (o *RealUserEndpointTestPageResult) GetResponseTime() int32 {
	if o == nil || utils.IsNil(o.ResponseTime) {
		var ret int32
		return ret
	}
	return *o.ResponseTime
}

// GetResponseTimeOk returns a tuple with the ResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestPageResult) GetResponseTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ResponseTime) {
		return nil, false
	}
	return o.ResponseTime, true
}

// HasResponseTime returns a boolean if a field has been set.
func (o *RealUserEndpointTestPageResult) HasResponseTime() bool {
	if o != nil && !utils.IsNil(o.ResponseTime) {
		return true
	}

	return false
}

// SetResponseTime gets a reference to the given int32 and assigns it to the ResponseTime field.
func (o *RealUserEndpointTestPageResult) SetResponseTime(v int32) {
	o.ResponseTime = &v
}

// GetSystemMetrics returns the SystemMetrics field value if set, zero value otherwise.
func (o *RealUserEndpointTestPageResult) GetSystemMetrics() SystemMetrics {
	if o == nil || utils.IsNil(o.SystemMetrics) {
		var ret SystemMetrics
		return ret
	}
	return *o.SystemMetrics
}

// GetSystemMetricsOk returns a tuple with the SystemMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestPageResult) GetSystemMetricsOk() (*SystemMetrics, bool) {
	if o == nil || utils.IsNil(o.SystemMetrics) {
		return nil, false
	}
	return o.SystemMetrics, true
}

// HasSystemMetrics returns a boolean if a field has been set.
func (o *RealUserEndpointTestPageResult) HasSystemMetrics() bool {
	if o != nil && !utils.IsNil(o.SystemMetrics) {
		return true
	}

	return false
}

// SetSystemMetrics gets a reference to the given SystemMetrics and assigns it to the SystemMetrics field.
func (o *RealUserEndpointTestPageResult) SetSystemMetrics(v SystemMetrics) {
	o.SystemMetrics = &v
}

func (o RealUserEndpointTestPageResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealUserEndpointTestPageResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PageId) {
		toSerialize["pageId"] = o.PageId
	}
	if !utils.IsNil(o.PageTitle) {
		toSerialize["pageTitle"] = o.PageTitle
	}
	if !utils.IsNil(o.PageUrl) {
		toSerialize["pageUrl"] = o.PageUrl
	}
	if !utils.IsNil(o.LoadDate) {
		toSerialize["loadDate"] = o.LoadDate
	}
	if !utils.IsNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !utils.IsNil(o.PageTimings) {
		toSerialize["pageTimings"] = o.PageTimings
	}
	if !utils.IsNil(o.AgentId) {
		toSerialize["agentId"] = o.AgentId
	}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.RoundId) {
		toSerialize["roundId"] = o.RoundId
	}
	if !utils.IsNil(o.ResponseTime) {
		toSerialize["responseTime"] = o.ResponseTime
	}
	if !utils.IsNil(o.SystemMetrics) {
		toSerialize["systemMetrics"] = o.SystemMetrics
	}
	return toSerialize, nil
}

type NullableRealUserEndpointTestPageResult struct {
	value *RealUserEndpointTestPageResult
	isSet bool
}

func (v NullableRealUserEndpointTestPageResult) Get() *RealUserEndpointTestPageResult {
	return v.value
}

func (v *NullableRealUserEndpointTestPageResult) Set(val *RealUserEndpointTestPageResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRealUserEndpointTestPageResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRealUserEndpointTestPageResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealUserEndpointTestPageResult(val *RealUserEndpointTestPageResult) *NullableRealUserEndpointTestPageResult {
	return &NullableRealUserEndpointTestPageResult{value: val, isSet: true}
}

func (v NullableRealUserEndpointTestPageResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealUserEndpointTestPageResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


