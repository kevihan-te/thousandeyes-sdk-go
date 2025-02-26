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

// checks if the NotificationEmail type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NotificationEmail{}

// NotificationEmail Email notifications.
type NotificationEmail struct {
	// An array containing the email addresses to receive notifications.
	Recipients []string `json:"recipients,omitempty"`
	// Custom text included in alert email notifications sent to recipients.
	Message *string `json:"message,omitempty"`
}

// NewNotificationEmail instantiates a new NotificationEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationEmail() *NotificationEmail {
	this := NotificationEmail{}
	return &this
}

// NewNotificationEmailWithDefaults instantiates a new NotificationEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationEmailWithDefaults() *NotificationEmail {
	this := NotificationEmail{}
	return &this
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *NotificationEmail) GetRecipients() []string {
	if o == nil || utils.IsNil(o.Recipients) {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEmail) GetRecipientsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *NotificationEmail) HasRecipients() bool {
	if o != nil && !utils.IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []string and assigns it to the Recipients field.
func (o *NotificationEmail) SetRecipients(v []string) {
	o.Recipients = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NotificationEmail) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEmail) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NotificationEmail) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *NotificationEmail) SetMessage(v string) {
	o.Message = &v
}

func (o NotificationEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableNotificationEmail struct {
	value *NotificationEmail
	isSet bool
}

func (v NullableNotificationEmail) Get() *NotificationEmail {
	return v.value
}

func (v *NullableNotificationEmail) Set(val *NotificationEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationEmail(val *NotificationEmail) *NullableNotificationEmail {
	return &NullableNotificationEmail{value: val, isSet: true}
}

func (v NullableNotificationEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


