/*
Alerts API

You can manage the following alert functionalities on the ThousandEyes platform using the Alerts API:  * **Alerts**: Retrieve alert details. Alerts are assigned to tests through alert rules.  * **Alert Rules**: Conditions that you configure in order to highlight or be notified of events of interest in your ThousandEyes tests. When an alert rule’s conditions are met, the associated alert is triggered and the alert becomes active. It remains active until the alert is cleared. Alert rules are reusable across multiple tests..  * **Alert Suppression Windows**: Suppress alerts for tests during periods such as planned maintenance. Windows can be one-time events or recurring events to handle periodic occurrences such as monthly downtime for maintenance.  For more information about the alerts, see [Alerts](https://docs.thousandeyes.com/product-documentation/alerts). 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alerts

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the Alerts type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Alerts{}

// Alerts struct for Alerts
type Alerts struct {
	Alerts []Alert `json:"alerts,omitempty"`
	Links *PaginationLinks `json:"_links,omitempty"`
}

// NewAlerts instantiates a new Alerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlerts() *Alerts {
	this := Alerts{}
	return &this
}

// NewAlertsWithDefaults instantiates a new Alerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsWithDefaults() *Alerts {
	this := Alerts{}
	return &this
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *Alerts) GetAlerts() []Alert {
	if o == nil || utils.IsNil(o.Alerts) {
		var ret []Alert
		return ret
	}
	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alerts) GetAlertsOk() ([]Alert, bool) {
	if o == nil || utils.IsNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *Alerts) HasAlerts() bool {
	if o != nil && !utils.IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []Alert and assigns it to the Alerts field.
func (o *Alerts) SetAlerts(v []Alert) {
	o.Alerts = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Alerts) GetLinks() PaginationLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alerts) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Alerts) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *Alerts) SetLinks(v PaginationLinks) {
	o.Links = &v
}

func (o Alerts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Alerts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableAlerts struct {
	value *Alerts
	isSet bool
}

func (v NullableAlerts) Get() *Alerts {
	return v.value
}

func (v *NullableAlerts) Set(val *Alerts) {
	v.value = val
	v.isSet = true
}

func (v NullableAlerts) IsSet() bool {
	return v.isSet
}

func (v *NullableAlerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlerts(val *Alerts) *NullableAlerts {
	return &NullableAlerts{value: val, isSet: true}
}

func (v NullableAlerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


