/*
Tests API

This API supports listing, creating, editing, and deleting Cloud and Enterprise Agent (CEA) based tests. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"fmt"
)

// checks if the ApiRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiRequest{}

// ApiRequest struct for ApiRequest
type ApiRequest struct {
	// List of assertion objects.
	Assertions []ApiRequestAssertion `json:"assertions,omitempty"`
	AuthType *ApiRequestAuthType `json:"authType,omitempty"`
	// The bearer token if `authType = bearer-token`.
	BearerToken *string `json:"bearerToken,omitempty"`
	// POST/PUT request body. Must be in JSON format.
	Body *string `json:"body,omitempty"`
	ClientAuthentication *ApiClientAuthentication `json:"clientAuthentication,omitempty"`
	// The application ID used when `authType` is set to \"oauth2\".
	ClientId *string `json:"clientId,omitempty"`
	// The private client secret used when `authType` is set to \"oauth2\".
	ClientSecret *string `json:"clientSecret,omitempty"`
	// Set to `true` if API response body should be collected and saved. Set to `false` if API response body should not be saved.
	CollectApiResponse *bool `json:"collectApiResponse,omitempty"`
	// Array of API Request Header objects.
	Headers []ApiRequestHeader `json:"headers,omitempty"`
	Method *ApiRequestMethod `json:"method,omitempty"`
	// API step name, must be unique.
	Name string `json:"name"`
	// The password if `authType = basic`.
	Password *string `json:"password,omitempty"`
	// Application-specific scope values for the access token when `authType` is \"oauth2\".
	Scope *string `json:"scope,omitempty"`
	// The endpoint used to request the access token when `authType` is \"oauth2\".
	TokenUrl *string `json:"tokenUrl,omitempty"`
	// Request url. Supports variables in the format `{{variableName}}`.
	Url string `json:"url"`
	// The username if `authType = basic`.
	Username *string `json:"username,omitempty"`
	// Array of API post request variable objects.
	Variables []ApiRequestVariable `json:"variables,omitempty"`
	// Post request delay before executing the next API requests, in milliseconds.
	WaitTimeMs *int32 `json:"waitTimeMs,omitempty"`
}

type _ApiRequest ApiRequest

// NewApiRequest instantiates a new ApiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRequest(name string, url string) *ApiRequest {
	this := ApiRequest{}
	var authType ApiRequestAuthType = APIREQUESTAUTHTYPE_NONE
	this.AuthType = &authType
	var collectApiResponse bool = true
	this.CollectApiResponse = &collectApiResponse
	this.Name = name
	this.Url = url
	return &this
}

// NewApiRequestWithDefaults instantiates a new ApiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRequestWithDefaults() *ApiRequest {
	this := ApiRequest{}
	var authType ApiRequestAuthType = APIREQUESTAUTHTYPE_NONE
	this.AuthType = &authType
	var collectApiResponse bool = true
	this.CollectApiResponse = &collectApiResponse
	return &this
}

// GetAssertions returns the Assertions field value if set, zero value otherwise.
func (o *ApiRequest) GetAssertions() []ApiRequestAssertion {
	if o == nil || utils.IsNil(o.Assertions) {
		var ret []ApiRequestAssertion
		return ret
	}
	return o.Assertions
}

// GetAssertionsOk returns a tuple with the Assertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetAssertionsOk() ([]ApiRequestAssertion, bool) {
	if o == nil || utils.IsNil(o.Assertions) {
		return nil, false
	}
	return o.Assertions, true
}

// HasAssertions returns a boolean if a field has been set.
func (o *ApiRequest) HasAssertions() bool {
	if o != nil && !utils.IsNil(o.Assertions) {
		return true
	}

	return false
}

// SetAssertions gets a reference to the given []ApiRequestAssertion and assigns it to the Assertions field.
func (o *ApiRequest) SetAssertions(v []ApiRequestAssertion) {
	o.Assertions = v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *ApiRequest) GetAuthType() ApiRequestAuthType {
	if o == nil || utils.IsNil(o.AuthType) {
		var ret ApiRequestAuthType
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetAuthTypeOk() (*ApiRequestAuthType, bool) {
	if o == nil || utils.IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *ApiRequest) HasAuthType() bool {
	if o != nil && !utils.IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given ApiRequestAuthType and assigns it to the AuthType field.
func (o *ApiRequest) SetAuthType(v ApiRequestAuthType) {
	o.AuthType = &v
}

// GetBearerToken returns the BearerToken field value if set, zero value otherwise.
func (o *ApiRequest) GetBearerToken() string {
	if o == nil || utils.IsNil(o.BearerToken) {
		var ret string
		return ret
	}
	return *o.BearerToken
}

// GetBearerTokenOk returns a tuple with the BearerToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetBearerTokenOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BearerToken) {
		return nil, false
	}
	return o.BearerToken, true
}

// HasBearerToken returns a boolean if a field has been set.
func (o *ApiRequest) HasBearerToken() bool {
	if o != nil && !utils.IsNil(o.BearerToken) {
		return true
	}

	return false
}

// SetBearerToken gets a reference to the given string and assigns it to the BearerToken field.
func (o *ApiRequest) SetBearerToken(v string) {
	o.BearerToken = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ApiRequest) GetBody() string {
	if o == nil || utils.IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetBodyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ApiRequest) HasBody() bool {
	if o != nil && !utils.IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *ApiRequest) SetBody(v string) {
	o.Body = &v
}

// GetClientAuthentication returns the ClientAuthentication field value if set, zero value otherwise.
func (o *ApiRequest) GetClientAuthentication() ApiClientAuthentication {
	if o == nil || utils.IsNil(o.ClientAuthentication) {
		var ret ApiClientAuthentication
		return ret
	}
	return *o.ClientAuthentication
}

// GetClientAuthenticationOk returns a tuple with the ClientAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetClientAuthenticationOk() (*ApiClientAuthentication, bool) {
	if o == nil || utils.IsNil(o.ClientAuthentication) {
		return nil, false
	}
	return o.ClientAuthentication, true
}

// HasClientAuthentication returns a boolean if a field has been set.
func (o *ApiRequest) HasClientAuthentication() bool {
	if o != nil && !utils.IsNil(o.ClientAuthentication) {
		return true
	}

	return false
}

// SetClientAuthentication gets a reference to the given ApiClientAuthentication and assigns it to the ClientAuthentication field.
func (o *ApiRequest) SetClientAuthentication(v ApiClientAuthentication) {
	o.ClientAuthentication = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ApiRequest) GetClientId() string {
	if o == nil || utils.IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetClientIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ApiRequest) HasClientId() bool {
	if o != nil && !utils.IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ApiRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ApiRequest) GetClientSecret() string {
	if o == nil || utils.IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ApiRequest) HasClientSecret() bool {
	if o != nil && !utils.IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ApiRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetCollectApiResponse returns the CollectApiResponse field value if set, zero value otherwise.
func (o *ApiRequest) GetCollectApiResponse() bool {
	if o == nil || utils.IsNil(o.CollectApiResponse) {
		var ret bool
		return ret
	}
	return *o.CollectApiResponse
}

// GetCollectApiResponseOk returns a tuple with the CollectApiResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetCollectApiResponseOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.CollectApiResponse) {
		return nil, false
	}
	return o.CollectApiResponse, true
}

// HasCollectApiResponse returns a boolean if a field has been set.
func (o *ApiRequest) HasCollectApiResponse() bool {
	if o != nil && !utils.IsNil(o.CollectApiResponse) {
		return true
	}

	return false
}

// SetCollectApiResponse gets a reference to the given bool and assigns it to the CollectApiResponse field.
func (o *ApiRequest) SetCollectApiResponse(v bool) {
	o.CollectApiResponse = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ApiRequest) GetHeaders() []ApiRequestHeader {
	if o == nil || utils.IsNil(o.Headers) {
		var ret []ApiRequestHeader
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetHeadersOk() ([]ApiRequestHeader, bool) {
	if o == nil || utils.IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ApiRequest) HasHeaders() bool {
	if o != nil && !utils.IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []ApiRequestHeader and assigns it to the Headers field.
func (o *ApiRequest) SetHeaders(v []ApiRequestHeader) {
	o.Headers = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *ApiRequest) GetMethod() ApiRequestMethod {
	if o == nil || utils.IsNil(o.Method) {
		var ret ApiRequestMethod
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetMethodOk() (*ApiRequestMethod, bool) {
	if o == nil || utils.IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *ApiRequest) HasMethod() bool {
	if o != nil && !utils.IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given ApiRequestMethod and assigns it to the Method field.
func (o *ApiRequest) SetMethod(v ApiRequestMethod) {
	o.Method = &v
}

// GetName returns the Name field value
func (o *ApiRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiRequest) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApiRequest) GetPassword() string {
	if o == nil || utils.IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetPasswordOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApiRequest) HasPassword() bool {
	if o != nil && !utils.IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApiRequest) SetPassword(v string) {
	o.Password = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ApiRequest) GetScope() string {
	if o == nil || utils.IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetScopeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ApiRequest) HasScope() bool {
	if o != nil && !utils.IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ApiRequest) SetScope(v string) {
	o.Scope = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *ApiRequest) GetTokenUrl() string {
	if o == nil || utils.IsNil(o.TokenUrl) {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetTokenUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TokenUrl) {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *ApiRequest) HasTokenUrl() bool {
	if o != nil && !utils.IsNil(o.TokenUrl) {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *ApiRequest) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

// GetUrl returns the Url field value
func (o *ApiRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ApiRequest) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiRequest) GetUsername() string {
	if o == nil || utils.IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetUsernameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiRequest) HasUsername() bool {
	if o != nil && !utils.IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiRequest) SetUsername(v string) {
	o.Username = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *ApiRequest) GetVariables() []ApiRequestVariable {
	if o == nil || utils.IsNil(o.Variables) {
		var ret []ApiRequestVariable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetVariablesOk() ([]ApiRequestVariable, bool) {
	if o == nil || utils.IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *ApiRequest) HasVariables() bool {
	if o != nil && !utils.IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []ApiRequestVariable and assigns it to the Variables field.
func (o *ApiRequest) SetVariables(v []ApiRequestVariable) {
	o.Variables = v
}

// GetWaitTimeMs returns the WaitTimeMs field value if set, zero value otherwise.
func (o *ApiRequest) GetWaitTimeMs() int32 {
	if o == nil || utils.IsNil(o.WaitTimeMs) {
		var ret int32
		return ret
	}
	return *o.WaitTimeMs
}

// GetWaitTimeMsOk returns a tuple with the WaitTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequest) GetWaitTimeMsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.WaitTimeMs) {
		return nil, false
	}
	return o.WaitTimeMs, true
}

// HasWaitTimeMs returns a boolean if a field has been set.
func (o *ApiRequest) HasWaitTimeMs() bool {
	if o != nil && !utils.IsNil(o.WaitTimeMs) {
		return true
	}

	return false
}

// SetWaitTimeMs gets a reference to the given int32 and assigns it to the WaitTimeMs field.
func (o *ApiRequest) SetWaitTimeMs(v int32) {
	o.WaitTimeMs = &v
}

func (o ApiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Assertions) {
		toSerialize["assertions"] = o.Assertions
	}
	if !utils.IsNil(o.AuthType) {
		toSerialize["authType"] = o.AuthType
	}
	if !utils.IsNil(o.BearerToken) {
		toSerialize["bearerToken"] = o.BearerToken
	}
	if !utils.IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !utils.IsNil(o.ClientAuthentication) {
		toSerialize["clientAuthentication"] = o.ClientAuthentication
	}
	if !utils.IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !utils.IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !utils.IsNil(o.CollectApiResponse) {
		toSerialize["collectApiResponse"] = o.CollectApiResponse
	}
	if !utils.IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !utils.IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	toSerialize["name"] = o.Name
	if !utils.IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !utils.IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !utils.IsNil(o.TokenUrl) {
		toSerialize["tokenUrl"] = o.TokenUrl
	}
	toSerialize["url"] = o.Url
	if !utils.IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !utils.IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	if !utils.IsNil(o.WaitTimeMs) {
		toSerialize["waitTimeMs"] = o.WaitTimeMs
	}
	return toSerialize, nil
}

func (o *ApiRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"url",
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

	varApiRequest := _ApiRequest{}

    err = json.Unmarshal(data, &varApiRequest)

	if err != nil {
		return err
	}

	*o = ApiRequest(varApiRequest)

	return err
}

type NullableApiRequest struct {
	value *ApiRequest
	isSet bool
}

func (v NullableApiRequest) Get() *ApiRequest {
	return v.value
}

func (v *NullableApiRequest) Set(val *ApiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRequest(val *ApiRequest) *NullableApiRequest {
	return &NullableApiRequest{value: val, isSet: true}
}

func (v NullableApiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


