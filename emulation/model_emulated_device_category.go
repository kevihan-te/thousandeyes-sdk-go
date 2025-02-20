/*
Emulation API

The Emulation API facilitates the retrieval of user-agent strings for HTTP, pageload, and transaction tests. It also enables the retrieval and addition of emulated devices for pageload and transaction tests.  To access Emulation API operations, the following permissions are required:  * `Settings Tests Read` for read operations. * `Settings Tests Update` for write operations. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emulation

import (
	"encoding/json"
	"fmt"
)

// EmulatedDeviceCategory The type of device being emulated.
type EmulatedDeviceCategory string

// List of EmulatedDeviceCategory
const (
	EMULATEDDEVICECATEGORY_DESKTOP EmulatedDeviceCategory = "desktop"
	EMULATEDDEVICECATEGORY_LAPTOP EmulatedDeviceCategory = "laptop"
	EMULATEDDEVICECATEGORY_PHONE EmulatedDeviceCategory = "phone"
	EMULATEDDEVICECATEGORY_TABLET EmulatedDeviceCategory = "tablet"
)

// All allowed values of EmulatedDeviceCategory enum
var AllowedEmulatedDeviceCategoryEnumValues = []EmulatedDeviceCategory{
	"desktop",
	"laptop",
	"phone",
	"tablet",
}

func (v *EmulatedDeviceCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmulatedDeviceCategory(value)
	for _, existing := range AllowedEmulatedDeviceCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmulatedDeviceCategory", value)
}

// NewEmulatedDeviceCategoryFromValue returns a pointer to a valid EmulatedDeviceCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmulatedDeviceCategoryFromValue(v string) (*EmulatedDeviceCategory, error) {
	ev := EmulatedDeviceCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmulatedDeviceCategory: valid values are %v", v, AllowedEmulatedDeviceCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmulatedDeviceCategory) IsValid() bool {
	for _, existing := range AllowedEmulatedDeviceCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmulatedDeviceCategory value
func (v EmulatedDeviceCategory) Ptr() *EmulatedDeviceCategory {
	return &v
}

type NullableEmulatedDeviceCategory struct {
	value *EmulatedDeviceCategory
	isSet bool
}

func (v NullableEmulatedDeviceCategory) Get() *EmulatedDeviceCategory {
	return v.value
}

func (v *NullableEmulatedDeviceCategory) Set(val *EmulatedDeviceCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableEmulatedDeviceCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableEmulatedDeviceCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmulatedDeviceCategory(val *EmulatedDeviceCategory) *NullableEmulatedDeviceCategory {
	return &NullableEmulatedDeviceCategory{value: val, isSet: true}
}

func (v NullableEmulatedDeviceCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmulatedDeviceCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

