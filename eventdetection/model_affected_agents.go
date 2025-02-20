/*
Event Detection API

 Event detection occurs when ThousandEyes identifies that error signals related to a component (proxy, network node, AS, server etc) have deviated from the baselines established by events. * To determine this, ThousandEyes takes the test results from all accounts groups within an organization, and analyzes that data. * Noisy test results (those that have too many errors in a short window) are removed until they stabilize, and the rest of the results are tagged with the components associated with that test result (for example, proxy, network, or server). * Next, any increase in failures from the test results and each component helps in determining the problem domain and which component may be at fault. * When this failure rate increases beyond a pre-defined threshold (set by the algorithm), an event is triggered and an email notification is sent to the user (if they've enabled email alerts).  With the Events API, you can perform the following tasks on the ThousandEyes platform: * **Retrieve Events**: Obtain a list of events and detailed information for each event. For more information about events, see [Event Detection](https://docs.thousandeyes.com/product-documentation/event-detection). 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eventdetection

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the AffectedAgents type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffectedAgents{}

// AffectedAgents struct for AffectedAgents
type AffectedAgents struct {
	// The total number affected.
	Total *int32 `json:"total,omitempty"`
	// Indicates if in the affected account group.
	InAccountGroup *int32 `json:"inAccountGroup,omitempty"`
	// List of affected agents.
	Agents []EventApiAffectedAgent `json:"agents,omitempty"`
}

// NewAffectedAgents instantiates a new AffectedAgents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffectedAgents() *AffectedAgents {
	this := AffectedAgents{}
	return &this
}

// NewAffectedAgentsWithDefaults instantiates a new AffectedAgents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffectedAgentsWithDefaults() *AffectedAgents {
	this := AffectedAgents{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *AffectedAgents) GetTotal() int32 {
	if o == nil || utils.IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedAgents) GetTotalOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *AffectedAgents) HasTotal() bool {
	if o != nil && !utils.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *AffectedAgents) SetTotal(v int32) {
	o.Total = &v
}

// GetInAccountGroup returns the InAccountGroup field value if set, zero value otherwise.
func (o *AffectedAgents) GetInAccountGroup() int32 {
	if o == nil || utils.IsNil(o.InAccountGroup) {
		var ret int32
		return ret
	}
	return *o.InAccountGroup
}

// GetInAccountGroupOk returns a tuple with the InAccountGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedAgents) GetInAccountGroupOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.InAccountGroup) {
		return nil, false
	}
	return o.InAccountGroup, true
}

// HasInAccountGroup returns a boolean if a field has been set.
func (o *AffectedAgents) HasInAccountGroup() bool {
	if o != nil && !utils.IsNil(o.InAccountGroup) {
		return true
	}

	return false
}

// SetInAccountGroup gets a reference to the given int32 and assigns it to the InAccountGroup field.
func (o *AffectedAgents) SetInAccountGroup(v int32) {
	o.InAccountGroup = &v
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *AffectedAgents) GetAgents() []EventApiAffectedAgent {
	if o == nil || utils.IsNil(o.Agents) {
		var ret []EventApiAffectedAgent
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedAgents) GetAgentsOk() ([]EventApiAffectedAgent, bool) {
	if o == nil || utils.IsNil(o.Agents) {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *AffectedAgents) HasAgents() bool {
	if o != nil && !utils.IsNil(o.Agents) {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []EventApiAffectedAgent and assigns it to the Agents field.
func (o *AffectedAgents) SetAgents(v []EventApiAffectedAgent) {
	o.Agents = v
}

func (o AffectedAgents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AffectedAgents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !utils.IsNil(o.InAccountGroup) {
		toSerialize["inAccountGroup"] = o.InAccountGroup
	}
	if !utils.IsNil(o.Agents) {
		toSerialize["agents"] = o.Agents
	}
	return toSerialize, nil
}

type NullableAffectedAgents struct {
	value *AffectedAgents
	isSet bool
}

func (v NullableAffectedAgents) Get() *AffectedAgents {
	return v.value
}

func (v *NullableAffectedAgents) Set(val *AffectedAgents) {
	v.value = val
	v.isSet = true
}

func (v NullableAffectedAgents) IsSet() bool {
	return v.isSet
}

func (v *NullableAffectedAgents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffectedAgents(val *AffectedAgents) *NullableAffectedAgents {
	return &NullableAffectedAgents{value: val, isSet: true}
}

func (v NullableAffectedAgents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffectedAgents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


