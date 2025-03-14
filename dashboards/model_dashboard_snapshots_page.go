/*
Dashboards API

Manage ThousandEyes Dashboards.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboards

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the DashboardSnapshotsPage type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DashboardSnapshotsPage{}

// DashboardSnapshotsPage Dashboard snapshots page.
type DashboardSnapshotsPage struct {
	// Deprecated
	Pages map[string]interface{} `json:"pages,omitempty"`
	DashboardSnapshots []ApiDashboardSnapshot `json:"dashboardSnapshots,omitempty"`
	Links *PaginationLinks `json:"_links,omitempty"`
}

// NewDashboardSnapshotsPage instantiates a new DashboardSnapshotsPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardSnapshotsPage() *DashboardSnapshotsPage {
	this := DashboardSnapshotsPage{}
	return &this
}

// NewDashboardSnapshotsPageWithDefaults instantiates a new DashboardSnapshotsPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardSnapshotsPageWithDefaults() *DashboardSnapshotsPage {
	this := DashboardSnapshotsPage{}
	return &this
}

// GetPages returns the Pages field value if set, zero value otherwise.
// Deprecated
func (o *DashboardSnapshotsPage) GetPages() map[string]interface{} {
	if o == nil || utils.IsNil(o.Pages) {
		var ret map[string]interface{}
		return ret
	}
	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *DashboardSnapshotsPage) GetPagesOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Pages) {
		return map[string]interface{}{}, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *DashboardSnapshotsPage) HasPages() bool {
	if o != nil && !utils.IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given map[string]interface{} and assigns it to the Pages field.
// Deprecated
func (o *DashboardSnapshotsPage) SetPages(v map[string]interface{}) {
	o.Pages = v
}

// GetDashboardSnapshots returns the DashboardSnapshots field value if set, zero value otherwise.
func (o *DashboardSnapshotsPage) GetDashboardSnapshots() []ApiDashboardSnapshot {
	if o == nil || utils.IsNil(o.DashboardSnapshots) {
		var ret []ApiDashboardSnapshot
		return ret
	}
	return o.DashboardSnapshots
}

// GetDashboardSnapshotsOk returns a tuple with the DashboardSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSnapshotsPage) GetDashboardSnapshotsOk() ([]ApiDashboardSnapshot, bool) {
	if o == nil || utils.IsNil(o.DashboardSnapshots) {
		return nil, false
	}
	return o.DashboardSnapshots, true
}

// HasDashboardSnapshots returns a boolean if a field has been set.
func (o *DashboardSnapshotsPage) HasDashboardSnapshots() bool {
	if o != nil && !utils.IsNil(o.DashboardSnapshots) {
		return true
	}

	return false
}

// SetDashboardSnapshots gets a reference to the given []ApiDashboardSnapshot and assigns it to the DashboardSnapshots field.
func (o *DashboardSnapshotsPage) SetDashboardSnapshots(v []ApiDashboardSnapshot) {
	o.DashboardSnapshots = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DashboardSnapshotsPage) GetLinks() PaginationLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSnapshotsPage) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DashboardSnapshotsPage) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *DashboardSnapshotsPage) SetLinks(v PaginationLinks) {
	o.Links = &v
}

func (o DashboardSnapshotsPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardSnapshotsPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Pages) {
		toSerialize["pages"] = o.Pages
	}
	if !utils.IsNil(o.DashboardSnapshots) {
		toSerialize["dashboardSnapshots"] = o.DashboardSnapshots
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableDashboardSnapshotsPage struct {
	value *DashboardSnapshotsPage
	isSet bool
}

func (v NullableDashboardSnapshotsPage) Get() *DashboardSnapshotsPage {
	return v.value
}

func (v *NullableDashboardSnapshotsPage) Set(val *DashboardSnapshotsPage) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardSnapshotsPage) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardSnapshotsPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardSnapshotsPage(val *DashboardSnapshotsPage) *NullableDashboardSnapshotsPage {
	return &NullableDashboardSnapshotsPage{value: val, isSet: true}
}

func (v NullableDashboardSnapshotsPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardSnapshotsPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


