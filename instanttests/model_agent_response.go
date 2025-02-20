/*
Instant Tests API

The Instant Tests API operations lets you create and run new instant tests. You will need to be a regular user or have the following permissions:   * `API Access`   * `View tests`  The response does not include the immediate test results. Use the Test Results endpoints to get test results after creating and executing an instant test. You can find the URLs for these endpoints in the _links section of the test definition that is returned when you create the instant test. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instanttests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"fmt"
)

// checks if the AgentResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AgentResponse{}

// AgentResponse struct for AgentResponse
type AgentResponse struct {
	// Array of private IP addresses.
	IpAddresses []string `json:"ipAddresses,omitempty"`
	// Array of public IP addresses.
	PublicIpAddresses []string `json:"publicIpAddresses,omitempty"`
	// Network (including ASN) of agent’s public IP.
	Network *string `json:"network,omitempty"`
	// Unique ID of the agent.
	AgentId *string `json:"agentId,omitempty"`
	// Name of the agent.
	AgentName *string `json:"agentName,omitempty"`
	// Location of the agent.
	Location *string `json:"location,omitempty"`
	// 2-digit ISO country code
	CountryId *string `json:"countryId,omitempty"`
	// Flag indicating if the agent is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Prefix containing agents public IP address.
	Prefix *string `json:"prefix,omitempty"`
	// Flag indicating if has normal SSL operations or  if instead it's set to ignore SSL errors on browserbot-based tests.
	VerifySslCertificates *bool `json:"verifySslCertificates,omitempty"`
	AgentType CloudEnterpriseAgentType `json:"agentType"`
}

type _AgentResponse AgentResponse

// NewAgentResponse instantiates a new AgentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentResponse(agentType CloudEnterpriseAgentType) *AgentResponse {
	this := AgentResponse{}
	this.AgentType = agentType
	return &this
}

// NewAgentResponseWithDefaults instantiates a new AgentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentResponseWithDefaults() *AgentResponse {
	this := AgentResponse{}
	return &this
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *AgentResponse) GetIpAddresses() []string {
	if o == nil || utils.IsNil(o.IpAddresses) {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentResponse) GetIpAddressesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.IpAddresses) {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *AgentResponse) HasIpAddresses() bool {
	if o != nil && !utils.IsNil(o.IpAddresses) {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *AgentResponse) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetPublicIpAddresses returns the PublicIpAddresses field value if set, zero value otherwise.
func (o *AgentResponse) GetPublicIpAddresses() []string {
	if o == nil || utils.IsNil(o.PublicIpAddresses) {
		var ret []string
		return ret
	}
	return o.PublicIpAddresses
}

// GetPublicIpAddressesOk returns a tuple with the PublicIpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentResponse) GetPublicIpAddressesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.PublicIpAddresses) {
		return nil, false
	}
	return o.PublicIpAddresses, true
}

// HasPublicIpAddresses returns a boolean if a field has been set.
func (o *AgentResponse) HasPublicIpAddresses() bool {
	if o != nil && !utils.IsNil(o.PublicIpAddresses) {
		return true
	}

	return false
}

// SetPublicIpAddresses gets a reference to the given []string and assigns it to the PublicIpAddresses field.
func (o *AgentResponse) SetPublicIpAddresses(v []string) {
	o.PublicIpAddresses = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *AgentResponse) GetNetwork() string {
	if o == nil || utils.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentResponse) GetNetworkOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *AgentResponse) HasNetwork() bool {
	if o != nil && !utils.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *AgentResponse) SetNetwork(v string) {
	o.Network = &v
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *AgentResponse) GetAgentId() string {
	if o == nil || utils.IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentResponse) GetAgentIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *AgentResponse) HasAgentId() bool {
	if o != nil && !utils.IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *AgentResponse) SetAgentId(v string) {
	o.AgentId = &v
}

// GetAgentName returns the AgentName field value if set, zero value otherwise.
func (o *AgentResponse) GetAgentName() string {
	if o == nil || utils.IsNil(o.AgentName) {
		var ret string
		return ret
	}
	return *o.AgentName
}

// GetAgentNameOk returns a tuple with the AgentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentResponse) GetAgentNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AgentName) {
		return nil, false
	}
	return o.AgentName, true
}

// HasAgentName returns a boolean if a field has been set.
func (o *AgentResponse) HasAgentName() bool {
	if o != nil && !utils.IsNil(o.AgentName) {
		return true
	}

	return false
}

// SetAgentName gets a reference to the given string and assigns it to the AgentName field.
func (o *AgentResponse) SetAgentName(v string) {
	o.AgentName = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AgentResponse) GetLocation() string {
	if o == nil || utils.IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentResponse) GetLocationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AgentResponse) HasLocation() bool {
	if o != nil && !utils.IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *AgentResponse) SetLocation(v string) {
	o.Location = &v
}

// GetCountryId returns the CountryId field value if set, zero value otherwise.
func (o *AgentResponse) GetCountryId() string {
	if o == nil || utils.IsNil(o.CountryId) {
		var ret string
		return ret
	}
	return *o.CountryId
}

// GetCountryIdOk returns a tuple with the CountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentResponse) GetCountryIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CountryId) {
		return nil, false
	}
	return o.CountryId, true
}

// HasCountryId returns a boolean if a field has been set.
func (o *AgentResponse) HasCountryId() bool {
	if o != nil && !utils.IsNil(o.CountryId) {
		return true
	}

	return false
}

// SetCountryId gets a reference to the given string and assigns it to the CountryId field.
func (o *AgentResponse) SetCountryId(v string) {
	o.CountryId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AgentResponse) GetEnabled() bool {
	if o == nil || utils.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AgentResponse) HasEnabled() bool {
	if o != nil && !utils.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AgentResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *AgentResponse) GetPrefix() string {
	if o == nil || utils.IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentResponse) GetPrefixOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *AgentResponse) HasPrefix() bool {
	if o != nil && !utils.IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *AgentResponse) SetPrefix(v string) {
	o.Prefix = &v
}

// GetVerifySslCertificates returns the VerifySslCertificates field value if set, zero value otherwise.
func (o *AgentResponse) GetVerifySslCertificates() bool {
	if o == nil || utils.IsNil(o.VerifySslCertificates) {
		var ret bool
		return ret
	}
	return *o.VerifySslCertificates
}

// GetVerifySslCertificatesOk returns a tuple with the VerifySslCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentResponse) GetVerifySslCertificatesOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.VerifySslCertificates) {
		return nil, false
	}
	return o.VerifySslCertificates, true
}

// HasVerifySslCertificates returns a boolean if a field has been set.
func (o *AgentResponse) HasVerifySslCertificates() bool {
	if o != nil && !utils.IsNil(o.VerifySslCertificates) {
		return true
	}

	return false
}

// SetVerifySslCertificates gets a reference to the given bool and assigns it to the VerifySslCertificates field.
func (o *AgentResponse) SetVerifySslCertificates(v bool) {
	o.VerifySslCertificates = &v
}

// GetAgentType returns the AgentType field value
func (o *AgentResponse) GetAgentType() CloudEnterpriseAgentType {
	if o == nil {
		var ret CloudEnterpriseAgentType
		return ret
	}

	return o.AgentType
}

// GetAgentTypeOk returns a tuple with the AgentType field value
// and a boolean to check if the value has been set.
func (o *AgentResponse) GetAgentTypeOk() (*CloudEnterpriseAgentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentType, true
}

// SetAgentType sets field value
func (o *AgentResponse) SetAgentType(v CloudEnterpriseAgentType) {
	o.AgentType = v
}

func (o AgentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.IpAddresses) {
		toSerialize["ipAddresses"] = o.IpAddresses
	}
	if !utils.IsNil(o.PublicIpAddresses) {
		toSerialize["publicIpAddresses"] = o.PublicIpAddresses
	}
	if !utils.IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !utils.IsNil(o.AgentId) {
		toSerialize["agentId"] = o.AgentId
	}
	if !utils.IsNil(o.AgentName) {
		toSerialize["agentName"] = o.AgentName
	}
	if !utils.IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !utils.IsNil(o.CountryId) {
		toSerialize["countryId"] = o.CountryId
	}
	if !utils.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !utils.IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !utils.IsNil(o.VerifySslCertificates) {
		toSerialize["verifySslCertificates"] = o.VerifySslCertificates
	}
	toSerialize["agentType"] = o.AgentType
	return toSerialize, nil
}

func (o *AgentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agentType",
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

	varAgentResponse := _AgentResponse{}

    err = json.Unmarshal(data, &varAgentResponse)

	if err != nil {
		return err
	}

	*o = AgentResponse(varAgentResponse)

	return err
}

type NullableAgentResponse struct {
	value *AgentResponse
	isSet bool
}

func (v NullableAgentResponse) Get() *AgentResponse {
	return v.value
}

func (v *NullableAgentResponse) Set(val *AgentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentResponse(val *AgentResponse) *NullableAgentResponse {
	return &NullableAgentResponse{value: val, isSet: true}
}

func (v NullableAgentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


