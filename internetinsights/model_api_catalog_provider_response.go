/*
Internet Insights API

We are happy to announce the release of the Internet Insights API set. This limited release includes endpoints that:  * Make our catalog provider and Internet outage data accessible to API users. * Provide access to advanced filtering, which is part of our next-generation API efforts to allow API users to fine-tune queries across all of our APIs in a consistent manner.  Internet Insights provide visibility into core Internet infrastructure, including ISPs, DNS providers, IaaS, CDNs , and SaaS providers. It tracks the macro-level impact of Internet events on individual users and enterprise networks connecting at the edge of the Internet. These events include Outages, Routing hijacks and leaks, DDoS attacks, And political interference, among others.  Future releases of the Internet Insights API set will further unlock access to core Internet Insights functionality, unlocking potential integrations to enrich customer process flows.  For more information about Internet Insights, see the [Internet Insights](https://docs.thousandeyes.com/product-documentation/internet-insights). 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package internetinsights

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the ApiCatalogProviderResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiCatalogProviderResponse{}

// ApiCatalogProviderResponse struct for ApiCatalogProviderResponse
type ApiCatalogProviderResponse struct {
	// List of catalog providers.
	Providers []ApiCatalogProvider `json:"providers,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewApiCatalogProviderResponse instantiates a new ApiCatalogProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCatalogProviderResponse() *ApiCatalogProviderResponse {
	this := ApiCatalogProviderResponse{}
	return &this
}

// NewApiCatalogProviderResponseWithDefaults instantiates a new ApiCatalogProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCatalogProviderResponseWithDefaults() *ApiCatalogProviderResponse {
	this := ApiCatalogProviderResponse{}
	return &this
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *ApiCatalogProviderResponse) GetProviders() []ApiCatalogProvider {
	if o == nil || utils.IsNil(o.Providers) {
		var ret []ApiCatalogProvider
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCatalogProviderResponse) GetProvidersOk() ([]ApiCatalogProvider, bool) {
	if o == nil || utils.IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *ApiCatalogProviderResponse) HasProviders() bool {
	if o != nil && !utils.IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []ApiCatalogProvider and assigns it to the Providers field.
func (o *ApiCatalogProviderResponse) SetProviders(v []ApiCatalogProvider) {
	o.Providers = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiCatalogProviderResponse) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCatalogProviderResponse) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiCatalogProviderResponse) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *ApiCatalogProviderResponse) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o ApiCatalogProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCatalogProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Providers) {
		toSerialize["providers"] = o.Providers
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableApiCatalogProviderResponse struct {
	value *ApiCatalogProviderResponse
	isSet bool
}

func (v NullableApiCatalogProviderResponse) Get() *ApiCatalogProviderResponse {
	return v.value
}

func (v *NullableApiCatalogProviderResponse) Set(val *ApiCatalogProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCatalogProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCatalogProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCatalogProviderResponse(val *ApiCatalogProviderResponse) *NullableApiCatalogProviderResponse {
	return &NullableApiCatalogProviderResponse{value: val, isSet: true}
}

func (v NullableApiCatalogProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCatalogProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


