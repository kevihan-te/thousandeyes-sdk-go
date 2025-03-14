/*
Instant Tests API

The Instant Tests API operations lets you create and run new instant tests. You will need to be a regular user or have the following permissions:   * `API Access`   * `View tests`  The response does not include the immediate test results. Use the Test Results endpoints to get test results after creating and executing an instant test. You can find the URLs for these endpoints in the _links section of the test definition that is returned when you create the instant test. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instanttests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"fmt"
)

// checks if the InstantTestRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InstantTestRequest{}

// InstantTestRequest struct for InstantTestRequest
type InstantTestRequest struct {
	// A list of test label identifiers (get `labelId` from `/labels` endpoint).
	Labels []string `json:"labels,omitempty"`
	// A list of account group identifiers that the test is shared with (get `aid` from `/account-groups` endpoint).
	SharedWithAccounts []string `json:"sharedWithAccounts,omitempty"`
	// A list of objects with `agentId` (required) and `sourceIpAddress` (optional).
	Agents []TestAgent `json:"agents"`
}

type _InstantTestRequest InstantTestRequest

// NewInstantTestRequest instantiates a new InstantTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstantTestRequest(agents []TestAgent) *InstantTestRequest {
	this := InstantTestRequest{}
	this.Agents = agents
	return &this
}

// NewInstantTestRequestWithDefaults instantiates a new InstantTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstantTestRequestWithDefaults() *InstantTestRequest {
	this := InstantTestRequest{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *InstantTestRequest) GetLabels() []string {
	if o == nil || utils.IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantTestRequest) GetLabelsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *InstantTestRequest) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *InstantTestRequest) SetLabels(v []string) {
	o.Labels = v
}

// GetSharedWithAccounts returns the SharedWithAccounts field value if set, zero value otherwise.
func (o *InstantTestRequest) GetSharedWithAccounts() []string {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		var ret []string
		return ret
	}
	return o.SharedWithAccounts
}

// GetSharedWithAccountsOk returns a tuple with the SharedWithAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantTestRequest) GetSharedWithAccountsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		return nil, false
	}
	return o.SharedWithAccounts, true
}

// HasSharedWithAccounts returns a boolean if a field has been set.
func (o *InstantTestRequest) HasSharedWithAccounts() bool {
	if o != nil && !utils.IsNil(o.SharedWithAccounts) {
		return true
	}

	return false
}

// SetSharedWithAccounts gets a reference to the given []string and assigns it to the SharedWithAccounts field.
func (o *InstantTestRequest) SetSharedWithAccounts(v []string) {
	o.SharedWithAccounts = v
}

// GetAgents returns the Agents field value
func (o *InstantTestRequest) GetAgents() []TestAgent {
	if o == nil {
		var ret []TestAgent
		return ret
	}

	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value
// and a boolean to check if the value has been set.
func (o *InstantTestRequest) GetAgentsOk() ([]TestAgent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Agents, true
}

// SetAgents sets field value
func (o *InstantTestRequest) SetAgents(v []TestAgent) {
	o.Agents = v
}

func (o InstantTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstantTestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !utils.IsNil(o.SharedWithAccounts) {
		toSerialize["sharedWithAccounts"] = o.SharedWithAccounts
	}
	toSerialize["agents"] = o.Agents
	return toSerialize, nil
}

func (o *InstantTestRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agents",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInstantTestRequest := _InstantTestRequest{}

    err = json.Unmarshal(data, &varInstantTestRequest)

	if err != nil {
		return err
	}

	*o = InstantTestRequest(varInstantTestRequest)

	return err
}

type NullableInstantTestRequest struct {
	value *InstantTestRequest
	isSet bool
}

func (v NullableInstantTestRequest) Get() *InstantTestRequest {
	return v.value
}

func (v *NullableInstantTestRequest) Set(val *InstantTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInstantTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInstantTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstantTestRequest(val *InstantTestRequest) *NullableInstantTestRequest {
	return &NullableInstantTestRequest{value: val, isSet: true}
}

func (v NullableInstantTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstantTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


