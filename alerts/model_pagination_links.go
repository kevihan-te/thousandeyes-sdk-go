/*
Alerts API

You can manage the following alert functionalities on the ThousandEyes platform using the Alerts API:  * **Alerts**: Retrieve alert details. Alerts are assigned to tests through alert rules.  * **Alert Rules**: Conditions that you configure in order to highlight or be notified of events of interest in your ThousandEyes tests. When an alert rule’s conditions are met, the associated alert is triggered and the alert becomes active. It remains active until the alert is cleared. Alert rules are reusable across multiple tests..  * **Alert Suppression Windows**: Suppress alerts for tests during periods such as planned maintenance. Windows can be one-time events or recurring events to handle periodic occurrences such as monthly downtime for maintenance.  For more information about the alerts, see [Alerts](https://docs.thousandeyes.com/product-documentation/alerts). 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alerts

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the PaginationLinks type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaginationLinks{}

// PaginationLinks A links object containing pagination related link(s).
type PaginationLinks struct {
	Previous *Link `json:"previous,omitempty"`
	Next *Link `json:"next,omitempty"`
	Self *Link `json:"self,omitempty"`
}

// NewPaginationLinks instantiates a new PaginationLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationLinks() *PaginationLinks {
	this := PaginationLinks{}
	return &this
}

// NewPaginationLinksWithDefaults instantiates a new PaginationLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationLinksWithDefaults() *PaginationLinks {
	this := PaginationLinks{}
	return &this
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *PaginationLinks) GetPrevious() Link {
	if o == nil || utils.IsNil(o.Previous) {
		var ret Link
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationLinks) GetPreviousOk() (*Link, bool) {
	if o == nil || utils.IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginationLinks) HasPrevious() bool {
	if o != nil && !utils.IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given Link and assigns it to the Previous field.
func (o *PaginationLinks) SetPrevious(v Link) {
	o.Previous = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginationLinks) GetNext() Link {
	if o == nil || utils.IsNil(o.Next) {
		var ret Link
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationLinks) GetNextOk() (*Link, bool) {
	if o == nil || utils.IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginationLinks) HasNext() bool {
	if o != nil && !utils.IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given Link and assigns it to the Next field.
func (o *PaginationLinks) SetNext(v Link) {
	o.Next = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PaginationLinks) GetSelf() Link {
	if o == nil || utils.IsNil(o.Self) {
		var ret Link
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationLinks) GetSelfOk() (*Link, bool) {
	if o == nil || utils.IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PaginationLinks) HasSelf() bool {
	if o != nil && !utils.IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Link and assigns it to the Self field.
func (o *PaginationLinks) SetSelf(v Link) {
	o.Self = &v
}

func (o PaginationLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	if !utils.IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !utils.IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullablePaginationLinks struct {
	value *PaginationLinks
	isSet bool
}

func (v NullablePaginationLinks) Get() *PaginationLinks {
	return v.value
}

func (v *NullablePaginationLinks) Set(val *PaginationLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationLinks(val *PaginationLinks) *NullablePaginationLinks {
	return &NullablePaginationLinks{value: val, isSet: true}
}

func (v NullablePaginationLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


