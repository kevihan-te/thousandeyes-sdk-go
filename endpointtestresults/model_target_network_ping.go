/*
Endpoint Test Results API

Retrieve results for scheduled and dynamic tests on endpoint agents.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointtestresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the TargetNetworkPing type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TargetNetworkPing{}

// TargetNetworkPing struct for TargetNetworkPing
type TargetNetworkPing struct {
	// Ping average response time.
	AvgRtt *int32 `json:"avgRtt,omitempty"`
	// Ping maximum response time.
	MaxRtt *int32 `json:"maxRtt,omitempty"`
	// Ping mean standard deviation response time.
	MeanDevRtt *int32 `json:"meanDevRtt,omitempty"`
	// Ping minimum response time.
	MinRtt *int32 `json:"minRtt,omitempty"`
	// Ping packets received.
	PktsReceived *int32 `json:"pktsReceived,omitempty"`
	// Ping packets sent.
	PktsSent *int32 `json:"pktsSent,omitempty"`
	// Only present when there is an error.
	Error *string `json:"error,omitempty"`
	InfoFlags []string `json:"infoFlags,omitempty"`
}

// NewTargetNetworkPing instantiates a new TargetNetworkPing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetNetworkPing() *TargetNetworkPing {
	this := TargetNetworkPing{}
	return &this
}

// NewTargetNetworkPingWithDefaults instantiates a new TargetNetworkPing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetNetworkPingWithDefaults() *TargetNetworkPing {
	this := TargetNetworkPing{}
	return &this
}

// GetAvgRtt returns the AvgRtt field value if set, zero value otherwise.
func (o *TargetNetworkPing) GetAvgRtt() int32 {
	if o == nil || utils.IsNil(o.AvgRtt) {
		var ret int32
		return ret
	}
	return *o.AvgRtt
}

// GetAvgRttOk returns a tuple with the AvgRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkPing) GetAvgRttOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.AvgRtt) {
		return nil, false
	}
	return o.AvgRtt, true
}

// HasAvgRtt returns a boolean if a field has been set.
func (o *TargetNetworkPing) HasAvgRtt() bool {
	if o != nil && !utils.IsNil(o.AvgRtt) {
		return true
	}

	return false
}

// SetAvgRtt gets a reference to the given int32 and assigns it to the AvgRtt field.
func (o *TargetNetworkPing) SetAvgRtt(v int32) {
	o.AvgRtt = &v
}

// GetMaxRtt returns the MaxRtt field value if set, zero value otherwise.
func (o *TargetNetworkPing) GetMaxRtt() int32 {
	if o == nil || utils.IsNil(o.MaxRtt) {
		var ret int32
		return ret
	}
	return *o.MaxRtt
}

// GetMaxRttOk returns a tuple with the MaxRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkPing) GetMaxRttOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MaxRtt) {
		return nil, false
	}
	return o.MaxRtt, true
}

// HasMaxRtt returns a boolean if a field has been set.
func (o *TargetNetworkPing) HasMaxRtt() bool {
	if o != nil && !utils.IsNil(o.MaxRtt) {
		return true
	}

	return false
}

// SetMaxRtt gets a reference to the given int32 and assigns it to the MaxRtt field.
func (o *TargetNetworkPing) SetMaxRtt(v int32) {
	o.MaxRtt = &v
}

// GetMeanDevRtt returns the MeanDevRtt field value if set, zero value otherwise.
func (o *TargetNetworkPing) GetMeanDevRtt() int32 {
	if o == nil || utils.IsNil(o.MeanDevRtt) {
		var ret int32
		return ret
	}
	return *o.MeanDevRtt
}

// GetMeanDevRttOk returns a tuple with the MeanDevRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkPing) GetMeanDevRttOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MeanDevRtt) {
		return nil, false
	}
	return o.MeanDevRtt, true
}

// HasMeanDevRtt returns a boolean if a field has been set.
func (o *TargetNetworkPing) HasMeanDevRtt() bool {
	if o != nil && !utils.IsNil(o.MeanDevRtt) {
		return true
	}

	return false
}

// SetMeanDevRtt gets a reference to the given int32 and assigns it to the MeanDevRtt field.
func (o *TargetNetworkPing) SetMeanDevRtt(v int32) {
	o.MeanDevRtt = &v
}

// GetMinRtt returns the MinRtt field value if set, zero value otherwise.
func (o *TargetNetworkPing) GetMinRtt() int32 {
	if o == nil || utils.IsNil(o.MinRtt) {
		var ret int32
		return ret
	}
	return *o.MinRtt
}

// GetMinRttOk returns a tuple with the MinRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkPing) GetMinRttOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MinRtt) {
		return nil, false
	}
	return o.MinRtt, true
}

// HasMinRtt returns a boolean if a field has been set.
func (o *TargetNetworkPing) HasMinRtt() bool {
	if o != nil && !utils.IsNil(o.MinRtt) {
		return true
	}

	return false
}

// SetMinRtt gets a reference to the given int32 and assigns it to the MinRtt field.
func (o *TargetNetworkPing) SetMinRtt(v int32) {
	o.MinRtt = &v
}

// GetPktsReceived returns the PktsReceived field value if set, zero value otherwise.
func (o *TargetNetworkPing) GetPktsReceived() int32 {
	if o == nil || utils.IsNil(o.PktsReceived) {
		var ret int32
		return ret
	}
	return *o.PktsReceived
}

// GetPktsReceivedOk returns a tuple with the PktsReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkPing) GetPktsReceivedOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PktsReceived) {
		return nil, false
	}
	return o.PktsReceived, true
}

// HasPktsReceived returns a boolean if a field has been set.
func (o *TargetNetworkPing) HasPktsReceived() bool {
	if o != nil && !utils.IsNil(o.PktsReceived) {
		return true
	}

	return false
}

// SetPktsReceived gets a reference to the given int32 and assigns it to the PktsReceived field.
func (o *TargetNetworkPing) SetPktsReceived(v int32) {
	o.PktsReceived = &v
}

// GetPktsSent returns the PktsSent field value if set, zero value otherwise.
func (o *TargetNetworkPing) GetPktsSent() int32 {
	if o == nil || utils.IsNil(o.PktsSent) {
		var ret int32
		return ret
	}
	return *o.PktsSent
}

// GetPktsSentOk returns a tuple with the PktsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkPing) GetPktsSentOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PktsSent) {
		return nil, false
	}
	return o.PktsSent, true
}

// HasPktsSent returns a boolean if a field has been set.
func (o *TargetNetworkPing) HasPktsSent() bool {
	if o != nil && !utils.IsNil(o.PktsSent) {
		return true
	}

	return false
}

// SetPktsSent gets a reference to the given int32 and assigns it to the PktsSent field.
func (o *TargetNetworkPing) SetPktsSent(v int32) {
	o.PktsSent = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *TargetNetworkPing) GetError() string {
	if o == nil || utils.IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkPing) GetErrorOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *TargetNetworkPing) HasError() bool {
	if o != nil && !utils.IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *TargetNetworkPing) SetError(v string) {
	o.Error = &v
}

// GetInfoFlags returns the InfoFlags field value if set, zero value otherwise.
func (o *TargetNetworkPing) GetInfoFlags() []string {
	if o == nil || utils.IsNil(o.InfoFlags) {
		var ret []string
		return ret
	}
	return o.InfoFlags
}

// GetInfoFlagsOk returns a tuple with the InfoFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkPing) GetInfoFlagsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.InfoFlags) {
		return nil, false
	}
	return o.InfoFlags, true
}

// HasInfoFlags returns a boolean if a field has been set.
func (o *TargetNetworkPing) HasInfoFlags() bool {
	if o != nil && !utils.IsNil(o.InfoFlags) {
		return true
	}

	return false
}

// SetInfoFlags gets a reference to the given []string and assigns it to the InfoFlags field.
func (o *TargetNetworkPing) SetInfoFlags(v []string) {
	o.InfoFlags = v
}

func (o TargetNetworkPing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetNetworkPing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AvgRtt) {
		toSerialize["avgRtt"] = o.AvgRtt
	}
	if !utils.IsNil(o.MaxRtt) {
		toSerialize["maxRtt"] = o.MaxRtt
	}
	if !utils.IsNil(o.MeanDevRtt) {
		toSerialize["meanDevRtt"] = o.MeanDevRtt
	}
	if !utils.IsNil(o.MinRtt) {
		toSerialize["minRtt"] = o.MinRtt
	}
	if !utils.IsNil(o.PktsReceived) {
		toSerialize["pktsReceived"] = o.PktsReceived
	}
	if !utils.IsNil(o.PktsSent) {
		toSerialize["pktsSent"] = o.PktsSent
	}
	if !utils.IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !utils.IsNil(o.InfoFlags) {
		toSerialize["infoFlags"] = o.InfoFlags
	}
	return toSerialize, nil
}

type NullableTargetNetworkPing struct {
	value *TargetNetworkPing
	isSet bool
}

func (v NullableTargetNetworkPing) Get() *TargetNetworkPing {
	return v.value
}

func (v *NullableTargetNetworkPing) Set(val *TargetNetworkPing) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetNetworkPing) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetNetworkPing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetNetworkPing(val *TargetNetworkPing) *NullableTargetNetworkPing {
	return &NullableTargetNetworkPing{value: val, isSet: true}
}

func (v NullableTargetNetworkPing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetNetworkPing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


