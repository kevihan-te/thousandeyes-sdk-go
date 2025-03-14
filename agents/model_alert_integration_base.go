/*
Agents API

 ## Overview Manage Cloud and Enterprise Agents available to your account in ThousandEyes.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package agents

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the AlertIntegrationBase type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AlertIntegrationBase{}

// AlertIntegrationBase struct for AlertIntegrationBase
type AlertIntegrationBase struct {
	// Unique ID of the integration.
	IntegrationId *string `json:"integrationId,omitempty"`
	// Name of the integration.
	IntegrationName *string `json:"integrationName,omitempty"`
	IntegrationType *AlertIntegrationType `json:"integrationType,omitempty"`
	// Target URL of the integration.
	Target *string `json:"target,omitempty"`
	// (PagerDuty only) Authentication method.
	AuthMethod *string `json:"authMethod,omitempty"`
	// (PagerDuty only) Authentication user.
	AuthUser *string `json:"authUser,omitempty"`
	// (PagerDuty only) Authentication token.
	AuthToken *string `json:"authToken,omitempty"`
	// (Slack only) Slack `#channel` or `@user`.
	Channel *string `json:"channel,omitempty"`
}

// NewAlertIntegrationBase instantiates a new AlertIntegrationBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertIntegrationBase() *AlertIntegrationBase {
	this := AlertIntegrationBase{}
	return &this
}

// NewAlertIntegrationBaseWithDefaults instantiates a new AlertIntegrationBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertIntegrationBaseWithDefaults() *AlertIntegrationBase {
	this := AlertIntegrationBase{}
	return &this
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *AlertIntegrationBase) GetIntegrationId() string {
	if o == nil || utils.IsNil(o.IntegrationId) {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertIntegrationBase) GetIntegrationIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IntegrationId) {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *AlertIntegrationBase) HasIntegrationId() bool {
	if o != nil && !utils.IsNil(o.IntegrationId) {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *AlertIntegrationBase) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

// GetIntegrationName returns the IntegrationName field value if set, zero value otherwise.
func (o *AlertIntegrationBase) GetIntegrationName() string {
	if o == nil || utils.IsNil(o.IntegrationName) {
		var ret string
		return ret
	}
	return *o.IntegrationName
}

// GetIntegrationNameOk returns a tuple with the IntegrationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertIntegrationBase) GetIntegrationNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IntegrationName) {
		return nil, false
	}
	return o.IntegrationName, true
}

// HasIntegrationName returns a boolean if a field has been set.
func (o *AlertIntegrationBase) HasIntegrationName() bool {
	if o != nil && !utils.IsNil(o.IntegrationName) {
		return true
	}

	return false
}

// SetIntegrationName gets a reference to the given string and assigns it to the IntegrationName field.
func (o *AlertIntegrationBase) SetIntegrationName(v string) {
	o.IntegrationName = &v
}

// GetIntegrationType returns the IntegrationType field value if set, zero value otherwise.
func (o *AlertIntegrationBase) GetIntegrationType() AlertIntegrationType {
	if o == nil || utils.IsNil(o.IntegrationType) {
		var ret AlertIntegrationType
		return ret
	}
	return *o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertIntegrationBase) GetIntegrationTypeOk() (*AlertIntegrationType, bool) {
	if o == nil || utils.IsNil(o.IntegrationType) {
		return nil, false
	}
	return o.IntegrationType, true
}

// HasIntegrationType returns a boolean if a field has been set.
func (o *AlertIntegrationBase) HasIntegrationType() bool {
	if o != nil && !utils.IsNil(o.IntegrationType) {
		return true
	}

	return false
}

// SetIntegrationType gets a reference to the given AlertIntegrationType and assigns it to the IntegrationType field.
func (o *AlertIntegrationBase) SetIntegrationType(v AlertIntegrationType) {
	o.IntegrationType = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *AlertIntegrationBase) GetTarget() string {
	if o == nil || utils.IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertIntegrationBase) GetTargetOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *AlertIntegrationBase) HasTarget() bool {
	if o != nil && !utils.IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *AlertIntegrationBase) SetTarget(v string) {
	o.Target = &v
}

// GetAuthMethod returns the AuthMethod field value if set, zero value otherwise.
func (o *AlertIntegrationBase) GetAuthMethod() string {
	if o == nil || utils.IsNil(o.AuthMethod) {
		var ret string
		return ret
	}
	return *o.AuthMethod
}

// GetAuthMethodOk returns a tuple with the AuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertIntegrationBase) GetAuthMethodOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AuthMethod) {
		return nil, false
	}
	return o.AuthMethod, true
}

// HasAuthMethod returns a boolean if a field has been set.
func (o *AlertIntegrationBase) HasAuthMethod() bool {
	if o != nil && !utils.IsNil(o.AuthMethod) {
		return true
	}

	return false
}

// SetAuthMethod gets a reference to the given string and assigns it to the AuthMethod field.
func (o *AlertIntegrationBase) SetAuthMethod(v string) {
	o.AuthMethod = &v
}

// GetAuthUser returns the AuthUser field value if set, zero value otherwise.
func (o *AlertIntegrationBase) GetAuthUser() string {
	if o == nil || utils.IsNil(o.AuthUser) {
		var ret string
		return ret
	}
	return *o.AuthUser
}

// GetAuthUserOk returns a tuple with the AuthUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertIntegrationBase) GetAuthUserOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AuthUser) {
		return nil, false
	}
	return o.AuthUser, true
}

// HasAuthUser returns a boolean if a field has been set.
func (o *AlertIntegrationBase) HasAuthUser() bool {
	if o != nil && !utils.IsNil(o.AuthUser) {
		return true
	}

	return false
}

// SetAuthUser gets a reference to the given string and assigns it to the AuthUser field.
func (o *AlertIntegrationBase) SetAuthUser(v string) {
	o.AuthUser = &v
}

// GetAuthToken returns the AuthToken field value if set, zero value otherwise.
func (o *AlertIntegrationBase) GetAuthToken() string {
	if o == nil || utils.IsNil(o.AuthToken) {
		var ret string
		return ret
	}
	return *o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertIntegrationBase) GetAuthTokenOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AuthToken) {
		return nil, false
	}
	return o.AuthToken, true
}

// HasAuthToken returns a boolean if a field has been set.
func (o *AlertIntegrationBase) HasAuthToken() bool {
	if o != nil && !utils.IsNil(o.AuthToken) {
		return true
	}

	return false
}

// SetAuthToken gets a reference to the given string and assigns it to the AuthToken field.
func (o *AlertIntegrationBase) SetAuthToken(v string) {
	o.AuthToken = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *AlertIntegrationBase) GetChannel() string {
	if o == nil || utils.IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertIntegrationBase) GetChannelOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *AlertIntegrationBase) HasChannel() bool {
	if o != nil && !utils.IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *AlertIntegrationBase) SetChannel(v string) {
	o.Channel = &v
}

func (o AlertIntegrationBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertIntegrationBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.IntegrationId) {
		toSerialize["integrationId"] = o.IntegrationId
	}
	if !utils.IsNil(o.IntegrationName) {
		toSerialize["integrationName"] = o.IntegrationName
	}
	if !utils.IsNil(o.IntegrationType) {
		toSerialize["integrationType"] = o.IntegrationType
	}
	if !utils.IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !utils.IsNil(o.AuthMethod) {
		toSerialize["authMethod"] = o.AuthMethod
	}
	if !utils.IsNil(o.AuthUser) {
		toSerialize["authUser"] = o.AuthUser
	}
	if !utils.IsNil(o.AuthToken) {
		toSerialize["authToken"] = o.AuthToken
	}
	if !utils.IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	return toSerialize, nil
}

type NullableAlertIntegrationBase struct {
	value *AlertIntegrationBase
	isSet bool
}

func (v NullableAlertIntegrationBase) Get() *AlertIntegrationBase {
	return v.value
}

func (v *NullableAlertIntegrationBase) Set(val *AlertIntegrationBase) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertIntegrationBase) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertIntegrationBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertIntegrationBase(val *AlertIntegrationBase) *NullableAlertIntegrationBase {
	return &NullableAlertIntegrationBase{value: val, isSet: true}
}

func (v NullableAlertIntegrationBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertIntegrationBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


