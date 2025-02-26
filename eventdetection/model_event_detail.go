/*
Event Detection API

 Event detection occurs when ThousandEyes identifies that error signals related to a component (proxy, network node, AS, server etc) have deviated from the baselines established by events. * To determine this, ThousandEyes takes the test results from all accounts groups within an organization, and analyzes that data. * Noisy test results (those that have too many errors in a short window) are removed until they stabilize, and the rest of the results are tagged with the components associated with that test result (for example, proxy, network, or server). * Next, any increase in failures from the test results and each component helps in determining the problem domain and which component may be at fault. * When this failure rate increases beyond a pre-defined threshold (set by the algorithm), an event is triggered and an email notification is sent to the user (if they've enabled email alerts).  With the Events API, you can perform the following tasks on the ThousandEyes platform: * **Retrieve Events**: Obtain a list of events and detailed information for each event. For more information about events, see [Event Detection](https://docs.thousandeyes.com/product-documentation/event-detection). 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eventdetection

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"fmt"
)

// EventDetail - struct for EventDetail
type EventDetail struct {
	AgentLocalEventDetail *AgentLocalEventDetail
	DnsEventDetail *DnsEventDetail
	NetworkEventDetail *NetworkEventDetail
	NetworkPopEventDetail *NetworkPopEventDetail
	ProxyEventDetail *ProxyEventDetail
	TargetEventDetail *TargetEventDetail
	TargetNetworkEventDetail *TargetNetworkEventDetail
}

// AgentLocalEventDetailAsEventDetail is a convenience function that returns AgentLocalEventDetail wrapped in EventDetail
func AgentLocalEventDetailAsEventDetail(v *AgentLocalEventDetail) EventDetail {
	return EventDetail{
		AgentLocalEventDetail: v,
	}
}

// DnsEventDetailAsEventDetail is a convenience function that returns DnsEventDetail wrapped in EventDetail
func DnsEventDetailAsEventDetail(v *DnsEventDetail) EventDetail {
	return EventDetail{
		DnsEventDetail: v,
	}
}

// NetworkEventDetailAsEventDetail is a convenience function that returns NetworkEventDetail wrapped in EventDetail
func NetworkEventDetailAsEventDetail(v *NetworkEventDetail) EventDetail {
	return EventDetail{
		NetworkEventDetail: v,
	}
}

// NetworkPopEventDetailAsEventDetail is a convenience function that returns NetworkPopEventDetail wrapped in EventDetail
func NetworkPopEventDetailAsEventDetail(v *NetworkPopEventDetail) EventDetail {
	return EventDetail{
		NetworkPopEventDetail: v,
	}
}

// ProxyEventDetailAsEventDetail is a convenience function that returns ProxyEventDetail wrapped in EventDetail
func ProxyEventDetailAsEventDetail(v *ProxyEventDetail) EventDetail {
	return EventDetail{
		ProxyEventDetail: v,
	}
}

// TargetEventDetailAsEventDetail is a convenience function that returns TargetEventDetail wrapped in EventDetail
func TargetEventDetailAsEventDetail(v *TargetEventDetail) EventDetail {
	return EventDetail{
		TargetEventDetail: v,
	}
}

// TargetNetworkEventDetailAsEventDetail is a convenience function that returns TargetNetworkEventDetail wrapped in EventDetail
func TargetNetworkEventDetailAsEventDetail(v *TargetNetworkEventDetail) EventDetail {
	return EventDetail{
		TargetNetworkEventDetail: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EventDetail) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = utils.NewStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'agent-local'
	if jsonDict["type"] == "agent-local" {
		// try to unmarshal JSON data into AgentLocalEventDetail
		err = json.Unmarshal(data, &dst.AgentLocalEventDetail)
		if err == nil {
			return nil // data stored in dst.AgentLocalEventDetail, return on the first match
		} else {
			dst.AgentLocalEventDetail = nil
			return fmt.Errorf("failed to unmarshal EventDetail as AgentLocalEventDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'dns'
	if jsonDict["type"] == "dns" {
		// try to unmarshal JSON data into DnsEventDetail
		err = json.Unmarshal(data, &dst.DnsEventDetail)
		if err == nil {
			return nil // data stored in dst.DnsEventDetail, return on the first match
		} else {
			dst.DnsEventDetail = nil
			return fmt.Errorf("failed to unmarshal EventDetail as DnsEventDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'network'
	if jsonDict["type"] == "network" {
		// try to unmarshal JSON data into NetworkEventDetail
		err = json.Unmarshal(data, &dst.NetworkEventDetail)
		if err == nil {
			return nil // data stored in dst.NetworkEventDetail, return on the first match
		} else {
			dst.NetworkEventDetail = nil
			return fmt.Errorf("failed to unmarshal EventDetail as NetworkEventDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'network-pop'
	if jsonDict["type"] == "network-pop" {
		// try to unmarshal JSON data into NetworkPopEventDetail
		err = json.Unmarshal(data, &dst.NetworkPopEventDetail)
		if err == nil {
			return nil // data stored in dst.NetworkPopEventDetail, return on the first match
		} else {
			dst.NetworkPopEventDetail = nil
			return fmt.Errorf("failed to unmarshal EventDetail as NetworkPopEventDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'proxy'
	if jsonDict["type"] == "proxy" {
		// try to unmarshal JSON data into ProxyEventDetail
		err = json.Unmarshal(data, &dst.ProxyEventDetail)
		if err == nil {
			return nil // data stored in dst.ProxyEventDetail, return on the first match
		} else {
			dst.ProxyEventDetail = nil
			return fmt.Errorf("failed to unmarshal EventDetail as ProxyEventDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'target'
	if jsonDict["type"] == "target" {
		// try to unmarshal JSON data into TargetEventDetail
		err = json.Unmarshal(data, &dst.TargetEventDetail)
		if err == nil {
			return nil // data stored in dst.TargetEventDetail, return on the first match
		} else {
			dst.TargetEventDetail = nil
			return fmt.Errorf("failed to unmarshal EventDetail as TargetEventDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'target-network'
	if jsonDict["type"] == "target-network" {
		// try to unmarshal JSON data into TargetNetworkEventDetail
		err = json.Unmarshal(data, &dst.TargetNetworkEventDetail)
		if err == nil {
			return nil // data stored in dst.TargetNetworkEventDetail, return on the first match
		} else {
			dst.TargetNetworkEventDetail = nil
			return fmt.Errorf("failed to unmarshal EventDetail as TargetNetworkEventDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AgentLocalEventDetail'
	if jsonDict["type"] == "AgentLocalEventDetail" {
		// try to unmarshal JSON data into AgentLocalEventDetail
		err = json.Unmarshal(data, &dst.AgentLocalEventDetail)
		if err == nil {
			return nil // data stored in dst.AgentLocalEventDetail, return on the first match
		} else {
			dst.AgentLocalEventDetail = nil
			return fmt.Errorf("failed to unmarshal EventDetail as AgentLocalEventDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DnsEventDetail'
	if jsonDict["type"] == "DnsEventDetail" {
		// try to unmarshal JSON data into DnsEventDetail
		err = json.Unmarshal(data, &dst.DnsEventDetail)
		if err == nil {
			return nil // data stored in dst.DnsEventDetail, return on the first match
		} else {
			dst.DnsEventDetail = nil
			return fmt.Errorf("failed to unmarshal EventDetail as DnsEventDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'NetworkEventDetail'
	if jsonDict["type"] == "NetworkEventDetail" {
		// try to unmarshal JSON data into NetworkEventDetail
		err = json.Unmarshal(data, &dst.NetworkEventDetail)
		if err == nil {
			return nil // data stored in dst.NetworkEventDetail, return on the first match
		} else {
			dst.NetworkEventDetail = nil
			return fmt.Errorf("failed to unmarshal EventDetail as NetworkEventDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'NetworkPopEventDetail'
	if jsonDict["type"] == "NetworkPopEventDetail" {
		// try to unmarshal JSON data into NetworkPopEventDetail
		err = json.Unmarshal(data, &dst.NetworkPopEventDetail)
		if err == nil {
			return nil // data stored in dst.NetworkPopEventDetail, return on the first match
		} else {
			dst.NetworkPopEventDetail = nil
			return fmt.Errorf("failed to unmarshal EventDetail as NetworkPopEventDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ProxyEventDetail'
	if jsonDict["type"] == "ProxyEventDetail" {
		// try to unmarshal JSON data into ProxyEventDetail
		err = json.Unmarshal(data, &dst.ProxyEventDetail)
		if err == nil {
			return nil // data stored in dst.ProxyEventDetail, return on the first match
		} else {
			dst.ProxyEventDetail = nil
			return fmt.Errorf("failed to unmarshal EventDetail as ProxyEventDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TargetEventDetail'
	if jsonDict["type"] == "TargetEventDetail" {
		// try to unmarshal JSON data into TargetEventDetail
		err = json.Unmarshal(data, &dst.TargetEventDetail)
		if err == nil {
			return nil // data stored in dst.TargetEventDetail, return on the first match
		} else {
			dst.TargetEventDetail = nil
			return fmt.Errorf("failed to unmarshal EventDetail as TargetEventDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TargetNetworkEventDetail'
	if jsonDict["type"] == "TargetNetworkEventDetail" {
		// try to unmarshal JSON data into TargetNetworkEventDetail
		err = json.Unmarshal(data, &dst.TargetNetworkEventDetail)
		if err == nil {
			return nil // data stored in dst.TargetNetworkEventDetail, return on the first match
		} else {
			dst.TargetNetworkEventDetail = nil
			return fmt.Errorf("failed to unmarshal EventDetail as TargetNetworkEventDetail: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EventDetail) MarshalJSON() ([]byte, error) {
	if src.AgentLocalEventDetail != nil {
		return json.Marshal(&src.AgentLocalEventDetail)
	}

	if src.DnsEventDetail != nil {
		return json.Marshal(&src.DnsEventDetail)
	}

	if src.NetworkEventDetail != nil {
		return json.Marshal(&src.NetworkEventDetail)
	}

	if src.NetworkPopEventDetail != nil {
		return json.Marshal(&src.NetworkPopEventDetail)
	}

	if src.ProxyEventDetail != nil {
		return json.Marshal(&src.ProxyEventDetail)
	}

	if src.TargetEventDetail != nil {
		return json.Marshal(&src.TargetEventDetail)
	}

	if src.TargetNetworkEventDetail != nil {
		return json.Marshal(&src.TargetNetworkEventDetail)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EventDetail) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AgentLocalEventDetail != nil {
		return obj.AgentLocalEventDetail
	}

	if obj.DnsEventDetail != nil {
		return obj.DnsEventDetail
	}

	if obj.NetworkEventDetail != nil {
		return obj.NetworkEventDetail
	}

	if obj.NetworkPopEventDetail != nil {
		return obj.NetworkPopEventDetail
	}

	if obj.ProxyEventDetail != nil {
		return obj.ProxyEventDetail
	}

	if obj.TargetEventDetail != nil {
		return obj.TargetEventDetail
	}

	if obj.TargetNetworkEventDetail != nil {
		return obj.TargetNetworkEventDetail
	}

	// all schemas are nil
	return nil
}

type NullableEventDetail struct {
	value *EventDetail
	isSet bool
}

func (v NullableEventDetail) Get() *EventDetail {
	return v.value
}

func (v *NullableEventDetail) Set(val *EventDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableEventDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableEventDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventDetail(val *EventDetail) *NullableEventDetail {
	return &NullableEventDetail{value: val, isSet: true}
}

func (v NullableEventDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


