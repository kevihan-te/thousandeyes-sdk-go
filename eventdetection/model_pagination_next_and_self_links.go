/*
Event Detection API

 Event detection occurs when ThousandEyes identifies that error signals related to a component (proxy, network node, AS, server etc) have deviated from the baselines established by events. * To determine this, ThousandEyes takes the test results from all accounts groups within an organization, and analyzes that data. * Noisy test results (those that have too many errors in a short window) are removed until they stabilize, and the rest of the results are tagged with the components associated with that test result (for example, proxy, network, or server). * Next, any increase in failures from the test results and each component helps in determining the problem domain and which component may be at fault. * When this failure rate increases beyond a pre-defined threshold (set by the algorithm), an event is triggered and an email notification is sent to the user (if they've enabled email alerts).  With the Events API, you can perform the following tasks on the ThousandEyes platform: * **Retrieve Events**: Obtain a list of events and detailed information for each event. For more information about events, see [Event Detection](https://docs.thousandeyes.com/product-documentation/event-detection). 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eventdetection

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the PaginationNextAndSelfLinks type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaginationNextAndSelfLinks{}

// PaginationNextAndSelfLinks A links object containing pagination-related links, including only next and self.
type PaginationNextAndSelfLinks struct {
	Next *Link `json:"next,omitempty"`
	Self *Link `json:"self,omitempty"`
}

// NewPaginationNextAndSelfLinks instantiates a new PaginationNextAndSelfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationNextAndSelfLinks() *PaginationNextAndSelfLinks {
	this := PaginationNextAndSelfLinks{}
	return &this
}

// NewPaginationNextAndSelfLinksWithDefaults instantiates a new PaginationNextAndSelfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationNextAndSelfLinksWithDefaults() *PaginationNextAndSelfLinks {
	this := PaginationNextAndSelfLinks{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginationNextAndSelfLinks) GetNext() Link {
	if o == nil || utils.IsNil(o.Next) {
		var ret Link
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationNextAndSelfLinks) GetNextOk() (*Link, bool) {
	if o == nil || utils.IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginationNextAndSelfLinks) HasNext() bool {
	if o != nil && !utils.IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given Link and assigns it to the Next field.
func (o *PaginationNextAndSelfLinks) SetNext(v Link) {
	o.Next = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PaginationNextAndSelfLinks) GetSelf() Link {
	if o == nil || utils.IsNil(o.Self) {
		var ret Link
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationNextAndSelfLinks) GetSelfOk() (*Link, bool) {
	if o == nil || utils.IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PaginationNextAndSelfLinks) HasSelf() bool {
	if o != nil && !utils.IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Link and assigns it to the Self field.
func (o *PaginationNextAndSelfLinks) SetSelf(v Link) {
	o.Self = &v
}

func (o PaginationNextAndSelfLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationNextAndSelfLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !utils.IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullablePaginationNextAndSelfLinks struct {
	value *PaginationNextAndSelfLinks
	isSet bool
}

func (v NullablePaginationNextAndSelfLinks) Get() *PaginationNextAndSelfLinks {
	return v.value
}

func (v *NullablePaginationNextAndSelfLinks) Set(val *PaginationNextAndSelfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationNextAndSelfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationNextAndSelfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationNextAndSelfLinks(val *PaginationNextAndSelfLinks) *NullablePaginationNextAndSelfLinks {
	return &NullablePaginationNextAndSelfLinks{value: val, isSet: true}
}

func (v NullablePaginationNextAndSelfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationNextAndSelfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


