/*
Test Results API

Get test result metrics for Cloud and Enterprise Agent tests.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
)

// checks if the PathVisDetailTestResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PathVisDetailTestResult{}

// PathVisDetailTestResult struct for PathVisDetailTestResult
type PathVisDetailTestResult struct {
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
	// Target server, including port (if method used is TCP)
	Server *string `json:"server,omitempty"`
	// IP of target server
	ServerIp *string `json:"serverIp,omitempty"`
	// IP address of source agent
	SourceIp *string `json:"sourceIp,omitempty"`
	// IP prefix of source agent
	SourcePrefix *string `json:"sourcePrefix,omitempty"`
	// Specifies whether the traces are targeting a proxy. If not set, it is considered as false.
	TargetIsProxy *bool `json:"targetIsProxy,omitempty"`
	Direction *PathVisDirection `json:"direction,omitempty"`
	// Shows 3 iterations of path trace, with each iteration specified by a pathId
	PathTraces []PathVisRoute `json:"pathTraces,omitempty"`
}

// NewPathVisDetailTestResult instantiates a new PathVisDetailTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathVisDetailTestResult() *PathVisDetailTestResult {
	this := PathVisDetailTestResult{}
	return &this
}

// NewPathVisDetailTestResultWithDefaults instantiates a new PathVisDetailTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathVisDetailTestResultWithDefaults() *PathVisDetailTestResult {
	this := PathVisDetailTestResult{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *PathVisDetailTestResult) GetDate() time.Time {
	if o == nil || utils.IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisDetailTestResult) GetDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *PathVisDetailTestResult) HasDate() bool {
	if o != nil && !utils.IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *PathVisDetailTestResult) SetDate(v time.Time) {
	o.Date = &v
}

// GetRoundId returns the RoundId field value if set, zero value otherwise.
func (o *PathVisDetailTestResult) GetRoundId() int32 {
	if o == nil || utils.IsNil(o.RoundId) {
		var ret int32
		return ret
	}
	return *o.RoundId
}

// GetRoundIdOk returns a tuple with the RoundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisDetailTestResult) GetRoundIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RoundId) {
		return nil, false
	}
	return o.RoundId, true
}

// HasRoundId returns a boolean if a field has been set.
func (o *PathVisDetailTestResult) HasRoundId() bool {
	if o != nil && !utils.IsNil(o.RoundId) {
		return true
	}

	return false
}

// SetRoundId gets a reference to the given int32 and assigns it to the RoundId field.
func (o *PathVisDetailTestResult) SetRoundId(v int32) {
	o.RoundId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PathVisDetailTestResult) GetLinks() TestResultAppLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestResultAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisDetailTestResult) GetLinksOk() (*TestResultAppLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PathVisDetailTestResult) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestResultAppLinks and assigns it to the Links field.
func (o *PathVisDetailTestResult) SetLinks(v TestResultAppLinks) {
	o.Links = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *PathVisDetailTestResult) GetStartTime() int32 {
	if o == nil || utils.IsNil(o.StartTime) {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisDetailTestResult) GetStartTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *PathVisDetailTestResult) HasStartTime() bool {
	if o != nil && !utils.IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *PathVisDetailTestResult) SetStartTime(v int32) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *PathVisDetailTestResult) GetEndTime() int32 {
	if o == nil || utils.IsNil(o.EndTime) {
		var ret int32
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisDetailTestResult) GetEndTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *PathVisDetailTestResult) HasEndTime() bool {
	if o != nil && !utils.IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int32 and assigns it to the EndTime field.
func (o *PathVisDetailTestResult) SetEndTime(v int32) {
	o.EndTime = &v
}

// GetAgent returns the Agent field value if set, zero value otherwise.
func (o *PathVisDetailTestResult) GetAgent() TestResultAgent {
	if o == nil || utils.IsNil(o.Agent) {
		var ret TestResultAgent
		return ret
	}
	return *o.Agent
}

// GetAgentOk returns a tuple with the Agent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisDetailTestResult) GetAgentOk() (*TestResultAgent, bool) {
	if o == nil || utils.IsNil(o.Agent) {
		return nil, false
	}
	return o.Agent, true
}

// HasAgent returns a boolean if a field has been set.
func (o *PathVisDetailTestResult) HasAgent() bool {
	if o != nil && !utils.IsNil(o.Agent) {
		return true
	}

	return false
}

// SetAgent gets a reference to the given TestResultAgent and assigns it to the Agent field.
func (o *PathVisDetailTestResult) SetAgent(v TestResultAgent) {
	o.Agent = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *PathVisDetailTestResult) GetServer() string {
	if o == nil || utils.IsNil(o.Server) {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisDetailTestResult) GetServerOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *PathVisDetailTestResult) HasServer() bool {
	if o != nil && !utils.IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *PathVisDetailTestResult) SetServer(v string) {
	o.Server = &v
}

// GetServerIp returns the ServerIp field value if set, zero value otherwise.
func (o *PathVisDetailTestResult) GetServerIp() string {
	if o == nil || utils.IsNil(o.ServerIp) {
		var ret string
		return ret
	}
	return *o.ServerIp
}

// GetServerIpOk returns a tuple with the ServerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisDetailTestResult) GetServerIpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ServerIp) {
		return nil, false
	}
	return o.ServerIp, true
}

// HasServerIp returns a boolean if a field has been set.
func (o *PathVisDetailTestResult) HasServerIp() bool {
	if o != nil && !utils.IsNil(o.ServerIp) {
		return true
	}

	return false
}

// SetServerIp gets a reference to the given string and assigns it to the ServerIp field.
func (o *PathVisDetailTestResult) SetServerIp(v string) {
	o.ServerIp = &v
}

// GetSourceIp returns the SourceIp field value if set, zero value otherwise.
func (o *PathVisDetailTestResult) GetSourceIp() string {
	if o == nil || utils.IsNil(o.SourceIp) {
		var ret string
		return ret
	}
	return *o.SourceIp
}

// GetSourceIpOk returns a tuple with the SourceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisDetailTestResult) GetSourceIpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SourceIp) {
		return nil, false
	}
	return o.SourceIp, true
}

// HasSourceIp returns a boolean if a field has been set.
func (o *PathVisDetailTestResult) HasSourceIp() bool {
	if o != nil && !utils.IsNil(o.SourceIp) {
		return true
	}

	return false
}

// SetSourceIp gets a reference to the given string and assigns it to the SourceIp field.
func (o *PathVisDetailTestResult) SetSourceIp(v string) {
	o.SourceIp = &v
}

// GetSourcePrefix returns the SourcePrefix field value if set, zero value otherwise.
func (o *PathVisDetailTestResult) GetSourcePrefix() string {
	if o == nil || utils.IsNil(o.SourcePrefix) {
		var ret string
		return ret
	}
	return *o.SourcePrefix
}

// GetSourcePrefixOk returns a tuple with the SourcePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisDetailTestResult) GetSourcePrefixOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SourcePrefix) {
		return nil, false
	}
	return o.SourcePrefix, true
}

// HasSourcePrefix returns a boolean if a field has been set.
func (o *PathVisDetailTestResult) HasSourcePrefix() bool {
	if o != nil && !utils.IsNil(o.SourcePrefix) {
		return true
	}

	return false
}

// SetSourcePrefix gets a reference to the given string and assigns it to the SourcePrefix field.
func (o *PathVisDetailTestResult) SetSourcePrefix(v string) {
	o.SourcePrefix = &v
}

// GetTargetIsProxy returns the TargetIsProxy field value if set, zero value otherwise.
func (o *PathVisDetailTestResult) GetTargetIsProxy() bool {
	if o == nil || utils.IsNil(o.TargetIsProxy) {
		var ret bool
		return ret
	}
	return *o.TargetIsProxy
}

// GetTargetIsProxyOk returns a tuple with the TargetIsProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisDetailTestResult) GetTargetIsProxyOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.TargetIsProxy) {
		return nil, false
	}
	return o.TargetIsProxy, true
}

// HasTargetIsProxy returns a boolean if a field has been set.
func (o *PathVisDetailTestResult) HasTargetIsProxy() bool {
	if o != nil && !utils.IsNil(o.TargetIsProxy) {
		return true
	}

	return false
}

// SetTargetIsProxy gets a reference to the given bool and assigns it to the TargetIsProxy field.
func (o *PathVisDetailTestResult) SetTargetIsProxy(v bool) {
	o.TargetIsProxy = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *PathVisDetailTestResult) GetDirection() PathVisDirection {
	if o == nil || utils.IsNil(o.Direction) {
		var ret PathVisDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisDetailTestResult) GetDirectionOk() (*PathVisDirection, bool) {
	if o == nil || utils.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *PathVisDetailTestResult) HasDirection() bool {
	if o != nil && !utils.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given PathVisDirection and assigns it to the Direction field.
func (o *PathVisDetailTestResult) SetDirection(v PathVisDirection) {
	o.Direction = &v
}

// GetPathTraces returns the PathTraces field value if set, zero value otherwise.
func (o *PathVisDetailTestResult) GetPathTraces() []PathVisRoute {
	if o == nil || utils.IsNil(o.PathTraces) {
		var ret []PathVisRoute
		return ret
	}
	return o.PathTraces
}

// GetPathTracesOk returns a tuple with the PathTraces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisDetailTestResult) GetPathTracesOk() ([]PathVisRoute, bool) {
	if o == nil || utils.IsNil(o.PathTraces) {
		return nil, false
	}
	return o.PathTraces, true
}

// HasPathTraces returns a boolean if a field has been set.
func (o *PathVisDetailTestResult) HasPathTraces() bool {
	if o != nil && !utils.IsNil(o.PathTraces) {
		return true
	}

	return false
}

// SetPathTraces gets a reference to the given []PathVisRoute and assigns it to the PathTraces field.
func (o *PathVisDetailTestResult) SetPathTraces(v []PathVisRoute) {
	o.PathTraces = v
}

func (o PathVisDetailTestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PathVisDetailTestResult) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !utils.IsNil(o.ServerIp) {
		toSerialize["serverIp"] = o.ServerIp
	}
	if !utils.IsNil(o.SourceIp) {
		toSerialize["sourceIp"] = o.SourceIp
	}
	if !utils.IsNil(o.SourcePrefix) {
		toSerialize["sourcePrefix"] = o.SourcePrefix
	}
	if !utils.IsNil(o.TargetIsProxy) {
		toSerialize["targetIsProxy"] = o.TargetIsProxy
	}
	if !utils.IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !utils.IsNil(o.PathTraces) {
		toSerialize["pathTraces"] = o.PathTraces
	}
	return toSerialize, nil
}

type NullablePathVisDetailTestResult struct {
	value *PathVisDetailTestResult
	isSet bool
}

func (v NullablePathVisDetailTestResult) Get() *PathVisDetailTestResult {
	return v.value
}

func (v *NullablePathVisDetailTestResult) Set(val *PathVisDetailTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePathVisDetailTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePathVisDetailTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathVisDetailTestResult(val *PathVisDetailTestResult) *NullablePathVisDetailTestResult {
	return &NullablePathVisDetailTestResult{value: val, isSet: true}
}

func (v NullablePathVisDetailTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathVisDetailTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


