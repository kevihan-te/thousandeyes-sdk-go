/*
Internet Insights API

We are happy to announce the release of the Internet Insights API set. This limited release includes endpoints that:  * Make our catalog provider and Internet outage data accessible to API users. * Provide access to advanced filtering, which is part of our next-generation API efforts to allow API users to fine-tune queries across all of our APIs in a consistent manner.  Internet Insights provide visibility into core Internet infrastructure, including ISPs, DNS providers, IaaS, CDNs , and SaaS providers. It tracks the macro-level impact of Internet events on individual users and enterprise networks connecting at the edge of the Internet. These events include Outages, Routing hijacks and leaks, DDoS attacks, And political interference, among others.  Future releases of the Internet Insights API set will further unlock access to core Internet Insights functionality, unlocking potential integrations to enrich customer process flows.  For more information about Internet Insights, see the [Internet Insights](https://docs.thousandeyes.com/product-documentation/internet-insights). 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package internetinsights

import (
	"encoding/json"
	"fmt"
)

// OutageScope Scope of the outage
type OutageScope string

// List of OutageScope
const (
	OUTAGESCOPE_ALL OutageScope = "all"
	OUTAGESCOPE_WITH_AFFECTED_TEST OutageScope = "with-affected-test"
)

// All allowed values of OutageScope enum
var AllowedOutageScopeEnumValues = []OutageScope{
	"all",
	"with-affected-test",
}

func (v *OutageScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OutageScope(value)
	for _, existing := range AllowedOutageScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OutageScope", value)
}

// NewOutageScopeFromValue returns a pointer to a valid OutageScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOutageScopeFromValue(v string) (*OutageScope, error) {
	ev := OutageScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OutageScope: valid values are %v", v, AllowedOutageScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OutageScope) IsValid() bool {
	for _, existing := range AllowedOutageScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OutageScope value
func (v OutageScope) Ptr() *OutageScope {
	return &v
}

type NullableOutageScope struct {
	value *OutageScope
	isSet bool
}

func (v NullableOutageScope) Get() *OutageScope {
	return v.value
}

func (v *NullableOutageScope) Set(val *OutageScope) {
	v.value = val
	v.isSet = true
}

func (v NullableOutageScope) IsSet() bool {
	return v.isSet
}

func (v *NullableOutageScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutageScope(val *OutageScope) *NullableOutageScope {
	return &NullableOutageScope{value: val, isSet: true}
}

func (v NullableOutageScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutageScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

