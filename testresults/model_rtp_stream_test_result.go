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

// checks if the RtpStreamTestResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RtpStreamTestResult{}

// RtpStreamTestResult struct for RtpStreamTestResult
type RtpStreamTestResult struct {
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
	// Target agent IP address
	ServerIp *string `json:"serverIp,omitempty"`
	// DSCP value received by the server from the agent
	Dscp *string `json:"dscp,omitempty"`
	// Name of DSCP value received by the server from the agent
	DscpName *string `json:"dscpName,omitempty"`
	// Mean opinion score for agent’s stream
	Mos *float32 `json:"mos,omitempty"`
	// Name of codec used by agen
	CodecName *string `json:"codecName,omitempty"`
	// Maximum value of Mean Opinion Score based on codec selection
	CodecMaxMos *float32 `json:"codecMaxMos,omitempty"`
	// Percentage value of packets sent from agent not received by server
	Loss *float32 `json:"loss,omitempty"`
	// Percentage of packets discarded
	Discards *float32 `json:"discards,omitempty"`
	// Time to send packets from source to server in milliseconds
	Latency *int32 `json:"latency,omitempty"`
	// Variation in packet delay in milliseconds
	Pdv *int32 `json:"pdv,omitempty"`
	// Error details, if an error was encountered
	ErrorDetail *string `json:"errorDetail,omitempty"`
}

// NewRtpStreamTestResult instantiates a new RtpStreamTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRtpStreamTestResult() *RtpStreamTestResult {
	this := RtpStreamTestResult{}
	return &this
}

// NewRtpStreamTestResultWithDefaults instantiates a new RtpStreamTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRtpStreamTestResultWithDefaults() *RtpStreamTestResult {
	this := RtpStreamTestResult{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetDate() time.Time {
	if o == nil || utils.IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasDate() bool {
	if o != nil && !utils.IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *RtpStreamTestResult) SetDate(v time.Time) {
	o.Date = &v
}

// GetRoundId returns the RoundId field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetRoundId() int32 {
	if o == nil || utils.IsNil(o.RoundId) {
		var ret int32
		return ret
	}
	return *o.RoundId
}

// GetRoundIdOk returns a tuple with the RoundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetRoundIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RoundId) {
		return nil, false
	}
	return o.RoundId, true
}

// HasRoundId returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasRoundId() bool {
	if o != nil && !utils.IsNil(o.RoundId) {
		return true
	}

	return false
}

// SetRoundId gets a reference to the given int32 and assigns it to the RoundId field.
func (o *RtpStreamTestResult) SetRoundId(v int32) {
	o.RoundId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetLinks() TestResultAppLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestResultAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetLinksOk() (*TestResultAppLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestResultAppLinks and assigns it to the Links field.
func (o *RtpStreamTestResult) SetLinks(v TestResultAppLinks) {
	o.Links = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetStartTime() int32 {
	if o == nil || utils.IsNil(o.StartTime) {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetStartTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasStartTime() bool {
	if o != nil && !utils.IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *RtpStreamTestResult) SetStartTime(v int32) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetEndTime() int32 {
	if o == nil || utils.IsNil(o.EndTime) {
		var ret int32
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetEndTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasEndTime() bool {
	if o != nil && !utils.IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int32 and assigns it to the EndTime field.
func (o *RtpStreamTestResult) SetEndTime(v int32) {
	o.EndTime = &v
}

// GetAgent returns the Agent field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetAgent() TestResultAgent {
	if o == nil || utils.IsNil(o.Agent) {
		var ret TestResultAgent
		return ret
	}
	return *o.Agent
}

// GetAgentOk returns a tuple with the Agent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetAgentOk() (*TestResultAgent, bool) {
	if o == nil || utils.IsNil(o.Agent) {
		return nil, false
	}
	return o.Agent, true
}

// HasAgent returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasAgent() bool {
	if o != nil && !utils.IsNil(o.Agent) {
		return true
	}

	return false
}

// SetAgent gets a reference to the given TestResultAgent and assigns it to the Agent field.
func (o *RtpStreamTestResult) SetAgent(v TestResultAgent) {
	o.Agent = &v
}

// GetServerIp returns the ServerIp field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetServerIp() string {
	if o == nil || utils.IsNil(o.ServerIp) {
		var ret string
		return ret
	}
	return *o.ServerIp
}

// GetServerIpOk returns a tuple with the ServerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetServerIpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ServerIp) {
		return nil, false
	}
	return o.ServerIp, true
}

// HasServerIp returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasServerIp() bool {
	if o != nil && !utils.IsNil(o.ServerIp) {
		return true
	}

	return false
}

// SetServerIp gets a reference to the given string and assigns it to the ServerIp field.
func (o *RtpStreamTestResult) SetServerIp(v string) {
	o.ServerIp = &v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetDscp() string {
	if o == nil || utils.IsNil(o.Dscp) {
		var ret string
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetDscpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Dscp) {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasDscp() bool {
	if o != nil && !utils.IsNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given string and assigns it to the Dscp field.
func (o *RtpStreamTestResult) SetDscp(v string) {
	o.Dscp = &v
}

// GetDscpName returns the DscpName field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetDscpName() string {
	if o == nil || utils.IsNil(o.DscpName) {
		var ret string
		return ret
	}
	return *o.DscpName
}

// GetDscpNameOk returns a tuple with the DscpName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetDscpNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DscpName) {
		return nil, false
	}
	return o.DscpName, true
}

// HasDscpName returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasDscpName() bool {
	if o != nil && !utils.IsNil(o.DscpName) {
		return true
	}

	return false
}

// SetDscpName gets a reference to the given string and assigns it to the DscpName field.
func (o *RtpStreamTestResult) SetDscpName(v string) {
	o.DscpName = &v
}

// GetMos returns the Mos field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetMos() float32 {
	if o == nil || utils.IsNil(o.Mos) {
		var ret float32
		return ret
	}
	return *o.Mos
}

// GetMosOk returns a tuple with the Mos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetMosOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Mos) {
		return nil, false
	}
	return o.Mos, true
}

// HasMos returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasMos() bool {
	if o != nil && !utils.IsNil(o.Mos) {
		return true
	}

	return false
}

// SetMos gets a reference to the given float32 and assigns it to the Mos field.
func (o *RtpStreamTestResult) SetMos(v float32) {
	o.Mos = &v
}

// GetCodecName returns the CodecName field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetCodecName() string {
	if o == nil || utils.IsNil(o.CodecName) {
		var ret string
		return ret
	}
	return *o.CodecName
}

// GetCodecNameOk returns a tuple with the CodecName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetCodecNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CodecName) {
		return nil, false
	}
	return o.CodecName, true
}

// HasCodecName returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasCodecName() bool {
	if o != nil && !utils.IsNil(o.CodecName) {
		return true
	}

	return false
}

// SetCodecName gets a reference to the given string and assigns it to the CodecName field.
func (o *RtpStreamTestResult) SetCodecName(v string) {
	o.CodecName = &v
}

// GetCodecMaxMos returns the CodecMaxMos field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetCodecMaxMos() float32 {
	if o == nil || utils.IsNil(o.CodecMaxMos) {
		var ret float32
		return ret
	}
	return *o.CodecMaxMos
}

// GetCodecMaxMosOk returns a tuple with the CodecMaxMos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetCodecMaxMosOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.CodecMaxMos) {
		return nil, false
	}
	return o.CodecMaxMos, true
}

// HasCodecMaxMos returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasCodecMaxMos() bool {
	if o != nil && !utils.IsNil(o.CodecMaxMos) {
		return true
	}

	return false
}

// SetCodecMaxMos gets a reference to the given float32 and assigns it to the CodecMaxMos field.
func (o *RtpStreamTestResult) SetCodecMaxMos(v float32) {
	o.CodecMaxMos = &v
}

// GetLoss returns the Loss field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetLoss() float32 {
	if o == nil || utils.IsNil(o.Loss) {
		var ret float32
		return ret
	}
	return *o.Loss
}

// GetLossOk returns a tuple with the Loss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetLossOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Loss) {
		return nil, false
	}
	return o.Loss, true
}

// HasLoss returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasLoss() bool {
	if o != nil && !utils.IsNil(o.Loss) {
		return true
	}

	return false
}

// SetLoss gets a reference to the given float32 and assigns it to the Loss field.
func (o *RtpStreamTestResult) SetLoss(v float32) {
	o.Loss = &v
}

// GetDiscards returns the Discards field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetDiscards() float32 {
	if o == nil || utils.IsNil(o.Discards) {
		var ret float32
		return ret
	}
	return *o.Discards
}

// GetDiscardsOk returns a tuple with the Discards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetDiscardsOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Discards) {
		return nil, false
	}
	return o.Discards, true
}

// HasDiscards returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasDiscards() bool {
	if o != nil && !utils.IsNil(o.Discards) {
		return true
	}

	return false
}

// SetDiscards gets a reference to the given float32 and assigns it to the Discards field.
func (o *RtpStreamTestResult) SetDiscards(v float32) {
	o.Discards = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetLatency() int32 {
	if o == nil || utils.IsNil(o.Latency) {
		var ret int32
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetLatencyOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Latency) {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasLatency() bool {
	if o != nil && !utils.IsNil(o.Latency) {
		return true
	}

	return false
}

// SetLatency gets a reference to the given int32 and assigns it to the Latency field.
func (o *RtpStreamTestResult) SetLatency(v int32) {
	o.Latency = &v
}

// GetPdv returns the Pdv field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetPdv() int32 {
	if o == nil || utils.IsNil(o.Pdv) {
		var ret int32
		return ret
	}
	return *o.Pdv
}

// GetPdvOk returns a tuple with the Pdv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetPdvOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Pdv) {
		return nil, false
	}
	return o.Pdv, true
}

// HasPdv returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasPdv() bool {
	if o != nil && !utils.IsNil(o.Pdv) {
		return true
	}

	return false
}

// SetPdv gets a reference to the given int32 and assigns it to the Pdv field.
func (o *RtpStreamTestResult) SetPdv(v int32) {
	o.Pdv = &v
}

// GetErrorDetail returns the ErrorDetail field value if set, zero value otherwise.
func (o *RtpStreamTestResult) GetErrorDetail() string {
	if o == nil || utils.IsNil(o.ErrorDetail) {
		var ret string
		return ret
	}
	return *o.ErrorDetail
}

// GetErrorDetailOk returns a tuple with the ErrorDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpStreamTestResult) GetErrorDetailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorDetail) {
		return nil, false
	}
	return o.ErrorDetail, true
}

// HasErrorDetail returns a boolean if a field has been set.
func (o *RtpStreamTestResult) HasErrorDetail() bool {
	if o != nil && !utils.IsNil(o.ErrorDetail) {
		return true
	}

	return false
}

// SetErrorDetail gets a reference to the given string and assigns it to the ErrorDetail field.
func (o *RtpStreamTestResult) SetErrorDetail(v string) {
	o.ErrorDetail = &v
}

func (o RtpStreamTestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RtpStreamTestResult) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.ServerIp) {
		toSerialize["serverIp"] = o.ServerIp
	}
	if !utils.IsNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	if !utils.IsNil(o.DscpName) {
		toSerialize["dscpName"] = o.DscpName
	}
	if !utils.IsNil(o.Mos) {
		toSerialize["mos"] = o.Mos
	}
	if !utils.IsNil(o.CodecName) {
		toSerialize["codecName"] = o.CodecName
	}
	if !utils.IsNil(o.CodecMaxMos) {
		toSerialize["codecMaxMos"] = o.CodecMaxMos
	}
	if !utils.IsNil(o.Loss) {
		toSerialize["loss"] = o.Loss
	}
	if !utils.IsNil(o.Discards) {
		toSerialize["discards"] = o.Discards
	}
	if !utils.IsNil(o.Latency) {
		toSerialize["latency"] = o.Latency
	}
	if !utils.IsNil(o.Pdv) {
		toSerialize["pdv"] = o.Pdv
	}
	if !utils.IsNil(o.ErrorDetail) {
		toSerialize["errorDetail"] = o.ErrorDetail
	}
	return toSerialize, nil
}

type NullableRtpStreamTestResult struct {
	value *RtpStreamTestResult
	isSet bool
}

func (v NullableRtpStreamTestResult) Get() *RtpStreamTestResult {
	return v.value
}

func (v *NullableRtpStreamTestResult) Set(val *RtpStreamTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRtpStreamTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRtpStreamTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRtpStreamTestResult(val *RtpStreamTestResult) *NullableRtpStreamTestResult {
	return &NullableRtpStreamTestResult{value: val, isSet: true}
}

func (v NullableRtpStreamTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRtpStreamTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


