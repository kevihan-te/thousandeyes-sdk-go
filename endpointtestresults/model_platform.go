/*
Endpoint Test Results API

Retrieve results for scheduled and dynamic tests on endpoint agents.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointtestresults

import (
	"encoding/json"
	"fmt"
)

// Platform OS platform types. Platform \"linux\" was recently renamed to \"roomos\".
type Platform string

// List of Platform
const (
	PLATFORM_WINDOWS Platform = "windows"
	PLATFORM_ROOMOS Platform = "roomos"
	PLATFORM_PHONEOS Platform = "phoneos"
	PLATFORM_ELUX Platform = "elux"
	PLATFORM_LINUX Platform = "linux"
	PLATFORM_MAC Platform = "mac"
	PLATFORM_UNKNOWN Platform = "unknown"
)

// All allowed values of Platform enum
var AllowedPlatformEnumValues = []Platform{
	"windows",
	"roomos",
	"phoneos",
	"elux",
	"linux",
	"mac",
	"unknown",
}

func (v *Platform) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Platform(value)
	for _, existing := range AllowedPlatformEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Platform", value)
}

// NewPlatformFromValue returns a pointer to a valid Platform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlatformFromValue(v string) (*Platform, error) {
	ev := Platform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Platform: valid values are %v", v, AllowedPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Platform) IsValid() bool {
	for _, existing := range AllowedPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Platform value
func (v Platform) Ptr() *Platform {
	return &v
}

type NullablePlatform struct {
	value *Platform
	isSet bool
}

func (v NullablePlatform) Get() *Platform {
	return v.value
}

func (v *NullablePlatform) Set(val *Platform) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatform) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatform(val *Platform) *NullablePlatform {
	return &NullablePlatform{value: val, isSet: true}
}

func (v NullablePlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

