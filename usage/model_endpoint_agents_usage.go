/*
Usage API

 These usage endpoints define the following operations:  * **Usage**: Retrieve usage data for the specified time period (default is one month).          * Users must have the `View organization usage` permission to access this endpoint.     * This operation offers visibility across all account groups within the organization.     * Users with `View organization usage` permission in multiple organizations should query the operation with the `aid` query string parameter (see optional parameters) for each organization.  * **Quotas**: Obtain organization and account usage quotas. Additionally, users with the appropriate permissions can create, update, or delete these quotas.          * Users must have the necessary permissions to perform quota-related actions.  Refer to the Usage API operations for detailed usage instructions and optional parameters. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package usage

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the EndpointAgentsUsage type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointAgentsUsage{}

// EndpointAgentsUsage struct for EndpointAgentsUsage
type EndpointAgentsUsage struct {
	// Unique identifier of the account group owning the endpoint agents.
	Aid *string `json:"aid,omitempty"`
	// Name of the account group which owns the endpoint agents.
	AccountGroupName *string `json:"accountGroupName,omitempty"`
	// Number of endpoint agents owned by the specific account group in the usage period.
	EndpointAgentsUsed *int64 `json:"endpointAgentsUsed,omitempty"`
}

// NewEndpointAgentsUsage instantiates a new EndpointAgentsUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointAgentsUsage() *EndpointAgentsUsage {
	this := EndpointAgentsUsage{}
	return &this
}

// NewEndpointAgentsUsageWithDefaults instantiates a new EndpointAgentsUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointAgentsUsageWithDefaults() *EndpointAgentsUsage {
	this := EndpointAgentsUsage{}
	return &this
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *EndpointAgentsUsage) GetAid() string {
	if o == nil || utils.IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgentsUsage) GetAidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *EndpointAgentsUsage) HasAid() bool {
	if o != nil && !utils.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *EndpointAgentsUsage) SetAid(v string) {
	o.Aid = &v
}

// GetAccountGroupName returns the AccountGroupName field value if set, zero value otherwise.
func (o *EndpointAgentsUsage) GetAccountGroupName() string {
	if o == nil || utils.IsNil(o.AccountGroupName) {
		var ret string
		return ret
	}
	return *o.AccountGroupName
}

// GetAccountGroupNameOk returns a tuple with the AccountGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgentsUsage) GetAccountGroupNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccountGroupName) {
		return nil, false
	}
	return o.AccountGroupName, true
}

// HasAccountGroupName returns a boolean if a field has been set.
func (o *EndpointAgentsUsage) HasAccountGroupName() bool {
	if o != nil && !utils.IsNil(o.AccountGroupName) {
		return true
	}

	return false
}

// SetAccountGroupName gets a reference to the given string and assigns it to the AccountGroupName field.
func (o *EndpointAgentsUsage) SetAccountGroupName(v string) {
	o.AccountGroupName = &v
}

// GetEndpointAgentsUsed returns the EndpointAgentsUsed field value if set, zero value otherwise.
func (o *EndpointAgentsUsage) GetEndpointAgentsUsed() int64 {
	if o == nil || utils.IsNil(o.EndpointAgentsUsed) {
		var ret int64
		return ret
	}
	return *o.EndpointAgentsUsed
}

// GetEndpointAgentsUsedOk returns a tuple with the EndpointAgentsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgentsUsage) GetEndpointAgentsUsedOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.EndpointAgentsUsed) {
		return nil, false
	}
	return o.EndpointAgentsUsed, true
}

// HasEndpointAgentsUsed returns a boolean if a field has been set.
func (o *EndpointAgentsUsage) HasEndpointAgentsUsed() bool {
	if o != nil && !utils.IsNil(o.EndpointAgentsUsed) {
		return true
	}

	return false
}

// SetEndpointAgentsUsed gets a reference to the given int64 and assigns it to the EndpointAgentsUsed field.
func (o *EndpointAgentsUsage) SetEndpointAgentsUsed(v int64) {
	o.EndpointAgentsUsed = &v
}

func (o EndpointAgentsUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointAgentsUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	if !utils.IsNil(o.AccountGroupName) {
		toSerialize["accountGroupName"] = o.AccountGroupName
	}
	if !utils.IsNil(o.EndpointAgentsUsed) {
		toSerialize["endpointAgentsUsed"] = o.EndpointAgentsUsed
	}
	return toSerialize, nil
}

type NullableEndpointAgentsUsage struct {
	value *EndpointAgentsUsage
	isSet bool
}

func (v NullableEndpointAgentsUsage) Get() *EndpointAgentsUsage {
	return v.value
}

func (v *NullableEndpointAgentsUsage) Set(val *EndpointAgentsUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointAgentsUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointAgentsUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointAgentsUsage(val *EndpointAgentsUsage) *NullableEndpointAgentsUsage {
	return &NullableEndpointAgentsUsage{value: val, isSet: true}
}

func (v NullableEndpointAgentsUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointAgentsUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


