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

// checks if the AlertNotification type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AlertNotification{}

// AlertNotification Alert notification object. See Alert notification integrations.
type AlertNotification struct {
	Email *NotificationEmail `json:"email,omitempty"`
	// Third party notifications.
	ThirdParty []NotificationThirdParty `json:"thirdParty,omitempty"`
	// Webhook notifications.
	Webhook []NotificationWebhook `json:"webhook,omitempty"`
	// Custom webhook notifications.
	CustomWebhook []NotificationCustomWebhook `json:"customWebhook,omitempty"`
}

// NewAlertNotification instantiates a new AlertNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertNotification() *AlertNotification {
	this := AlertNotification{}
	return &this
}

// NewAlertNotificationWithDefaults instantiates a new AlertNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertNotificationWithDefaults() *AlertNotification {
	this := AlertNotification{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AlertNotification) GetEmail() NotificationEmail {
	if o == nil || utils.IsNil(o.Email) {
		var ret NotificationEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertNotification) GetEmailOk() (*NotificationEmail, bool) {
	if o == nil || utils.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AlertNotification) HasEmail() bool {
	if o != nil && !utils.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NotificationEmail and assigns it to the Email field.
func (o *AlertNotification) SetEmail(v NotificationEmail) {
	o.Email = &v
}

// GetThirdParty returns the ThirdParty field value if set, zero value otherwise.
func (o *AlertNotification) GetThirdParty() []NotificationThirdParty {
	if o == nil || utils.IsNil(o.ThirdParty) {
		var ret []NotificationThirdParty
		return ret
	}
	return o.ThirdParty
}

// GetThirdPartyOk returns a tuple with the ThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertNotification) GetThirdPartyOk() ([]NotificationThirdParty, bool) {
	if o == nil || utils.IsNil(o.ThirdParty) {
		return nil, false
	}
	return o.ThirdParty, true
}

// HasThirdParty returns a boolean if a field has been set.
func (o *AlertNotification) HasThirdParty() bool {
	if o != nil && !utils.IsNil(o.ThirdParty) {
		return true
	}

	return false
}

// SetThirdParty gets a reference to the given []NotificationThirdParty and assigns it to the ThirdParty field.
func (o *AlertNotification) SetThirdParty(v []NotificationThirdParty) {
	o.ThirdParty = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *AlertNotification) GetWebhook() []NotificationWebhook {
	if o == nil || utils.IsNil(o.Webhook) {
		var ret []NotificationWebhook
		return ret
	}
	return o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertNotification) GetWebhookOk() ([]NotificationWebhook, bool) {
	if o == nil || utils.IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *AlertNotification) HasWebhook() bool {
	if o != nil && !utils.IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given []NotificationWebhook and assigns it to the Webhook field.
func (o *AlertNotification) SetWebhook(v []NotificationWebhook) {
	o.Webhook = v
}

// GetCustomWebhook returns the CustomWebhook field value if set, zero value otherwise.
func (o *AlertNotification) GetCustomWebhook() []NotificationCustomWebhook {
	if o == nil || utils.IsNil(o.CustomWebhook) {
		var ret []NotificationCustomWebhook
		return ret
	}
	return o.CustomWebhook
}

// GetCustomWebhookOk returns a tuple with the CustomWebhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertNotification) GetCustomWebhookOk() ([]NotificationCustomWebhook, bool) {
	if o == nil || utils.IsNil(o.CustomWebhook) {
		return nil, false
	}
	return o.CustomWebhook, true
}

// HasCustomWebhook returns a boolean if a field has been set.
func (o *AlertNotification) HasCustomWebhook() bool {
	if o != nil && !utils.IsNil(o.CustomWebhook) {
		return true
	}

	return false
}

// SetCustomWebhook gets a reference to the given []NotificationCustomWebhook and assigns it to the CustomWebhook field.
func (o *AlertNotification) SetCustomWebhook(v []NotificationCustomWebhook) {
	o.CustomWebhook = v
}

func (o AlertNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !utils.IsNil(o.ThirdParty) {
		toSerialize["thirdParty"] = o.ThirdParty
	}
	if !utils.IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	if !utils.IsNil(o.CustomWebhook) {
		toSerialize["customWebhook"] = o.CustomWebhook
	}
	return toSerialize, nil
}

type NullableAlertNotification struct {
	value *AlertNotification
	isSet bool
}

func (v NullableAlertNotification) Get() *AlertNotification {
	return v.value
}

func (v *NullableAlertNotification) Set(val *AlertNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertNotification(val *AlertNotification) *NullableAlertNotification {
	return &NullableAlertNotification{value: val, isSet: true}
}

func (v NullableAlertNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


